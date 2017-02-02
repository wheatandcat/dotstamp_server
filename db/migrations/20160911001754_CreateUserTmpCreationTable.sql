
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `tmp_user_creation` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT 'ユーザID',
    `category_id` INT NOT NULL COMMENT 'カテゴリ',
    `producer_id` INT NOT NULL COMMENT '製作者',
    `name` VARCHAR(255) NOT NULL COMMENT 'タイトル',
    `release_date` DATE NULL COMMENT '公開日',
    `delete_flag` DATETIME DEFAULT NULL COMMENT '削除フラグ',
    `created` DATETIME COMMENT '作成日時',
    `updated` DATETIME COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `tmp_user_creation`;
