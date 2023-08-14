SELECT DISTINCT u.id, u.name FROM orders AS o
JOIN users AS u ON u.id = o.user_id
ORDER BY u.name ASC, u.id ASC;