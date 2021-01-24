-- +migrate Up
CREATE TABLE IF NOT EXISTS `product_types`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `value` INT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE `product_types`;