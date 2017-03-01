
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `contribution_total_follows` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'ユーザ投稿ID',
    `count` BIGINT NOT NULL COMMENT 'フォロー数',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    `deleted_at` DATETIME NULL COMMENT '削除日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `contribution_total_follows`;
