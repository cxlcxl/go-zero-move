CREATE TABLE `jobs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '数据日期',
  `api_module` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '数据模块',
  `total_page` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总页数',
  `total_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总数',
  `job_schedule` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'cron 调度',
  `pause_rule` int(3) unsigned NOT NULL DEFAULT '0' COMMENT '调度截止规则：0 调度到当天；-1 停止调度此任务；> 0 为当前日期减{pause_rule}天',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_module` (`api_module`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;