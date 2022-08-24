CREATE TABLE `tokens`
(
    `id`            bigint(15) unsigned NOT NULL AUTO_INCREMENT,
    `account_id`    bigint(15) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID',
    `access_token`  varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `refresh_token` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `expired_at`    datetime                                 NOT NULL COMMENT 'access_token 过期时间',
    `created_at`    datetime                                 NOT NULL COMMENT '添加时间',
    `updated_at`    timestamp                                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    `token_type`    varchar(20) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_account_id` (`account_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
