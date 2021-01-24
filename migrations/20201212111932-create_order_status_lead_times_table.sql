-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_status_lead_times`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `status_id` TINYINT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NULL
);

-- +migrate Down
DROP TABLE `order_status_lead_times`;