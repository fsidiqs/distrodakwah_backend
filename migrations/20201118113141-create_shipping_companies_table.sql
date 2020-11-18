-- +migrate Up
CREATE TABLE IF NOT EXISTS `shipping_companies` (
    `id` TINYINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `shipping_companies`;