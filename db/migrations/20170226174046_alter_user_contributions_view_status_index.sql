
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contributions` ADD INDEX `view_status_index` (view_status);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contributions` DROP INDEX `view_status_index`;
