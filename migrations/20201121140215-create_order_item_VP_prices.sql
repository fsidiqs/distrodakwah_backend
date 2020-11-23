
-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_item_VP_prices`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `oitvp_id` BIGINT UNSIGNED NOT NULL,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT order_item_VP_prices_oitvp_id FOREIGN KEY(oitvp_id) references order_item_variant_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `order_item_SP_prices`;
