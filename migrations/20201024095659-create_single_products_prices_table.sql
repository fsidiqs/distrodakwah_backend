-- +migrate Up
CREATE TABLE IF NOT EXISTS `single_products_prices`(
    `single_product_id` BIGINT UNSIGNED NOT NULL,
    `price_id` BIGINT UNSIGNED NOT NULL,
    CONSTRAINT single_products_prices_single_product_id FOREIGN KEY(single_product_id) references single_products(id) ON DELETE CASCADE,
    CONSTRAINT single_products_prices_price_id FOREIGN KEY(price_id) references prices(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `single_products_prices`;