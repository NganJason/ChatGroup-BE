CREATE DATABASE `chat_group_db`;

CREATE TABLE `chat_group_db`.`user_tab` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `username` varchar(32) DEFAULT '',
    `hashed_password` blob,
    `salt` varchar(32) DEFAULT '',
    `email_address` varchar(32) DEFAULT '',
    `photo_url` varchar(32) DEFAULT '',
    `created_at` bigint(20),
    `updated_at` bigint(20),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`id`,`user_id`)
);