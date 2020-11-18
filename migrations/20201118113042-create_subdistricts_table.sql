-- +migrate Up
CREATE TABLE IF NOT EXISTS `subdistricts`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `city_id` BIGINT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT subdistricts_city_id FOREIGN KEY(city_id) references cities(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `subdistricts`;