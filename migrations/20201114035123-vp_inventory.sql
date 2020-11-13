-- +migrate Up
CREATE TABLE IF NOT EXISTS `vp_inventory` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `variant_product_id` BIGINT UNSIGNED NOT NULL,
    `stock` INT NOT NULL,
    `keep` INT NOT NULL,
    CONSTRAINT vp_inventory_variant_product_id FOREIGN KEY(variant_product_id) references variant_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `vp_inventory`;