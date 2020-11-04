-- +migrate Up
CREATE TABLE IF NOT EXISTS `single_products_prices`(
    `product_id` BIGINT UNSIGNED NOT NULL,
    `price_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT single_products_prices_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE,
    CONSTRAINT single_products_prices_price_id FOREIGN KEY(price_id) references prices(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `single_products_prices`;