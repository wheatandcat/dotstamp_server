
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `tmp_work` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT 'ユーザID',
    `name` VARCHAR(100) COMMENT 'タイトル',
    `category_id` INT NOT NULL COMMENT 'カテゴリID',
    `author_id` INT NOT NULL COMMENT '著者ID',
    `country_id` INT COMMENT '国ID',
    `released` DATE COMMENT '公開日',
    `delete_flag` TINYINT(1) DEFAULT 0 COMMENT '削除フラグ',
    `created` DATETIME COMMENT '作成日時',
    `updated` DATETIME COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `tmp_work`;
