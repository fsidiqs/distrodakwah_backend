-- +migrate Up
CREATE TABLE IF NOT EXISTS `brands` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_vendor_id` INT UNSIGNED DEFAULT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT brands_user_vendor_id foreign key(user_vendor_id) references user_vendors(id),
    CONSTRAINT brands_image_id foreign key(image_id) references images(id)
);

-- +migrate Down
DROP TABLE `brands`;