SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `payinfo`;

CREATE TABLE `payinfo` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付信息表id',
   `sn` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '流水单号',
   `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
   `order_sn` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '业务单号',
   `platform` tinyint(4) NOT NULL DEFAULT 0 COMMENT '支付平台:1-支付宝,2-微信',
   `pay_total` decimal(20,2) NOT NULL DEFAULT '0' COMMENT '支付总金额',

   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   `pay_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '平台内交易状态   -1:支付失败 0:未支付 1:支付成功 2:已退款',
   `pay_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:00' COMMENT '支付成功时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `idx_sn` (`sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付信息表';
