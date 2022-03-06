CREATE DATABASE tblog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
create user 'dev'@'%' identified by 'machine03';
grant all privileges on tblog.* to dev;

drop table if exists t_menu  ;
create table t_menu (
    id bigint(20) not null primary key auto_increment ,
    pid_ bigint(20) not null default 0 comment '' ,
    name_ varchar(50) not null comment '名称' ,
    sorted_ int(3) not null comment '排序',
    deleted_ tinyint(1) not null default 0 comment '是否删除。0=未删除，1=已删除',
    created_by varchar(20) default 0 comment '创建人' ,
    update_by varchar(20) default 0 comment '变更人' ,
    created_time datetime comment '创建时间' ,
    update_time datetime comment '变更时间'
) comment '菜单';


drop table if exists t_topic  ;
create table t_topic (
        id bigint(20) not null primary key auto_increment ,
        name_ varchar(50) not null comment '名称' ,
        image_url varchar(255) comment '图片地址' ,
        type_ tinyint(1) default 0 comment '栏目类型，0为普通类型，1为轮播图' ,
        published_ tinyint(1) default 0 comment '栏目是否发布' ,
        sorted_ int(3) not null comment '排序',
        deleted_ tinyint(1) not null default 0 comment '是否删除。0=未删除，1=已删除',
        created_by varchar(20) default 0 comment '创建人' ,
        update_by varchar(20) default 0 comment '变更人' ,
        created_time datetime comment '创建时间' ,
        update_time datetime comment '变更时间'
) comment '菜单';