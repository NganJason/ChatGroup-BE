CREATE DATABASE `chat_group_db`;

CREATE TABLE `chat_group_db`.`channel_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `channel_id` bigint(20) unsigned NOT NULL,
    `channel_name` varchar(32) DEFAULT '',
    `channel_desc` TEXT,
    `status` int(11) unsigned DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`, `channel_id`)
);