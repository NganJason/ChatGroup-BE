CREATE DATABASE `chat_group_db`;

CREATE TABLE `chat_group_db`.`message_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `message_id` bigint(20) unsigned NOT NULL,
    `channel_id` bigint(20) unsigned NOT NULL,
    `user_id` bigint(20) unsigned NOT NULL,
    `content` TEXT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`, `channel_id`)
);