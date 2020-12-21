-- +migrate Up
CREATE TABLE IF NOT EXISTS `users`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `role_id` TINYINT UNSIGNED NOT NULL,
    `gender` enum('male', 'female') NOT NULL DEFAULT 'male',
    `birthdate` date NOT NULL,
    `status` VARCHAR(10) UNSIGNED NOT NULL DEFAULT "A",
    `created_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `forgot_papssword_token` VARCHAR(255) DEFAULT NULL

) Engine = InnoDB;

ALTER TABLE
    `users`
ADD
    UNIQUE KEY `users_email_unique` (`email`);

-- +migrate Down
DROP TABLE `users`;