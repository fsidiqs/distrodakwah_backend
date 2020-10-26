
-- +migrate Up
CREATE TABLE IF NOT EXISTS sku_values(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `product_image_id` BIGINT UNSIGNED NOT NULL,
    `sku` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,

    CONSTRAINT sku_values_product_id FOREIGN KEY(product_id) references products(id)
    ON DELETE CASCADE,
    CONSTRAINT sku_values_product_image_id foreign key(product_image_id) references product_images(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +migrate Down
DROP TABLE `sku_values`;