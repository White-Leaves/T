SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `orders_operation`;

CREATE TABLE `product_operation` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
     `product_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '商品id',
     `status` int NOT NULL DEFAULT '1' COMMENT '运营商品状态 0-下线 1-上线',
     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='商品运营表';
