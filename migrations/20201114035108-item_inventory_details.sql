-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_inventory_details` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `item_inventory_id` BIGINT UNSIGNED NOT NULL,
    `subdistrict_id` INT NOT NULL,
    CONSTRAINT item_inventory_details_item_inventory_id FOREIGN KEY(item_inventory_id) references item_inventories(id) ON DELETE CASCADE,
    CONSTRAINT item_inventory_details_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `item_inventory_details`;