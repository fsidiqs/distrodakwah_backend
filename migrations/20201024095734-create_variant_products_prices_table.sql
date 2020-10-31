-- +migrate Up
CREATE TABLE IF NOT EXISTS `variant_products_prices`(
    `variant_product_id` BIGINT UNSIGNED NOT NULL,
    `price_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT variant_products_prices_variant_product_id FOREIGN KEY(variant_product_id) references variant_products(id) ON DELETE CASCADE,
    CONSTRAINT variant_products_prices_price_id FOREIGN KEY(price_id) references prices(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `variant_products_prices`;