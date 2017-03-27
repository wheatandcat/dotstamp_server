
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contribution_movies` ADD INDEX `user_contribution_id_movie_type_movie_status_index` (user_contribution_id, movie_type, movie_status);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contribution_movies` DROP INDEX `user_contribution_id_movie_type_movie_status_index`;
