
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_contribution_searches ADD follow_count BIGINT NOT NULL AFTER search;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_contribution_searches DROP COLUMN follow_count;
