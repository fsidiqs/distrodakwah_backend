-- +migrate Up
CREATE TABLE IF NOT EXISTS `vp_inventory_details` (
    `vp_inventory_id` BIGINT UNSIGNED NOT NULL,
    `vendor_id` INT UNSIGNED NOT NULL,
    CONSTRAINT vp_inventory_details_vp_inventory_id FOREIGN KEY(vp_inventory_id) references vp_inventory(id) ON DELETE CASCADE,
    CONSTRAINT vp_inventory_details_vendor_id FOREIGN KEY(vendor_id) references users_vendors(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `vp_inventory_details`;