-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_shippings` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `shipping_cost` BIGINT DEFAULT 0,
    `cost_total` BIGINT NOT NULL,
    `weight_total` INT NOT NULL,
    `shipping_company_id` INT UNSIGNED NOT NULL,
    `shipping_service_name` VARCHAR(255) NOT NULL,
    `item_origin_location_id` INT NOT NULL,
    `item_origin_location_type` VARCHAR(255) NOT NULL,
    `awb` VARCHAR(255) NOT NULL,
    `type_id` TINYINT UNSIGNED,
    CONSTRAINT order_shippings_order_id FOREIGN KEY(order_id) references orders(id),
    CONSTRAINT order_shippings_shipping_company_id FOREIGN KEY(shipping_company_id) references shipping_companies(id)
);

-- +migrate Down
DROP TABLE `order_shippings`;