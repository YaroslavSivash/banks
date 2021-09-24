
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS banks
(
    id serial primary key unique,
    bank_name varchar unique not null,
    interest_rate integer not null,
    maximum_loan integer not null,
    minimum_down_payment integer not null,
    loan_term integer not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS banks;
