---- create ----
CREATE TABLE IF NOT EXISTS `products`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(100) NOT NULL UNIQUE,
 `detail`           VARCHAR(10000) NOT NULL,
 `price`              INT NOT NULL,
 `img`	            MEDIUMBLOB NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;