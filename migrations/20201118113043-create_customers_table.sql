-- +migrate Up
CREATE TABLE IF NOT EXISTS `customers` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `address_detail` TEXT NOT NULL,
    `subdistrict_id` int NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    CONSTRAINT customers_subdistrict_id FOREIGN KEY(subdistrict_id) references tb_ro_subdistricts(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `customers`;