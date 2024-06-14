SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for homestay
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `del_state` tinyint NOT NULL DEFAULT '0',
                            `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                            `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
                            `info` varchar(4069) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '介绍',
                            `row_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:下架 1:上架',


                            `price` bigint NOT NULL DEFAULT '0' COMMENT '民宿价格（分）',

                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品';

