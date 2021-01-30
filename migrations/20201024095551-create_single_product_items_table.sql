-- +migrate Up
CREATE TABLE IF NOT EXISTS `SP_items` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `SP_id` BIGINT UNSIGNED NOT NULL,
    `weight` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT SPItems_SP_id FOREIGN KEY(SP_id) references single_products(id)
);

-- +migrate Down
DROP TABLE `SP_items`;