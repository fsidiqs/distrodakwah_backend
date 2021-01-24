-- +migrate Up
CREATE TABLE IF NOT EXISTS `SP_item_inventories` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `SP_item_id` BIGINT UNSIGNED NOT NULL,
    `stock` INT NOT NULL,
    `keep` INT NOT NULL
);

-- +migrate Down
DROP TABLE `SP_item_inventories`;