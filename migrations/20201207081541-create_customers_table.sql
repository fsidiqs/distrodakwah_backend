-- +migrate Up
CREATE TABLE IF NOT EXISTS `customers`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `phone` VARCHAR(255) NOT NULL,
    `subdistrict_id` INT NOT NULL,
    `address` TEXT NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `user_reseller_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT customers_user_reseller_id FOREIGN KEY(user_reseller_id) references user_resellers(id),
    CONSTRAINT customers_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `customers`;