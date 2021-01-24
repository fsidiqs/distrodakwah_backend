-- +migrate Up
CREATE TABLE IF NOT EXISTS `SPI_inventory_details` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `SPI_inventory_id` BIGINT UNSIGNED NOT NULL,
    `subdistrict_id` INT NOT NULL,
    CONSTRAINT SPIIDs_SPI_inventory_id FOREIGN KEY(SPI_inventory_id) references SP_item_inventories(id) ON DELETE CASCADE,
    CONSTRAINT SPIIDs_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
);

-- +migrate Down
DROP TABLE `SPI_inventory_details`;