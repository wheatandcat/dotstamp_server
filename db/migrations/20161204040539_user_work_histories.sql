
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_work_histories` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT 'ユーザID',
    `work_id` BIGINT NOT NULL COMMENT '作品ID',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    `deleted_at` DATETIME NULL COMMENT '削除日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_work_histories`;
