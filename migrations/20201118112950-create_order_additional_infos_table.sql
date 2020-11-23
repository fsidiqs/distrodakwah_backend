-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_additional_infos`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `sender_name` VARCHAR(255),
    `sender_phone` VARCHAR(255),
    CONSTRAINT order_additional_infos_order_id FOREIGN KEY(order_id) references orders(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `order_additional_infos`;