-- +migrate Up
CREATE TABLE IF NOT EXISTS `variant_product_images` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `VP_id` BIGINT UNSIGNED NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  CONSTRAINT VP_images_VP_id FOREIGN KEY(VP_id) references variant_products(id)
);

-- +migrate Down
DROP TABLE `variant_product_images`;