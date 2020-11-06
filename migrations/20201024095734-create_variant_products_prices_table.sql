-- +migrate Up
CREATE TABLE IF NOT EXISTS `variant_products_prices`(
    `variant_product_id` BIGINT UNSIGNED NOT NULL,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT variant_products_prices_variant_product_id FOREIGN KEY(variant_product_id) references variant_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `variant_products_prices`;