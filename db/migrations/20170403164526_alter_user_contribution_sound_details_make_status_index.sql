
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contribution_sound_details` ADD INDEX `make_status` (make_status);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contribution_sound_details` DROP INDEX `make_status`;
