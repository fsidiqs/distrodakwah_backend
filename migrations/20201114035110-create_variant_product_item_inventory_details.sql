-- +migrate Up
CREATE TABLE IF NOT EXISTS `VPI_inventory_details` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VPI_inventory_id` BIGINT UNSIGNED NOT NULL,
    `subdistrict_id` INT NOT NULL,
    CONSTRAINT VPIIDs_VPI_inventory_id FOREIGN KEY(VPI_inventory_id) references VP_item_inventories(id) ON DELETE CASCADE,
    CONSTRAINT VPIIDs_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
);

-- +migrate Down
DROP TABLE `VPI_inventory_details`;