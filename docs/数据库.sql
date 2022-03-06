CREATE DATABASE tblog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
create user 'dev'@'%' identified by 'machine03';
grant all privileges on tblog.* to dev;

drop table if exists t_menu  ;
create table t_menu (
    id bigint(20) not null primary key auto_increment ,
    pid_ bigint(20) not null default 0 comment '' ,
    name_ varchar(50) not null comment '名称' ,
    deleted_ tinyint(1) not null default 0 comment '是否删除。0=未删除，1=已删除',
    created_by varchar(20) default 0 comment '创建人' ,
    update_by varchar(20) default 0 comment '变更人' ,
    created_time datetime comment '创建时间' ,
    update_time datetime comment '变更时间'
) comment '菜单';