-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_status_lead_times` (
    `order_id` BIGINT UNSIGNED NOT NULL,
    `order_status_id` TINYINT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `order_status_lead_times`;