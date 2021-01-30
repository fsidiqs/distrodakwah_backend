-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_variants`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VP_id` BIGINT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    
  CONSTRAINT VP_variants_VP_id FOREIGN KEY(VP_id) references variant_products(id)
);    

-- +migrate Down
DROP TABLE `VP_variants`;