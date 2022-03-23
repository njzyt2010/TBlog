CREATE DATABASE tblog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
create user 'dev'@'%' identified by 'machine03';
grant all privileges on tblog.* to dev;

drop table if exists t_topic;
create table t_topic
(
    id           bigint(20)  not null primary key auto_increment,
    pid_         bigint(20)  not null default 0 comment '父级栏目',
    title_        varchar(50) not null comment '名称',
    image_url    varchar(255) comment '图片地址',
    rotation_        tinyint(1)           default 0 comment '是否为设置为轮播',
    published_   tinyint(1)           default 0 comment '栏目是否发布',
    sorted_      int(3)      not null comment '排序',
    deleted_     tinyint(1)  not null default 0 comment '是否删除。0=未删除，1=已删除',
    created_by   varchar(20)          default 0 comment '创建人',
    update_by    varchar(20)          default 0 comment '变更人',
    created_time datetime comment '创建时间',
    update_time  datetime comment '变更时间'
) comment '栏目';

drop table if exists t_article;
create table t_article
(
    id           bigint(20)   not null primary key auto_increment,
    title_       varchar(100) not null comment '标题',
    content_     longtext comment '文章内容',
    reprint_     tinyint(1) comment '是否转载',
    reprint_url  varchar(255) comment '转载文章的地址',
    topic_id     bigint(20) comment '所属栏目',
    published_   tinyint(1)            default 0 comment '文章是否发布',
    deleted_     tinyint(1)   not null default 0 comment '是否删除。0=未删除，1=已删除',
    created_by   varchar(20)           default 0 comment '创建人',
    update_by    varchar(20)           default 0 comment '变更人',
    created_time datetime comment '创建时间',
    update_time  datetime comment '变更时间'

) comment '文章';

drop table if exists t_tag;
create table t_tag
(
    id           bigint(20)   not null primary key auto_increment,
    name_        varchar(100) not null comment '标签',
    deleted_     tinyint(1)   not null default 0 comment '是否删除。0=未删除，1=已删除',
    created_by   varchar(20)           default 0 comment '创建人',
    update_by    varchar(20)           default 0 comment '变更人',
    created_time datetime comment '创建时间',
    update_time  datetime comment '变更时间'
) comment '标签';

drop table if exists t_article_tag;
create table t_article_tag
(
    id         bigint(20) not null primary key auto_increment,
    tag_id     bigint(20) comment '标签id',
    article_id bigint(20) comment '文章id'
) comment '文章-标签';


alter table t_topic  add newer_ tinyint(1) default 0 not null comment '是否显示在最近更新';