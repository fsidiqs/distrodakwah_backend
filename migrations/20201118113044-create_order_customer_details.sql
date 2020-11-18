-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_customer_details`(
    `order_id` BIGINT UNSIGNED NOT NULL,
    `customer_id` BIGINT UNSIGNED NOT NULL,
    `address_detail` TEXT NOT NULL,
    `subdistrict_id` BIGINT UNSIGNED NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `order_customer_details`;