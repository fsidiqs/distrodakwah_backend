
-- +migrate Up
CREATE TABLE IF NOT EXISTs `product_images` (
  `id` bigint UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE product_images;