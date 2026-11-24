SELECT id, project_id, title, status, priority
FROM tasks
WHERE status = 'new' AND project_id = 1
ORDER BY priority DESC, id ASC;

SELECT p.name, COUNT(t.id) AS tasks_total
FROM projects p
LEFT JOIN tasks t ON t.project_id = p.id
GROUP BY p.id, p.name
ORDER BY p.id ASC;

INSERT INTO projects (name) VALUES('Аналитика');
SELECT p.name, COUNT(t.id) AS tasks_total
FROM projects p
LEFT JOIN tasks t ON t.project_id = p.id
GROUP BY p.id, p.name
ORDER BY p.id ASC; 

UPDATE tasks
SET status = 'done', priority = 2
WHERE id = 1
RETURNING id, title, status, priority;