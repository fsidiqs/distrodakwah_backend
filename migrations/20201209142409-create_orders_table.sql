-- +migrate Up
CREATE TABLE IF NOT EXISTS `orders` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `invoice` STRING NOT NULL,
    `order_status_id` TINYINT UNSIGNED NOT NULL,
    `total` BIGINT NOT NULL DEFAULT 0,
    `unique_code` INT NOT NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    `user_reseller_id` BIGINT UNSIGNED NOT NULL,
    `order_status_id_1_expires` TIMESTAMP NOT NULL,
    CONSTRAINT orders_user_reseller_id FOREIGN KEY(user_reseller_id) references user_resellers(id),
    CONSTRAINT orders_order_status_id FOREIGN KEY(order_status_id) references order_statuses(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `orders`;