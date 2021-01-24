-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_variants`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    
    CONSTRAINT FK_VPVs_products_product_id FOREIGN KEY(product_id) references products(id) ON DELETE CASCADE
);    

-- +migrate Down
DROP TABLE `VP_variants`;