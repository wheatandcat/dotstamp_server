
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_character_images ADD voice_type INT NOT NULL DEFAULT 1 COMMENT '音声状態' AFTER priority;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_character_images DROP COLUMN voice_type;
