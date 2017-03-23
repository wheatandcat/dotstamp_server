
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_movies` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'ユーザ投稿ID',
    `movie_type` INT(3) DEFAULT 1 COMMENT '動画タイプ',
    `movie_id` VARCHAR(100) COMMENT '動画ID',
    `movie_status` INT(3) DEFAULT 1 COMMENT '動画状態',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    `deleted_at` DATETIME NULL COMMENT '削除日時',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_movie_status_index` (user_contribution_id, movie_status)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_movies`;
