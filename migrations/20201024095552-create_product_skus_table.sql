-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_skus`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    CONSTRAINT product_skus_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `product_skus`;