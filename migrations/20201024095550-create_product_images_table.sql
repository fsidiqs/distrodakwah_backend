-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_images` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  CONSTRAINT product_images_product_id FOREIGN KEY(product_id) references products(id)
);

-- +migrate Down
DROP TABLE `product_images`;