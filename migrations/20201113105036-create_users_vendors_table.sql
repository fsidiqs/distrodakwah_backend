-- +migrate Up
CREATE TABLE IF NOT EXISTS `users_vendors`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `province` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL,
    `subdistrict` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `gender` enum('male', 'female') NOT NULL DEFAULT 'male',
    `birthdate` date NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `users_vendors`;