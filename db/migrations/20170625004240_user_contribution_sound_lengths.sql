
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_sound_lengths` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'ユーザ投稿ID',
    `second` INT(10) DEFAULT 0 COMMENT '再生時間',
    `length` INT(10) DEFAULT 0 COMMENT '文字数',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    `deleted_at` DATETIME NULL COMMENT '削除日時',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_index` (user_contribution_id)
) COMMENT = '投稿音声長さ' ENGINE = InnoDB ;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_sound_lengths`;
