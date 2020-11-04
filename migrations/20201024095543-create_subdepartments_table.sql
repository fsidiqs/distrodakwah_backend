-- +migrate Up
CREATE TABLE IF NOT EXISTS `subdepartments`(
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `department_id` INT UNSIGNED DEFAULT NULL,
    `image_id` BIGINT UNSIGNED DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT subdepartments_department_id foreign key(department_id) references departments(id),
    CONSTRAINT subdepartments_image_id foreign key(image_id) references images(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `subdepartments`;