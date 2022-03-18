insert into tblog.t_menu (id, pid_, title_, sorted_, deleted_, created_by, update_by, created_time, update_time)
values  (1, 0, '首页', 1, 0, '0', '0', null, null),
        (2, 0, 'Java', 2, 0, '0', '0', null, null),
        (3, 2, 'Java基础', 3, 0, '0', '0', null, null),
        (4, 2, 'JVM', 4, 0, '0', '0', null, null),
        (5, 0, 'Spring', 5, 0, '0', '0', null, null),
        (6, 5, 'Spring Framework', 6, 0, '0', '0', null, null),
        (7, 5, 'Spring Boot', 7, 0, '0', '0', null, null),
        (8, 5, 'Spring Cloud', 8, 0, '0', '0', null, null);


insert into t_article(title_, content_, reprint_, reprint_url, topic_id, published_, deleted_) VALUES
("测试文章","测试的内容",0,'',1,1,0)