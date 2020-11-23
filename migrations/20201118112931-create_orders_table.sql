-- +migrate Up
CREATE TABLE IF NOT EXISTS `orders` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `invoice` VARCHAR(255) NOT NULL,
    `order_status_id` TINYINT UNSIGNED NOT NULL,
    `total` DECIMAL(19, 2) DEFAULT 0.0 NOT NULL,
    `shipping_cost` DECIMAL(19, 2) DEFAULT 0.0 NOT NULL,
    `grand_total` DECIMAL(19, 2) DEFAULT 0.0 NOT NULL,
    `unique_code` DECIMAL(5, 2) DEFAULT 0.0 NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL,
    `status_id_1_expires` TIMESTAMP NULL,
    CONSTRAINT orders_user_id FOREIGN KEY(user_id) references users(id),
    CONSTRAINT orders_order_status_id FOREIGN KEY(order_status_id) references order_statuses(id)
) Engine = InnoDB;

-- +migrate Down
DROP TABLE `orders`;