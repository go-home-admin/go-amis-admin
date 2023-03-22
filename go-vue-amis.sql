/*
 Navicat Premium Data Transfer

 Source Server         : 华为
 Source Server Type    : MySQL
 Source Server Version : 80024 (8.0.24)
 Source Host           : 120.46.149.36:3306
 Source Schema         : go-vue-amis

 Target Server Type    : MySQL
 Target Server Version : 80024 (8.0.24)
 File Encoding         : 65001

 Date: 22/03/2023 15:59:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组件名称',
  `component` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组件',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '地址',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '重定向',
  `meta` json NOT NULL COMMENT '元数据',
  `hidden` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '0忽略1隐藏2显示',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `api_list` json DEFAULT NULL COMMENT 'api',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'created_at',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'updated_at',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` (`id`, `parent_id`, `name`, `component`, `path`, `redirect`, `meta`, `hidden`, `sort`, `api_list`, `created_at`, `updated_at`) VALUES (1, 0, '', 'Layout', '/sys', '/menus', '{\"icon\": \"el-icon-s-help\", \"title\": \"系统设置\"}', 0, 1, NULL, '2023-02-22 00:00:00', '2023-02-22 19:40:27');
INSERT INTO `admin_menu` (`id`, `parent_id`, `name`, `component`, `path`, `redirect`, `meta`, `hidden`, `sort`, `api_list`, `created_at`, `updated_at`) VALUES (91, 1, 'users', 'amis/index', 'users', '', '{\"amis\": \"users\", \"icon\": \"example\", \"title\": \"用户\"}', 0, 1, NULL, '2023-02-22 00:00:00', '2023-02-27 14:33:27');
INSERT INTO `admin_menu` (`id`, `parent_id`, `name`, `component`, `path`, `redirect`, `meta`, `hidden`, `sort`, `api_list`, `created_at`, `updated_at`) VALUES (112, 1, 'menus', 'auth/admin_menus', 'menus', '', '{\"icon\": \"el-icon-s-help\", \"title\": \"菜单管理\"}', 0, 500, NULL, NULL, '2023-02-27 13:55:09');
INSERT INTO `admin_menu` (`id`, `parent_id`, `name`, `component`, `path`, `redirect`, `meta`, `hidden`, `sort`, `api_list`, `created_at`, `updated_at`) VALUES (113, 1, 'roles', 'amis/index', 'roles', '', '{\"amis\": \"/roles\", \"icon\": \"example\", \"title\": \"角色管理\"}', 0, 500, NULL, NULL, '2023-02-28 09:59:40');
COMMIT;

-- ----------------------------
-- Table structure for admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menu`;
CREATE TABLE `admin_role_menu` (
  `role_id` int NOT NULL COMMENT 'role_id',
  `menu_id` int NOT NULL COMMENT 'menu_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'created_at',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'updated_at',
  KEY `laravel_admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色关联菜单';

-- ----------------------------
-- Records of admin_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_role_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_users`;
CREATE TABLE `admin_role_users` (
  `role_id` int NOT NULL COMMENT 'role_id',
  `user_id` int NOT NULL COMMENT 'user_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'created_at',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'updated_at',
  KEY `laravel_admin_role_users_role_id_user_id_index` (`role_id`,`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户关联角色';

-- ----------------------------
-- Records of admin_role_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES (1, 2, '2023-03-22 14:02:57', '2023-03-22 14:02:57');
INSERT INTO `admin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES (2, 8, '2023-03-22 14:03:03', '2023-03-22 14:03:03');
INSERT INTO `admin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES (3, 11, '2023-03-22 14:03:08', '2023-03-22 14:03:08');
INSERT INTO `admin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES (3, 12, '2023-03-22 14:03:38', '2023-03-22 14:03:38');
COMMIT;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名',
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '默认权限',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'created_at',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'updated_at',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `laravel_admin_roles_name_unique` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='权限角色表';

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles` (`id`, `name`, `slug`, `remark`, `created_at`, `updated_at`) VALUES (1, '超级管理员', '', '', '2023-03-22 13:56:38', '2023-03-22 13:56:38');
INSERT INTO `admin_roles` (`id`, `name`, `slug`, `remark`, `created_at`, `updated_at`) VALUES (2, '普通管理员', '', 'id=1拥有所有权限', '2023-02-28 10:23:55', '2023-02-28 10:45:32');
INSERT INTO `admin_roles` (`id`, `name`, `slug`, `remark`, `created_at`, `updated_at`) VALUES (3, '测试', '', '', '2023-03-22 13:56:50', '2023-03-22 13:56:50');
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(190) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账户',
  `password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '显示名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'created_at',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_at',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `laravel_admin_users_username_unique` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='后台用户';

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users` (`id`, `username`, `password`, `name`, `avatar`, `deleted_at`, `created_at`, `updated_at`) VALUES (2, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://avatars.githubusercontent.com/u/18717080?s=30&v=4', NULL, NULL, '2023-03-22 14:02:57');
INSERT INTO `admin_users` (`id`, `username`, `password`, `name`, `avatar`, `deleted_at`, `created_at`, `updated_at`) VALUES (8, 'root', 'e10adc3949ba59abbe56e057f20f883e', '普通管理员', 'http://127.0.0.1:5000/files/images/avatar/2023-03-03/01d381cf-c1b6-4fe6-8fe8-c55b66689f9e..jpeg', NULL, '2023-03-03 14:06:43', '2023-03-22 14:03:03');
INSERT INTO `admin_users` (`id`, `username`, `password`, `name`, `avatar`, `deleted_at`, `created_at`, `updated_at`) VALUES (11, 'test', 'e10adc3949ba59abbe56e057f20f883e', '测试', NULL, NULL, '2023-03-22 13:57:08', '2023-03-22 14:03:08');
INSERT INTO `admin_users` (`id`, `username`, `password`, `name`, `avatar`, `deleted_at`, `created_at`, `updated_at`) VALUES (12, 'zhangsan', 'e10adc3949ba59abbe56e057f20f883e', '张三', NULL, NULL, '2023-03-22 14:03:37', '2023-03-22 14:03:38');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
