
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_masters ADD profile_image_id BIGINT NOT NULL COMMENT 'プロフィール画像ID' AFTER password;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_masters DROP COLUMN profile_image_id;
