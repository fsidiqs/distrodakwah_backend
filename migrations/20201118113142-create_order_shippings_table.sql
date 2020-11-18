-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_shippings` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `shipping_cost` DECIMAL(19, 2) NOT NULL,
    `total_cost` DECIMAL(19, 2) NOT NULL,
    `weight` INT NOT NULL,
    `shipping_company_id` TINYINT UNSIGNED NOT NULL,
    `shipping_service_name` VARCHAR(255) NOT NULL,
    `subdistrict_id_origin` BIGINT UNSIGNED NOT NULL,
    `awb` VARCHAR(255) NOT NULL,
    `type_id` TINYINT UNSIGNED NOT NULL,
    CONSTRAINT customers_order_id FOREIGN KEY(order_id) references orders(id),
    CONSTRAINT customers_shipping_company_id FOREIGN KEY(shipping_company_id) references shipping_companies(id),
    CONSTRAINT customers_subdistrict_id_origin FOREIGN KEY(subdistrict_id_origin) references subdistricts(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `order_shippings`;