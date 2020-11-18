-- +migrate Up
CREATE TABLE IF NOT EXISTS `provinces`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `provinces`;