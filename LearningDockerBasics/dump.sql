CREATE TABLE ITEMS (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(255),
    created_at timestamp
);