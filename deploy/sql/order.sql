SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
      `id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
      `sn` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
      `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
      `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '产品id',
      `payment` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT '实际付款金额,单位是元,保留两位小数',
      `postage` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT '运费',
      `paymenttype` tinyint(4) NOT NULL DEFAULT 1 COMMENT '支付类型,1-在线支付',
      `status` smallint(6) NOT NULL DEFAULT 10 COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
      `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `idx_sn` (`sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

