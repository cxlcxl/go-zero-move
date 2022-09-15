CREATE TABLE `campaigns`
(
    `id`                           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `campaign_id`                  varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '计划 ID',
    `campaign_name`                varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '计划名称',
    `account_id`                   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户 ID',
    `advertiser_id`                varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '广告主账户 ID',
    `opt_status`                   varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '操作状态',
    `campaign_daily_budget_status` varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '计划日预算状态',
    `product_type`                 varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '推广产品类型',
    `show_status`                  varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '计划状态',
    `user_balance_status`          varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '账户余额状态',
    `flow_resource`                varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '投放网络',
    `sync_flow_resource`           varchar(5) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '' COMMENT '同时同步投放搜索广告网络',
    `campaign_type`                varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '计划类型',
    `today_daily_budget`           int(11) unsigned NOT NULL DEFAULT '0' COMMENT '当日计划日限额',
    `tomorrow_daily_budget`        int(11) unsigned NOT NULL DEFAULT '0' COMMENT '次日计划日限额，不返回表示与当日计划日限额相同',
    `marketing_goal`               varchar(35) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '营销目标',
    `is_callback`                  tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否通过查询计划任务回调完整信息',
    `created_at`                   datetime                                NOT NULL COMMENT '添加时间',
    `updated_at`                   timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_campaign_id` (`campaign_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;