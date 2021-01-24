-- +migrate Up
CREATE TABLE IF NOT EXISTS `products` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `brand_id` INT UNSIGNED NOT NULL,
    `category_id` INT UNSIGNED NOT NULL,
    `product_kind_id` TINYINT UNSIGNED NOT NULL,
    `product_type_id` INT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `status` VARCHAR(10) NOT NULL DEFAULT "A",
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT products_brand_id FOREIGN KEY(brand_id) references brands(id),
    CONSTRAINT products_category_id FOREIGN KEY(category_id) references categories(id),
    CONSTRAINT products_product_kind_id FOREIGN KEY(product_kind_id) references product_kinds(id)
);

-- +migrate Down
DROP TABLE `products`;