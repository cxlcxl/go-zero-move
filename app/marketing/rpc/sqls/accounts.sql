CREATE TABLE `accounts`
(
    `id`            bigint(15) unsigned NOT NULL AUTO_INCREMENT,
    `parent_id`     bigint(15) unsigned NOT NULL DEFAULT '0' COMMENT '所属上级服务商',
    `advertiser_id` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '广告主账户ID',
    `developer_id`  varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '开发者ID',
    `account_type`  tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '账户类型',
    `state`         tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
    `account_name`  varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账户名',
    `client_id`     varchar(15)                            NOT NULL DEFAULT '' COMMENT '客户端ID',
    `is_auth`       tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否已认证',
    `secret`        varchar(70) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密钥',
    `created_at`    datetime                               NOT NULL COMMENT '添加时间',
    `updated_at`    timestamp                              NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_advertiser_id` (`advertiser_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table tokens
(
    id            bigint(15) unsigned auto_increment not null,
    client_id     varchar(15)                            not null default '' comment '客户端ID',
    access_token  varchar(1000)                          not null default '',
    refresh_token varchar(1000)                          not null default '',
    expired_at    datetime                               not null comment 'access_token 过期时间',
    created_at    datetime                               not null comment '添加时间',
    updated_at    TIMESTAMP                              not null DEFAULT CURRENT_TIMESTAMP comment '最后一次修改时间',
    token_type    varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    primary key (id),
    unique KEY idx_client_id (client_id) using btree
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;