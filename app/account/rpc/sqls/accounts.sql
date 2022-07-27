create table accounts
(
    id           BIGINT(15) unsigned auto_increment not null,
    account_id   varchar(30) not null DEFAULT '' comment '账户ID',
    account_type tinyint(2) unsigned not null DEFAULT 1,
    email        varchar(50) not null DEFAULT '' comment '邮箱',
    account_name varchar(50) not null default '' comment '账户名',
    created_at   datetime    not null comment '添加时间',
    updated_at   TIMESTAMP   not null DEFAULT CURRENT_TIMESTAMP comment '最后一次修改时间',
    state        tinyint(2) unsigned not null DEFAULT 1,
    PRIMARY key (id),
    unique key idx_account_id(account_id) using btree
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table account_client_ids
(
    id         BIGINT(15) unsigned auto_increment not null,
    account_id varchar(30)  not null DEFAULT '' comment '账户ID',
    client_id  int(11) unsigned not null DEFAULT 0 comment '客户端ID',
    secret     varchar(70)  not null default '' comment '密钥',
    comment    varchar(100) not null default '' comment '备注',
    created_at datetime     not null comment '添加时间',
    updated_at TIMESTAMP    not null DEFAULT CURRENT_TIMESTAMP comment '最后一次修改时间',
    PRIMARY key (id),
    key        idx_account_id(account_id) using btree
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table tokens
(
    id            bigint(15) unsigned auto_increment not null,
    client_id     int(11) unsigned not null default 0 comment '客户端ID',
    access_token  varchar(1000) not null default '',
    refresh_token varchar(1000) not null default '',
    expired_at    datetime      not null comment 'access_token 过期时间',
    created_at    datetime      not null comment '添加时间',
    updated_at    TIMESTAMP     not null DEFAULT CURRENT_TIMESTAMP comment '最后一次修改时间',
    primary key (id),
    unique KEY idx_client_id (client_id) using btree
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
