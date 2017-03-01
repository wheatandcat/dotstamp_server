
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_contribution_sound_details` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'ユーザ投稿ID',
    `priority` INT COMMENT '優先度',
    `talk_type` INT COMMENT '会話タイプ',
    `body` VARCHAR(256) COMMENT '本文',
    `body_sound` VARCHAR(256) COMMENT '本文読み上げ',
    `voice_type` INT COMMENT '音声タイプ',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    `deleted_at` DATETIME NULL COMMENT '削除日時',
    PRIMARY KEY (`id`),
    INDEX `user_contribution_id_priority_index` (user_contribution_id, priority)
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_sound_details`;
