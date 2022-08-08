CREATE TABLE `apps`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID;对应 accounts 表的id字段',
    `app_id`     varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
    `app_name`   varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用名称',
    `pkg_name`   varchar(128)                                                  NOT NULL DEFAULT '' COMMENT '应用包名或BundleID',
    `channel`    tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '系统平台(渠道)：华为 AppGallery；GooglePlay; AppStore',
    `app_type`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '应用分类',
    `tags`       varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '应用标签',
    `created_at` datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;