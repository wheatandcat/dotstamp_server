
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_master ADD profile_image_id BIGINT NOT NULL AFTER password;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_master DROP COLUMN profile_image_id;
