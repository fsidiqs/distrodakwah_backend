-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_items` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `VP_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT VP_items_VP_id FOREIGN KEY(VP_id) references variant_products(id)
);

-- +migrate Down
DROP TABLE `VP_items`;