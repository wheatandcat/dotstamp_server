
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_masters` ADD INDEX `email_index` (email);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_masters` DROP INDEX `email_index`;
