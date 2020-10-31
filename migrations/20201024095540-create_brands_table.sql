-- +migrate Up
CREATE TABLE IF NOT EXISTS `brands` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT brands_image_id foreign key(image_id) references images(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `brands`;