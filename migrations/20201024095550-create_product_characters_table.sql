-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_characters`(
    `id` TINYINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `product_characters`;