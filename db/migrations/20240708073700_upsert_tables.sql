-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, surname, patronymic, address, passport, passport_series, passport_number)
VALUES ('First', 'First 2', 'First 3', 'First 4', '1234 123456', '1234', '123456');

INSERT INTO users (name, surname, patronymic, address, passport, passport_series, passport_number)
VALUES ('Second', 'Second 2', 'Second 3', 'Second 4', '2234 223456', '2234', '223456');

INSERT INTO users (name, surname, patronymic, address, passport, passport_series, passport_number)
VALUES ('Third', 'Third 2', 'Third 3', 'Third 4', '3234 323456', '3234', '323456');

INSERT INTO tasks (name, hours, start_task, end_task, user_id)
VALUES ('First Task', 1.5, 1633644600, 1633644872, (SELECT id FROM users WHERE passport = '1234 123456')),
       ('Second Task', 2.0, 1633645000, 1633645300, (SELECT id FROM users WHERE passport = '2234 223456')),
       ('Third Task', 1.0, 1633645500, 1633645600, (SELECT id FROM users WHERE passport = '1234 123456')),
       ('First Task', 1.5, 1633644600, 1633644872, (SELECT id FROM users WHERE passport = '2234 223456')),
       ('Second Task', 2.0, 1633645000, 1633645300, (SELECT id FROM users WHERE passport = '2234 223456')),
       ('Third Task', 1.0, 1633645500, 1633645600, (SELECT id FROM users WHERE passport = '2234 223456')),
       ('First Task', 1.5, 1633644600, 1633644872, (SELECT id FROM users WHERE passport = '3234 323456')),
       ('Second Task', 2.0, 1633645000, 1633645300, (SELECT id FROM users WHERE passport = '3234 323456')),
       ('Third Task', 1.0, 1633645500, 1633645600, (SELECT id FROM users WHERE passport = '3234 323456'));
-- +goose StatementEnd
