-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_inventory_details` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `item_inventory_id` BIGINT UNSIGNED NOT NULL,
    `vendor_id` INT UNSIGNED NOT NULL,
    CONSTRAINT item_inventory_details_item_inventory_id FOREIGN KEY(item_inventory_id) references item_inventories(id) ON DELETE CASCADE,
    CONSTRAINT item_inventory_details_vendor_id FOREIGN KEY(vendor_id) references user_vendors(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `item_inventory_details`;