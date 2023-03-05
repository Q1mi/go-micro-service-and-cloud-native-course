CREATE TABLE `order`(
                        `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
                        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                        `update_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新者',
                        `version` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
                        `is_del` tinyint(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',

                        `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户id',
                        `order_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '订单id',
                        `trade_id` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '交易单号',
                        `pay_channel` tinyint(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付方式',
                        `status` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单状态:100创建订单/待支付 200已支付 300交易关闭 400完成',
                        `pay_amount` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付金额（分）',
                        `pay_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '支付时间',

                        INDEX (user_id),
                        INDEX (order_id),
                        INDEX (trade_id),
                        INDEX (is_del)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '订单表';