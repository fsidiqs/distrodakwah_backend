-- +migrate Up
CREATE TABLE IF NOT EXISTS `single_products_prices`(
    `single_product_id` BIGINT UNSIGNED NOT NULL,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT single_products_prices_single_product_id FOREIGN KEY(single_product_id) references single_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `single_products_prices`;