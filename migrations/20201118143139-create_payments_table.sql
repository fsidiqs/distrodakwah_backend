-- +migrate Up
CREATE TABLE IF NOT EXISTS `payments`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `total_paid` DECIMAL(19, 2) NOT NULL,
    `image_url` VARCHAR(255),
    CONSTRAINT payments_order_id FOREIGN KEY(order_id) references orders(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `payments`;