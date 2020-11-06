-- +migrate Up
CREATE TABLE IF NOT EXISTS `single_products`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    CONSTRAINT single_products_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `single_products`;