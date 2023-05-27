CREATE TABLE `sys_role` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
                            `name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名称',
                            `remark` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
                            `create_by` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
                            `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `last_update_by` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '更新人',
                            `last_update_time` datetime DEFAULT NULL COMMENT '更新时间',
                            `del_flag` tinyint DEFAULT '0' COMMENT '是否删除  -1：已删除  0：正常',
                            `status` bigint DEFAULT '1' COMMENT '状态  1:启用,0:禁用',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色管理';

INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, 'admin', '超级管理员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 19:07:18', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, 'mng', '项目经理', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:39:28', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (3, 'dev', '开发人员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:39:28', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (4, 'test', '测试人员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:11:11', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (5, 'demo', '1', 'admin', '2020-11-26 14:52:20', 'admin', '2020-11-26 14:50:18', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (6, '1', '1', 'admin', '2020-11-26 15:35:42', 'admin', '2020-11-26 15:01:45', 0);