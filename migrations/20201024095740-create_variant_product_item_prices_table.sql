-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_item_prices`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VP_item_id` BIGINT UNSIGNED NOT NULL,
    `value` BIGINT DEFAULT 0,
    `name` VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE `VP_item_prices`;