-- +migrate Up
CREATE TABLE IF NOT EXISTS `cities`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `type` VARCHAR(255) NOT NULL,
    `province_id` BIGINT UNSIGNED NOT NULL,
    `postal_code` VARCHAR(255) NOT NULL,
    CONSTRAINT cities_province_id FOREIGN KEY(province_id) references provinces(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `cities`;