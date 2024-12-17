CREATE TABLE IF NOT EXISTS users_tasks
(
    id      SERIAL PRIMARY KEY,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users (id),
    task_id int,
    FOREIGN KEY (task_id) REFERENCES tasks (id)
);