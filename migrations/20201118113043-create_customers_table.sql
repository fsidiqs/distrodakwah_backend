-- +migrate Up
CREATE TABLE IF NOT EXISTS `customers` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `address_detail` TEXT NOT NULL,
    `subdistrict_id` BIGINT UNSIGNED NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    CONSTRAINT customers_subdistrict_id FOREIGN KEY(subdistrict_id) references subdistricts(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `customers`;