
-- +migrate Up
CREATE TABLE IF NOT EXISTS `options`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `variant_id` BIGINT UNSIGNED NOT NULL,
    `variant_product_id` BIGINT UNSIGNED NOT NULL,
    `name`VARCHAR(255) NOT NULL,

    CONSTRAINT options_variant_id FOREIGN KEY(variant_id) REFERENCES variants(id)
    ON DELETE CASCADE,
    CONSTRAINT options_variant_product_id FOREIGN KEY(variant_product_id) REFERENCES variant_products(id)
    ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `options`;