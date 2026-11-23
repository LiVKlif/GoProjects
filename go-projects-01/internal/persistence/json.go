package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

// TODO: if path is absent, return an empty state with NextID=1.
// If the file exists, decode exactly one JSON document, validate it via Restore,
// and return a meaningful error for corrupted data. Do NOT silently reset it.
func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return tasks.State{NextID: 1}, nil
	}
	if err != nil {
		return tasks.State{}, fmt.Errorf("чтение файла %s:%w", path, err)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var state tasks.State
	if err := dec.Decode(&state); err != nil {
		return tasks.State{}, fmt.Errorf("декодирование json: %w", err)
	}
	var extra any
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return tasks.State{}, fmt.Errorf("лишние данные после json-документа: %w", tasks.ErrInvalid)
		}
		return tasks.State{}, fmt.Errorf("проверка конца json-документа: %w", err)
	}
	if err := tasks.NewMemoryStore().Restore(state); err != nil {
		return tasks.State{}, err
	}
	return state, nil
}

// TODO: validate the state, marshal it using encoding/json, save to a
// temporary file in the same directory and replace the target using os.Rename.
// Do not remove the old file before a successful temporary write.
func Save(path string, state tasks.State) error {
	if err := tasks.NewMemoryStore().Restore(state); err != nil {
		return err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("кодирование json: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("временный файл: %w", err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("запись json: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("синхронизация файла: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("закрытие файла: %w", err)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("замена %s: %w", path, err)
	}
	return nil
}
