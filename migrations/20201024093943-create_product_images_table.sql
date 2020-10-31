-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_images` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `url` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `product_images`;