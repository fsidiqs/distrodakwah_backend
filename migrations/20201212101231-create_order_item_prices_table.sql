-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_item_prices` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_item_id` BIGINT UNSIGNED NOT NULL,
    `value` BIGINT DEFAULT 0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT order_item_prices_order_item_id FOREIGN KEY(order_item_id) references order_items(id) ON DELETE CASCADE
) Engine = InnoDB;

-- +migrate Down

DROP TABLE `order_item_prices`;