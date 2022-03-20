insert into t_topic (id, pid_, title_,image_url,rotation_, sorted_, deleted_,published_)
values  (1, 0, '首页',null,false, 1, 0,1),
        (2, 0, 'Java',null,false, 2, 0,1),
        (3, 2, 'Java基础',null,false, 3, 0,1),
        (4, 2, 'JVM',null,false, 4, 0,1),
        (5, 0, 'Spring',null,false, 5, 0,1 ),
        (6, 5, 'Spring Framework',null,false, 6, 0,1 ),
        (7, 5, 'Spring Boot',null,false, 7, 0 ,1),
        (8, 5, 'Spring Cloud',null,false, 8, 0 ,1);


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


