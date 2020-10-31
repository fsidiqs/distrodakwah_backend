-- +migrate Up
CREATE TABLE IF NOT EXISTS `prices`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `prices`;