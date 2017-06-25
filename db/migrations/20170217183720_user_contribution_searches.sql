
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSTALL PLUGIN Mroonga SONAME 'ha_mroonga.so';
CREATE FUNCTION last_insert_grn_id RETURNS INTEGER SONAME 'ha_mroonga.so';
CREATE FUNCTION mroonga_snippet RETURNS STRING SONAME 'ha_mroonga.so';
CREATE FUNCTION mroonga_command RETURNS STRING SONAME 'ha_mroonga.so';
CREATE FUNCTION mroonga_escape RETURNS STRING SONAME 'ha_mroonga.so';

CREATE TABLE IF NOT EXISTS `user_contribution_searches` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_contribution_id` BIGINT NOT NULL COMMENT 'ユーザ投稿ID',
    `search` TEXT COMMENT '検索',
    `created_at` DATETIME NULL COMMENT '作成日時',
    `updated_at` DATETIME NULL COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE = mroonga;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_contribution_searches`;
