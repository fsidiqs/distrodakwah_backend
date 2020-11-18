-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_items_single_products`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `single_product_id` BIGINT UNSIGNED NOT NULL,
    `order_shipping_id` BIGINT UNSIGNED NOT NULL,
    `qty` INT NOT NULL,
    `unit_weight` INT NOT NULL,
    `dropshipper_item_price` DECIMAL(19, 2) NOT NULL,
    `retail_item_price` DECIMAL(19, 2) NOT NULL,
    CONSTRAINT order_items_single_products_order_id FOREIGN KEY(order_id) references orders(id),
    CONSTRAINT oisp_single_product_id FOREIGN KEY(single_product_id) references single_products(id),
    CONSTRAINT oisp_order_shipping_id FOREIGN KEY(order_shipping_id) references order_shippings(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `order_items_single_products`;