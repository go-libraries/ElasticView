/*
Navicat MySQL Data Transfer

Source Server         : 47.110.85.147
Source Server Version : 50731
Source Host           : 47.110.85.147:3306
Source Database       : es_view

Target Server Type    : MYSQL
Target Server Version : 50731
File Encoding         : 65001

Date: 2021-03-11 18:40:12
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of es_link
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
INSERT INTO `gm_role` VALUES ('1', 'admin', '超级管理员', '[{\"path\":\"/permission\",\"component\":\"layout\",\"redirect\":\"/permission/index\",\"alwaysShow\":true,\"meta\":{\"title\":\"权限\",\"icon\":\"el-icon-user-solid\"},\"children\":[{\"path\":\"role\",\"component\":\"views/permission/role\",\"name\":\"RolePermission\",\"meta\":{\"title\":\"角色管理\",\"icon\":\"el-icon-s-check\"}},{\"path\":\"user\",\"component\":\"views/permission/user\",\"name\":\"user\",\"meta\":{\"title\":\"用户管理\",\"icon\":\"el-icon-user\"}}]},{\"path\":\"/connect-tree\",\"component\":\"layout\",\"redirect\":\"/connect-tree/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"连接树管理\",\"icon\":\"el-icon-link\"},\"children\":[{\"path\":\"/connect-tree/index\",\"component\":\"views/connect-tree/index\",\"name\":\"index\",\"meta\":{\"title\":\"连接树管理\",\"icon\":\"el-icon-link\"}}]},{\"path\":\"/cat\",\"component\":\"layout\",\"redirect\":\"/cat/index\",\"alwaysShow\":false,\"meta\":{\"title\":\"ES状态\",\"icon\":\"el-icon-pie-chart\"},\"children\":[{\"path\":\"/cat/index\",\"component\":\"views/cat/index\",\"name\":\"index\",\"meta\":{\"title\":\"ES状态\",\"icon\":\"el-icon-pie-chart\"}}]}]');

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of gm_user
-- ----------------------------
INSERT INTO `gm_user` VALUES ('1', 'admin', 'admin', '1', '肖文龙');
