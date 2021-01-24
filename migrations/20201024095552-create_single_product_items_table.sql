-- +migrate Up
CREATE TABLE IF NOT EXISTS `SP_items` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_SP_items_products_product_id FOREIGN KEY(product_id) references products(id)
);

-- +migrate Down
DROP TABLE `SP_items`;