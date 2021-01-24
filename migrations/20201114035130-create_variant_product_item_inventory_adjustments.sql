-- +migrate Up
CREATE TABLE IF NOT EXISTS `VPI_inventory_adjustments` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VPI_inventory_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `stock_before` INT UNSIGNED NOT NULL,
    `stock_after` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL,
    CONSTRAINT VPI_inventory_adjustments_VPI_inventory_id FOREIGN KEY(VPI_inventory_id) references VP_item_inventories(id),
    CONSTRAINT VPI_inventory_adjustments_user_id FOREIGN KEY(user_id) references users(id)
);

-- +migrate Down
DROP TABLE `VPI_inventory_adjustments`;