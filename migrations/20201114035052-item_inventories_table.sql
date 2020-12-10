-- +migrate Up
CREATE TABLE IF NOT EXISTS `item_inventories` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `item_id` BIGINT UNSIGNED NOT NULL,
    `stock` INT NOT NULL,
    `keep` INT NOT NULL,
    CONSTRAINT item_inventories FOREIGN KEY(item_id) references items(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `item_inventories`;