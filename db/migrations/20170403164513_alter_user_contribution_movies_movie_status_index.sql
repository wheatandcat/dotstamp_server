
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contribution_movies` ADD INDEX `movie_status_index` (movie_status);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contribution_movies` DROP INDEX `movie_status_index`;
