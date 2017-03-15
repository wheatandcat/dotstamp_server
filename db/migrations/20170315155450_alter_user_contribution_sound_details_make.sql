
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_contribution_sound_details ADD make_status INT NOT NULL DEFAULT 0 COMMENT '作成状態' AFTER voice_type;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_contribution_sound_details DROP COLUMN make_status;
