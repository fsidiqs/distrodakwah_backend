-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_prices`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `item_id` BIGINT UNSIGNED NOT NULL,
    `value` BIGINT DEFAULT 0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT item_prices_item_id FOREIGN KEY(item_id) references items(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `item_prices`;