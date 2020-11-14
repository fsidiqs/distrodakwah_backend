-- +migrate Up
CREATE TABLE IF NOT EXISTS `sp_inventory_adjustments` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `sp_inventory_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `stock_before` INT UNSIGNED NOT NULL,
    `stock_after` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL,
    CONSTRAINT sp_inventory_adjustments_sp_inventory_id FOREIGN KEY(sp_inventory_id) references sp_inventory(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `sp_inventory_adjustments`;