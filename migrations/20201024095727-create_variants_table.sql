-- +migrate Up
CREATE TABLE IF NOT EXISTS variants(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    
    CONSTRAINT variants_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE
) ENGINE = InnoDB;    

-- +migrate Down
DROP TABLE `variants`;