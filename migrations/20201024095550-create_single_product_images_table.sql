-- +migrate Up
CREATE TABLE IF NOT EXISTS `single_product_images` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `SP_id` BIGINT UNSIGNED NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  CONSTRAINT SP_images_SP_id FOREIGN KEY(SP_id) references single_products(id)
);

-- +migrate Down
DROP TABLE `single_product_images`;