-- +migrate Up
CREATE TABLE IF NOT EXISTS `items` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT items_product_id FOREIGN KEY(product_id) references products(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `items`;