insert into tblog.t_menu (id, pid_, title_, sorted_, deleted_, created_by, update_by, created_time, update_time)
values  (1, 0, '首页', 1, 0, '0', '0', null, null),
        (2, 0, 'Java', 2, 0, '0', '0', null, null),
        (3, 2, 'Java基础', 3, 0, '0', '0', null, null),
        (4, 2, 'JVM', 4, 0, '0', '0', null, null),
        (5, 0, 'Spring', 5, 0, '0', '0', null, null),
        (6, 5, 'Spring Framework', 6, 0, '0', '0', null, null),
        (7, 5, 'Spring Boot', 7, 0, '0', '0', null, null),
        (8, 5, 'Spring Cloud', 8, 0, '0', '0', null, null);


insert into tblog.t_article (id, title_, content_, reprint_, reprint_url, topic_id, published_, deleted_, created_by, update_by, created_time, update_time)
values

(1, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(2, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(3, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(4, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(5, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(6, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(7, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(8, '测试文章', '测试的内容', 0, '', 4, 1, 0, '0', '0', null, null),
(9, '测试文章', '测试的内容', 0, '', 5, 1, 0, '0', '0', null, null),
(10, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null),
(11, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null),
(12, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null),
(13, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null),
(14, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null),
(15, '测试文章', '测试的内容', 0, '', 3, 1, 0, '0', '0', null, null);


