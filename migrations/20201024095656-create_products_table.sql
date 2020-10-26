-- +migrate Up
CREATE TABLE IF NOT EXISTS products(
    `id` bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `product_image_id` BIGINT UNSIGNED DEFAULT NULL,
    `status` char(1) default 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT products_product_image_id foreign key(product_image_id) references product_images(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +migrate Down
DROP TABLE `products`;
