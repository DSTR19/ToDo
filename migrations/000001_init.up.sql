CREATE SCHEMA todoapp;
CREATE TABLE todoapp.users (
    id           SERIAL                   PRIMARY KEY,
    version      BIGINT       NOT NULL    DEFAULT 1,
    full_name    varchar(100) NOT NULL    CHECK (char_length(full_name) between 3 and 100),
    phone_number varchar(15)              CHECK (
        phone_number ~ '^\+[0-9]+$' 
        AND
        char_length(phone_number) between 10 and 15
    )
);

CREATE TABLE todoapp.tasks (
    id          SERIAL                   PRIMARY KEY,
    version     BIGINT       NOT NULL    DEFAULT 1,
    user_id     INT          NOT NULL    REFERENCES todoapp.users(id) ON DELETE CASCADE,
    title       varchar(100) NOT NULL    CHECK (char_length(title) between 1 and 100),
    description VARCHAR(1000)            CHECK (char_length(description) between 1 and 1000), 
    completed   BOOLEAN      NOT NULL    DEFAULT FALSE,
    created_at  TIMESTAMPTZ  NOT NULL    DEFAULT NOW(),
    completed_at TIMESTAMPTZ,
    CHECK (
    (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= created_at) 
    OR
    (completed = FALSE AND completed_at IS NULL)
    ),

    autor_user_id INT NOT NULL REFERENCES todoapp.users(id)
);   
