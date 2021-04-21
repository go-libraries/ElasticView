/*
Navicat MySQL Data Transfer

Source Server         : 测试服
Source Server Version : 50727
Source Host           : 192.168.1.236:3306
Source Database       : es_view

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2021-04-13 19:17:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `es_link`
-- ----------------------------
DROP TABLE IF EXISTS `es_link`;
CREATE TABLE `es_link` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(100) NOT NULL,
  `user` varchar(100) NOT NULL,
  `pwd` varchar(100) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `remark` varchar(100) DEFAULT '默认连接',
  `version` tinyint(10) NOT NULL DEFAULT '6',
  PRIMARY KEY (`id`),
  UNIQUE KEY `es_remark` (`remark`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of es_link
-- ----------------------------

-- ----------------------------
-- Table structure for `gm_dsl_history`
-- ----------------------------
DROP TABLE IF EXISTS `gm_dsl_history`;
CREATE TABLE `gm_dsl_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT '0',
  `method` varchar(50) DEFAULT '',
  `path` varchar(255) DEFAULT '',
  `body` text,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=829 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gm_dsl_history
-- ----------------------------

-- ----------------------------
-- Table structure for `gm_guid`
-- ----------------------------
DROP TABLE IF EXISTS `gm_guid`;
CREATE TABLE `gm_guid` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `guid_name` varchar(100) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid_name` (`uid`,`guid_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gm_guid
-- ----------------------------

-- ----------------------------
-- Table structure for `gm_role`
-- ----------------------------
DROP TABLE IF EXISTS `gm_role`;
CREATE TABLE `gm_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `role_list` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gm_role
-- ----------------------------
INSERT INTO `gm_role` VALUES ('1', 'admin', '超级管理员', '[{\"path\":\"/permission\",\"component\":\"layout\",\"redirect\":\"/permission/index\",\"alwaysShow\":true,\"meta\":{\"title\":\"权限\",\"icon\":\"el-icon-user-solid\"},\"children\":[{\"path\":\"role\",\"component\":\"views/permission/role\",\"name\":\"RolePermission\",\"meta\":{\"title\":\"角色管理\",\"icon\":\"el-icon-s-check\"}},{\"path\":\"user\",\"component\":\"views/permission/user\",\"name\":\"user\",\"meta\":{\"title\":\"用户管理\",\"icon\":\"el-icon-user\"}}]},{\"path\":\"/connect-tree\",\"component\":\"layout\",\"redirect\":\"/connect-tree/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"连接树管理\",\"icon\":\"el-icon-link\"},\"children\":[{\"path\":\"/connect-tree/index\",\"component\":\"views/connect-tree/index\",\"name\":\"index\",\"meta\":{\"title\":\"连接树管理\",\"icon\":\"el-icon-link\"}}]},{\"path\":\"/cat\",\"component\":\"layout\",\"redirect\":\"/cat/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"ES状态\",\"icon\":\"el-icon-pie-chart\"},\"children\":[{\"path\":\"/cat/index\",\"component\":\"views/cat/index\",\"name\":\"index\",\"meta\":{\"title\":\"ES状态\",\"icon\":\"el-icon-pie-chart\"}}]},{\"path\":\"/rest\",\"component\":\"layout\",\"redirect\":\"/rest/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"开发工具\",\"icon\":\"el-icon-edit\"},\"children\":[{\"path\":\"/rest/index\",\"component\":\"views/rest/index\",\"name\":\"index\",\"meta\":{\"title\":\"开发工具\",\"icon\":\"el-icon-search\"}}]},{\"path\":\"/indices\",\"component\":\"layout\",\"redirect\":\"/indices/index\",\"alwaysShow\":true,\"meta\":{\"title\":\"索引管理\",\"icon\":\"el-icon-coin\"},\"children\":[{\"path\":\"index\",\"component\":\"views/indices/index\",\"name\":\"index\",\"meta\":{\"title\":\"索引管理\",\"icon\":\"el-icon-coin\"}},{\"path\":\"reindex\",\"component\":\"views/indices/reindex\",\"name\":\"reindex\",\"meta\":{\"title\":\"重建索引\",\"icon\":\"el-icon-document-copy\"}}]},{\"path\":\"/task\",\"component\":\"layout\",\"redirect\":\"/task/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"任务\",\"icon\":\"el-icon-notebook-2\"},\"children\":[{\"path\":\"/task/index\",\"component\":\"views/task/index\",\"name\":\"index\",\"meta\":{\"title\":\"任务\",\"icon\":\"el-icon-notebook-2\"}}]},{\"path\":\"/back-up\",\"component\":\"layout\",\"redirect\":\"/back-up/index\",\"alwaysShow\":true,\"meta\":{\"title\":\"备份\",\"icon\":\"el-icon-copy-document\"},\"children\":[{\"path\":\"index\",\"component\":\"views/back-up/index\",\"name\":\"index\",\"meta\":{\"title\":\"快照存储库\",\"icon\":\"el-icon-first-aid-kit\"}},{\"path\":\"snapshot\",\"component\":\"views/back-up/snapshot\",\"name\":\"index\",\"meta\":{\"title\":\"快照管理\",\"icon\":\"el-icon-shopping-bag-2\"}}]}]');

-- ----------------------------
-- Table structure for `gm_user`
-- ----------------------------
DROP TABLE IF EXISTS `gm_user`;
CREATE TABLE `gm_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL COMMENT '角色id',
  `realname` varchar(255) DEFAULT '' COMMENT '真实姓名',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `gm_user_username` (`username`) USING BTREE COMMENT '角色名唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gm_user
-- ----------------------------
INSERT INTO `gm_user` VALUES ('1', 'admin', '21232f297a57a5a743894a0e4a801fc3', '1', '肖文龙');
