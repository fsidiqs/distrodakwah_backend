-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_vendors`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `subdistrict_id` INT NOT NULL,
    `address` TEXT NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT user_vendors_user_id FOREIGN KEY(user_id) references users(id),
    CONSTRAINT user_vendors_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
);

-- +migrate Down
DROP TABLE `users_vendors`;