
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT 'ユーザID',
    `title` VARCHAR(100) COMMENT 'タイトル',
    `view_status` INT(10) DEFAULT 1 COMMENT '表示状態',
    `delete_flag` TINYINT(1) DEFAULT 0 COMMENT '削除フラグ',
    `created` DATETIME COMMENT '作成日時',
    `updated` DATETIME COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution`;
