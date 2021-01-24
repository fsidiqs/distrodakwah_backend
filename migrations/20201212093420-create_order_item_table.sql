-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_items` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `itemable_id` BIGINT UNSIGNED NOT NULL,
    `qty` INT NOT NULL,
    `unit_weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    `order_shipping_id` BIGINT UNSIGNED,
    CONSTRAINT order_items_order_id FOREIGN KEY(order_id) references orders(id),
    CONSTRAINT order_items_order_shipping_id FOREIGN KEY(order_shipping_id) references order_shippings(id)
);

-- +migrate Down
DROP TABLE `order_items`;