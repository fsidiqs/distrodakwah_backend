-- +migrate Up
CREATE TABLE IF NOT EXISTS `departments`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    CONSTRAINT departments_image_id foreign key(image_id) references images(id)
);

-- +migrate Down
DROP TABLE `departments`;