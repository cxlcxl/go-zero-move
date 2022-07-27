CREATE TABLE `oauth_access_tokens`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(15) NOT NULL DEFAULT '0',
    `client_id`   int(10) unsigned NOT NULL DEFAULT '1' COMMENT '普通用户的授权，默认为1',
    `token`       varchar(2000) NOT NULL,
    `action_name` varchar(128)  NOT NULL DEFAULT '' COMMENT 'login|refresh|reset表示token生成动作',
    `scopes`      varchar(128)  NOT NULL DEFAULT '[*]' COMMENT '暂时预留,未启用',
    `revoked`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否撤销',
    `client_ip`   varchar(128)  NOT NULL COMMENT 'ipv6最长为128位',
    `created_at`  datetime      NOT NULL COMMENT '添加时间',
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    `expires_at`  datetime      NOT NULL COMMENT '过期时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `oauth_access_tokens_user_id_index` (`user_id`) USING BTREE,
    KEY           `idx_user_id_expires_at` (`user_id`,`expires_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `roles`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `role_name`  varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `state`      tinyint(2) NOT NULL DEFAULT '1' COMMENT '1正常0停用',
    `sys`        tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '角色所属系统',
    `created_at` datetime                               NOT NULL COMMENT '添加时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `user_roles`
(
    `id`      int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(10) unsigned NOT NULL DEFAULT '0',
    `role_id` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_id` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `users`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `email`      varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登陆邮箱',
    `username`   varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `mobile`     varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `state`      tinyint(2) NOT NULL DEFAULT '1' COMMENT '账号状态',
    `secret`     varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码加密符',
    `pass`       varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
    `created_at` datetime                               NOT NULL COMMENT '添加时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`),
    UNIQUE KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

