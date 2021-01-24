-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_resellers`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `subdistrict_id` INT NOT NULL,
    `address` TEXT NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT user_resellers_user_id FOREIGN KEY(user_id) references users(id),
    CONSTRAINT user_resellers_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
);

-- +migrate Down
DROP TABLE `user_resellers`;