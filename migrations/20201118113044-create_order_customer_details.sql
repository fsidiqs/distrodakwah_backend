-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_customer_details`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `customer_id` BIGINT UNSIGNED NOT NULL,
    `address_detail` TEXT NOT NULL,
    `subdistrict_id` INT NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    CONSTRAINT ocds_order_id FOREIGN KEY(order_id) references orders(id),
    CONSTRAINT ocds_customer_id FOREIGN KEY(customer_id) references customers(id),
    CONSTRAINT ocds_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `order_customer_details`;