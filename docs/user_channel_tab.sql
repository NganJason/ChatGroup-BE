CREATE DATABASE `chat_group_db`;

CREATE TABLE `chat_group_db`.`user_channel_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `channel_id` bigint(20) unsigned NOT NULL,
    `unread` int(11) unsigned DEFAULT 0,
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`)
);