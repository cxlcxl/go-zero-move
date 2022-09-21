CREATE TABLE `apps`
(
    `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `account_id`             int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID;对应 accounts 表的id字段',
    `advertiser_id`          varchar(30)                                                   NOT NULL DEFAULT '' COMMENT '广告主账户ID',
    `app_id`                 varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
    `app_name`               varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用名称',
    `app_type`               varchar(50)                                                   NOT NULL DEFAULT '' COMMENT '产品/应用类型',
    `pkg_name`               varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '应用包名或BundleID',
    `channel`                tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '系统平台(渠道)：华为 AppGallery；GooglePlay; AppStore',
    `tags`                   varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '应用标签',
    `icon_url`               varchar(1000)                                                 NOT NULL DEFAULT '' COMMENT '图标',
    `product_id`             varchar(32)                                                   NOT NULL DEFAULT '' COMMENT '产品ID，创建任务时需要',
    `app_store_download_url` varchar(1000)                                                 NOT NULL DEFAULT '',
    `created_at`             datetime                                                      NOT NULL COMMENT '创建时间',
    `updated_at`             timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;