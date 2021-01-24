-- +migrate Up
CREATE TABLE IF NOT EXISTS `SPI_inventory_adjustments` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `SPI_inventory_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `stock_before` INT UNSIGNED NOT NULL,
    `stock_after` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL,
    CONSTRAINT SPI_inventory_adjustments_SPI_inventory_id FOREIGN KEY(SPI_inventory_id) references SP_item_inventories(id),
    CONSTRAINT SPI_inventory_adjustments_user_id FOREIGN KEY(user_id) references users(id)
);

-- +migrate Down
DROP TABLE `SPI_inventory_adjustments`;