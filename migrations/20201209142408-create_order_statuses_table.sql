-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_statuses` (
    `id` TINYINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `order_statuses`;