-- +migrate Up
CREATE TABLE IF NOT EXISTS `shipping_companies`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `shipping_companies`;