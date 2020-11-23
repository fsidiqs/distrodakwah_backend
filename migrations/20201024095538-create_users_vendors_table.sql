-- +migrate Up
CREATE TABLE IF NOT EXISTS `users_vendors`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `subdistrict_id` INT NOT NULL,
    `address` TEXT NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT users_vendors_user_id FOREIGN KEY(user_id) references users(id),
    CONSTRAINT usvs_subdistrict_id FOREIGN KEY(subdistrict_id) REFERENCES tb_ro_subdistricts(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `users_vendors`;