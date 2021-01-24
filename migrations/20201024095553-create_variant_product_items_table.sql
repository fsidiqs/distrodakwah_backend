-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_items` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT FK_VPIs_products_product_id FOREIGN KEY(product_id) references products(id)
);

-- +migrate Down
DROP TABLE `VP_items`;