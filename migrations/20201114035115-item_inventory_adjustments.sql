-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_inventory_adjustments` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `item_inventory_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `stock_before` INT UNSIGNED NOT NULL,
    `stock_after` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL,
    CONSTRAINT item_inventory_adjustments_item_inventory_id FOREIGN KEY(item_inventory_id) references item_inventories(id),
    CONSTRAINT item_inventory_adjustments_user_id FOREIGN KEY(user_id) references users(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `item_inventory_adjustments`;