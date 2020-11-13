-- +migrate Up
CREATE TABLE IF NOT EXISTS `sp_inventory_details` (
    `sp_inventory_id` BIGINT UNSIGNED NOT NULL,
    `vendor_id` INT UNSIGNED NOT NULL,
    CONSTRAINT sp_inventory_details_sp_inventory_id FOREIGN KEY(sp_inventory_id) references sp_inventory(id) ON DELETE CASCADE,
    CONSTRAINT sp_inventory_details_vendor_id FOREIGN KEY(vendor_id) references users_vendors(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `sp_inventory_details`;