-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_resellers`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `reseller_role_id` TINYINT UNSIGNED NOT NULL,
    `location_type` VARCHAR(255) NOT NULL,
    `location_id` INT NOT NULL,
    `address` TEXT NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT user_resellers_user_id FOREIGN KEY(user_id) references users(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `user_resellers`;