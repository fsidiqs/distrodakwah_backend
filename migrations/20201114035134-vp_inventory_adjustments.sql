-- +migrate Up
CREATE TABLE IF NOT EXISTS `vp_inventory_adjustments` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `vp_inventory_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `stock_before` INT UNSIGNED NOT NULL,
    `stock_after` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL,
    CONSTRAINT vp_inventory_adjustments_vp_inventory_id FOREIGN KEY(vp_inventory_id) references vp_inventory(id),
    CONSTRAINT vp_inventory_adjustments_user_id FOREIGN KEY(user_id) references users(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `vp_inventory_adjustments`;