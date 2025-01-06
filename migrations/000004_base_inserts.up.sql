INSERT INTO "users"(name, email)
VALUES ('Dima', 'hihihaha@golang.com'),
       ('Sofa', 'spletni@csharp.com'),
       ('Lusya', 'jwt_auth@za-pol-minuti.com'),
       ('Sasha', 'otvet-mne-pz@ymolyau.com');

INSERT INTO "tasks"(name, description, is_completed)
VALUES ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false),
       ('помыть попу', 'помыть попу', false);

INSERT INTO "users_tasks"(user_id, task_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (2, 4),
       (2, 5),
       (2, 6),
       (3, 7),
       (3, 8),
       (3, 9),
       (4, 10),
       (4, 11),
       (4, 12);
