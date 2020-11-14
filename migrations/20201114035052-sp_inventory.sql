-- +migrate Up
CREATE TABLE IF NOT EXISTS `sp_inventory` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `single_product_id` BIGINT UNSIGNED NOT NULL,
    `stock` INT NOT NULL,
    `keep` INT NOT NULL,
    CONSTRAINT sp_inventory_single_product_id FOREIGN KEY(single_product_id) references single_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `sp_inventory`;