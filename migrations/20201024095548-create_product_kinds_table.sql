-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_kinds`(
    `id` TINYINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE `product_kinds`;