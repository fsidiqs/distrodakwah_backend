-- +migrate Up
CREATE TABLE IF NOT EXISTS `products` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `brand_id` INT UNSIGNED NOT NULL,
    `category_id` INT UNSIGNED NOT NULL,
    `product_character_id` TINYINT UNSIGNED NOT NULL,
    `product_type_id` INT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT products_brand_id FOREIGN KEY(brand_id) references brands(id),
    CONSTRAINT products_category_id FOREIGN KEY(category_id) references categories(id),
    CONSTRAINT products_product_character_id FOREIGN KEY(product_character_id) references product_characters(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `products`;