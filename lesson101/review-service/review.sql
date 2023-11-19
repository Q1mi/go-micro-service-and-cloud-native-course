CREATE TABLE review_info (
        `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
        `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT '创建方标识',
        `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT '更新方标识',
        `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `delete_at` timestamp COMMENT '逻辑删除标记',
        `version`   int(10) unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁标记',

        `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '评价id',
        `content` varchar(512) NOT NULL COMMENT '评价内容',
        `score` tinyint(4) NOT NULL DEFAULT '0' COMMENT '评分',
        `service_score` tinyint(4) NOT NULL DEFAULT '0' COMMENT '商家服务评分',
        `express_score` tinyint(4) NOT NULL DEFAULT '0' COMMENT '物流评分',
        `has_media` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有图或视频',
        `order_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '订单id',
        `sku_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 'sku id',
        `spu_id` bigint(32) NOT NULL DEFAULT '0' COMMENT 'spu id',
        `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '店铺id',
        `user_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '用户id',
        `anonymous` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否匿名',
        `tags` varchar(1024) NOT NULL DEFAULT '' COMMENT '标签json',
        `pic_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：图片',
        `video_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：视频',
        `status` tinyint(4) NOT NULL DEFAULT '10' COMMENT '状态:10待审核；20审核通过；30审核不通过；40隐藏',
        `is_default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认评价',
        `has_reply` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有商家回复:0无;1有',
        `op_reason` varchar(512) NOT NULL DEFAULT '' COMMENT '运营审核拒绝原因',
        `op_remarks` varchar(512) NOT NULL DEFAULT '' COMMENT '运营备注',
        `op_user` varchar(64) NOT NULL DEFAULT '' COMMENT '运营者标识',

        `goods_snapshoot` varchar(2048) NOT NULL DEFAULT '' COMMENT '商品快照信息',
        `ext_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '信息扩展',
        `ctrl_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '控制扩展',
        PRIMARY KEY (`id`),
        KEY `idx_delete_at` (`delete_at`) COMMENT '逻辑删除索引',
        UNIQUE KEY `uk_review_id` (`review_id`) COMMENT '评价id索引',
        KEY `idx_order_id` (`order_id`) COMMENT '订单id索引',
        KEY `idx_user_id` (`user_id`) COMMENT '用户id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价表';

CREATE TABLE review_reply_info (
        `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
        `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT '创建方标识',
        `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT '更新方标识',
        `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `delete_at` timestamp COMMENT '逻辑删除标记',
        `version` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁标记',

        `reply_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '回复id',
        `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '评价id',
        `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '店铺id',
        `content` varchar(512) NOT NULL COMMENT '评价内容',
        `pic_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：图片',
        `video_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：视频',

        `ext_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '信息扩展',
        `ctrl_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '控制扩展',
        PRIMARY KEY (`id`),
        KEY `idx_delete_at` (`delete_at`) COMMENT '逻辑删除索引',
        UNIQUE KEY `uk_reply_id` (`reply_id`) COMMENT '回复id索引',
        KEY `idx_review_id` (`review_id`) COMMENT '评价id索引',
        KEY `idx_store_id` (`store_id`) COMMENT '店铺id索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价商家回复表';


CREATE TABLE review_appeal_info (
        `id` bigint(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
        `create_by` varchar(48) NOT NULL DEFAULT '' COMMENT '创建方标识',
        `update_by` varchar(48) NOT NULL DEFAULT '' COMMENT '更新方标识',
        `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `delete_at` timestamp COMMENT '逻辑删除标记',
        `version` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁标记',

        `appeal_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '回复id',
        `review_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '评价id',
        `store_id` bigint(32) NOT NULL DEFAULT '0' COMMENT '店铺id',
        `status` tinyint(4) NOT NULL DEFAULT '10' COMMENT '状态:10待审核；20申诉通过；30申诉驳回',
        `reason` varchar(255) NOT NULL COMMENT '申诉原因类别',
        `content` varchar(255) NOT NULL COMMENT '申诉内容描述',
        `pic_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：图片',
        `video_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '媒体信息：视频',

        `op_remarks` varchar(512) NOT NULL DEFAULT '' COMMENT '运营备注',
        `op_user` varchar(64) NOT NULL DEFAULT '' COMMENT '运营者标识',

        `ext_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '信息扩展',
        `ctrl_json` varchar(1024) NOT NULL DEFAULT '' COMMENT '控制扩展',
        PRIMARY KEY (`id`),
        KEY `idx_delete_at` (`delete_at`) COMMENT '逻辑删除索引',
        KEY `idx_appeal_id` (`appeal_id`) COMMENT '申诉id索引',
        UNIQUE KEY `uk_review_id` (`review_id`) COMMENT '评价id索引',
        KEY `idx_store_id` (`store_id`) COMMENT '店铺id索引'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评价商家申诉表';