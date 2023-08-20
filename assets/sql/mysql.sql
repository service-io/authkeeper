drop table if exists `account`;
create table `account`
(
    `id`                    bigint       not null comment '主键' auto_increment,
    `name`                  varchar(100) not null default '' comment '名称',
    `password`              varchar(100) null     default '' comment '密码',
    `password_salt`         varchar(100) null     default '' comment '盐值',
    `encrypt_type`          varchar(100) null     default '' comment '加密类型',
    `nick_name`             varchar(100) null     default '' comment '昵称',
    `display_name`          varchar(100) null     default '' comment '显示名称',
    `first_name`            varchar(100) null     default '' comment '名字',
    `last_name`             varchar(100) null     default '' comment '姓氏',
    `avatar`                varchar(500) null     default '' comment '头像',
    `avatar_type`           varchar(100) null     default '' comment '头像类型',
    `email`                 varchar(100) null     default '' comment '邮箱',
    `email_verified`        tinyint      null     default 0 comment '邮箱是否验证',
    `tel`                   varchar(20)  null     default '' comment '电话',
    `tel_verified`          tinyint      null     default 0 comment '电话是否验证',
    `country_code`          varchar(6)   null     default 'CN' comment '国家代码',
    `region`                varchar(100) null     default '' comment '区域',
    `city`                  varchar(100) null     default '' comment '城市',
    `location`              varchar(100) null     default '' comment '地理位置',
    `address`               varchar(500) null     default '' comment '地址',
    `id_card`               varchar(100) null     default '' comment '身份证',
    `id_card_type`          tinyint      null     default 0 comment '身份证类型',
    `homepage`              varchar(100) null     default '' comment '主页',
    `self_intro`            varchar(200) null     default '' comment '自我介绍',
    `language`              varchar(100) null     default '' comment '语言',
    `gender`                tinyint      null     default 0 comment '性别',
    `birthday`              varchar(30)  null     default '' comment '生日',
    `education`             varchar(100) null     default '' comment '教育',
    `last_sign_in_at`       timestamp    null     default current_timestamp comment '最后一次登录时间',
    `last_sign_in_ip`       varchar(50)  null     default '' comment '最后一次登录IP',
    `last_sign_in_wrong_at` timestamp    null     default current_timestamp comment '最后一次登录失败时间',
    `sign_in_wrong_times`   int          null     default 0 comment '登录失败次数',
    `is_admin`              tinyint      null     default 0 comment '是否组织管理员',
    `is_global_admin`       tinyint      null     default 0 comment '是否超级管理员',
    `is_forbidden`          tinyint      null     default 0 comment '是否禁用',
    `create_at`             timestamp    null     default current_timestamp comment '创建时间',
    `create_by`             bigint       null     default 0 comment '创建人',
    `modify_at`             timestamp    null     default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by`             bigint       null     default 0 comment '更新人',
    `deleted`               tinyint      null     default 0 comment '逻辑删除',
    primary key (`id`) using btree,
    unique key (`name`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4
  collate utf8mb4_bin comment ='账号表';

drop table if exists `tag`;
create table `tag`
(
    `id`        bigint       not null comment '主键' auto_increment,
    `name`      varchar(100) null default '' comment '标签名称',
    `status`    tinyint      null default 0 comment '状态, 0: 禁用, 1: 启用',
    `create_at` timestamp    null default current_timestamp comment '创建时间',
    `create_by` bigint       null default 0 comment '创建人',
    `modify_at` timestamp    null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint       null default 0 comment '更新人',
    `deleted`   tinyint      null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='标签表';

drop table if exists `type`;
create table `type`
(
    `id`        bigint       not null comment '主键' auto_increment,
    `name`      varchar(100) null default '' comment '类型名称',
    `code`      varchar(100) null default '' comment '类型编码',
    `status`    tinyint      null default 0 comment '状态, 0: 禁用, 1: 启用',
    `create_at` timestamp    null default current_timestamp comment '创建时间',
    `create_by` bigint       null default 0 comment '创建人',
    `modify_at` timestamp    null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint       null default 0 comment '更新人',
    `deleted`   tinyint      null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='类型表';

drop table if exists `application`;
create table `application`
(
    `id`            bigint       not null comment '主键' auto_increment,
    `name`          varchar(100) null default '' comment '应用名称',
    `logo`          varchar(200) null default '' comment '应用LOGO',
    `homepage`      varchar(100) null default '' comment '主页',
    `org_id`        bigint       null default 0 comment '组织ID',
    `cert`          varchar(500) null default '' comment '证书',
    `description`   varchar(200) null default '' comment '描述',
    `client_id`     varchar(200) null default '' comment '客户端ID',
    `client_secret` varchar(200) null default '' comment '客户端秘钥',
    `redirect_uris` varchar(200) null default '' comment '重定向地址, 多个地址使用,分隔',
    `sign_up_url`   varchar(200) null default '' comment '注册地址',
    `sign_in_url`   varchar(200) null default '' comment '登录地址',
    `sign_out_url`  varchar(200) null default '' comment '登出地址',
    `terms_of_use`  varchar(200) null default '' comment '用户协议',
    `create_at`     timestamp    null default current_timestamp comment '创建时间',
    `create_by`     bigint       null default 0 comment '创建人',
    `modify_at`     timestamp    null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by`     bigint       null default 0 comment '更新人',
    `deleted`       tinyint      null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='应用表';

drop table if exists `application`;
create table `application`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='应用表';

drop table if exists `social`;
create table `social`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='三方社交表';

drop table if exists `company`;
create table `company`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='公司表';

drop table if exists `ldap`;
create table `ldap`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='LDAP表';

drop table if exists `multi_factor_auth`;
create table `multi_factor_auth`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='MFA多因素认证表';

drop table if exists `role`;
create table `role`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='角色表';

drop table if exists `permission`;
create table `permission`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='权限表';

drop table if exists `account_group`;
create table `account_group`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='账号组表';

drop table if exists `sign_in_token`;
create table `sign_in_token`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='登录令牌表';

# account tag
# account score
# account app
# account social
# account company
# account ldap
# account webauthn
# account mfa
# account role
# account perm
# account group
# account token

# provider
# provider account


drop table if exists `account`;
create table `account`
(
    `id`        bigint    not null comment '主键' auto_increment,
    `create_at` timestamp null default current_timestamp comment '创建时间',
    `create_by` bigint    null default 0 comment '创建人',
    `modify_at` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    `modify_by` bigint    null default 0 comment '更新人',
    `deleted`   tinyint   null default 0 comment '逻辑删除',
    primary key (`id`) using btree
) engine = innodb
  auto_increment = 100
  default charset = utf8mb4 comment ='表';
 