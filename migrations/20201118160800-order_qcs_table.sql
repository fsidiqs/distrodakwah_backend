-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_qcs` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` BIGINT UNSIGNED NOT NULL,
    `bahan` BOOLEAN NOT NULL,
    `desain` BOOLEAN NOT NULL,
    `qc` BOOLEAN NOT NULL,
    `packing` BOOLEAN NOT NULL,
    `pickup` BOOLEAN NOT NULL,
    `jurnal` BOOLEAN NOT NULL,
    CONSTRAINT order_qcs_order_id FOREIGN KEY(order_id) references orders(id)
) ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `order_qcs`;