CREATE DATABASE `chat_group_db`;

CREATE TABLE `chat_group_db`.`message_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `message_id` bigint(20) unsigned NOT NULL,
    `channel_id` bigint(20) unsigned NOT NULL,
    `user_id` bigint(20) unsigned NOT NULL,
    `content` TEXT,
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`, `channel_id`)
);