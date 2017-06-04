
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table contribution_total_follows comment '投稿統計フォロワー';
alter table goose_db_version comment 'DBバージョン';
alter table log_bug_reports comment 'バグレポートログ';
alter table log_contribution_images comment '投稿画像ログ';
alter table log_problem_contribution_reports comment '不適切投稿レポートログ';
alter table log_questions comment '質問ログ';
alter table log_user_contributions comment '投稿ログ';
alter table user_character_images comment 'キャラクタ画像';
alter table user_contribution_details comment '投稿詳細';
alter table user_contribution_follows comment '投稿フォロー';
alter table user_contribution_movies comment '投稿動画';
alter table user_contribution_searches comment '投稿検索';
alter table user_contribution_sounds comment '投稿音声';
alter table user_contribution_sound_details comment '投稿音声詳細';
alter table user_contribution_tags comment '投稿タグ';
alter table user_contribution_uploads comment '投稿アップロード';
alter table user_contributions comment '投稿';
alter table user_forget_passwords comment '忘れたパスワード';
alter table user_masters comment 'ユーザ';
alter table user_profile_images comment 'プロフィール画像';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table contribution_total_follows comment '';
alter table goose_db_version comment '';
alter table log_bug_reports comment '';
alter table log_contribution_images comment '';
alter table log_problem_contribution_reports comment '';
alter table log_questions comment '';
alter table log_user_contributions comment '';
alter table user_character_images comment '';
alter table user_contribution_details comment '';
alter table user_contribution_follows comment '';
alter table user_contribution_movies comment '';
alter table user_contribution_searches comment '';
alter table user_contribution_sounds comment '';
alter table user_contribution_sound_details comment '';
alter table user_contribution_tags comment '';
alter table user_contribution_uploads comment '';
alter table user_contributions comment '';
alter table user_forget_passwords comment '';
alter table user_masters comment '';
alter table user_profile_images comment '';
