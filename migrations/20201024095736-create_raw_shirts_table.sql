-- +migrate Up
CREATE TABLE IF NOT EXISTS `raw_shirts`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `shirt_type` VARCHAR(9) NOT NULL,
    `size` VARCHAR(9) NOT NULL,
    `color` VARCHAR(100) NOT NULL,
    `weight` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE `raw_shirts`;