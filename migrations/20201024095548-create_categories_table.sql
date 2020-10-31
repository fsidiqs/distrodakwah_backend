-- +migrate Up
CREATE TABLE IF NOT EXISTS `categories`(
    `id` int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `parent_id` int UNSIGNED DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    CONSTRAINT categories_image_id foreign key(image_id) references images(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `categories`;