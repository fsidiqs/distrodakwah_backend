-- +migrate Up
CREATE TABLE IF NOT EXISTS `VP_options`(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `VP_variant_id` BIGINT UNSIGNED NOT NULL,
    `VP_item_id` BIGINT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    CONSTRAINT FK_VPOs_VPs_VP_id FOREIGN KEY(VP_variant_id) REFERENCES VP_variants(id) ON DELETE CASCADE,
    CONSTRAINT FK_VPOs_VPIs_VP_item_id FOREIGN KEY(VP_item_id) REFERENCES VP_items(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE `VP_options`;