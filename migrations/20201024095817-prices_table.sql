-- +migrate Up
CREATE TABLE IF NOT EXISTS prices(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `sku_value_id` BIGINT UNSIGNED NOT NULL,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,

    CONSTRAINT prices_sku_value_id FOREIGN KEY(sku_value_id) REFERENCES sku_values(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down

DROP TABLE `prices`;