-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_item_inventories` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VP_item_id` BIGINT UNSIGNED NOT NULL,
    `stock` INT NOT NULL,
    `keep` INT NOT NULL
);

-- +migrate Down
DROP TABLE `VP_item_inventories`;