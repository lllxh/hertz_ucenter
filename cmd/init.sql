CREATE DATABASE jason CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use jason;

-- auto-generated definition
create table if not exists user
(
    id           bigint auto_increment comment 'id'
    primary key,
    username     varchar(256)                       not null comment '用户昵称',
    userAccount  varchar(256)                       null comment '账号',
    avatarUrl    varchar(1024)                      null comment '用户头像',
    gender       tinyint                            null comment '性别',
    userPassword varchar(512)                       not null comment '密码',
    phone        varchar(128)                       null comment '电话',
    email        varchar(256)                       null comment 'email',
    userStatus   int      default 0                 not null comment '用户状态',
    createTime   datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    isDelete     tinyint                            null comment '是否删除'
    )
    comment '用户表';

