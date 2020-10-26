
-- +migrate Up
CREATE TABLE IF NOT EXISTS options(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `variant_id` BIGINT UNSIGNED NOT NULL,
    `sku_value_id` BIGINT UNSIGNED NOT NULL,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,

    CONSTRAINT options_variant_id FOREIGN KEY(variant_id) REFERENCES variants(id)
    ON DELETE CASCADE,
    CONSTRAINT options_sku_value_id FOREIGN KEY(sku_value_id) REFERENCES sku_values(id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +migrate Down
DROP TABLE `options`;