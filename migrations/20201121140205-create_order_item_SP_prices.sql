
-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_item_SP_prices`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `oitsp_id` BIGINT UNSIGNED NOT NULL,
    `value` DECIMAL(19, 2) DEFAULT 0.0,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT order_item_SP_prices_oitsp_id FOREIGN KEY(oitsp_id) references order_item_single_products(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `order_item_SP_prices`;
