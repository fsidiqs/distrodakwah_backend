-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_types` (
    `id` INT UNSIGNED AUTO_INCREMENT,
    `name` VARCHAR(99) NOT NULL,
    PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE `item_types`;