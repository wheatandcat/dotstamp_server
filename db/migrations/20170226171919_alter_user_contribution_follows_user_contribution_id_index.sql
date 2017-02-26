
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contribution_follows` ADD INDEX `user_contribution_id_index` (user_id, user_contribution_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contribution_follows` DROP INDEX `user_id_user_contribution_id_index`;
