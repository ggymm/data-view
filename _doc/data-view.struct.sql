/*
 Navicat Premium Data Transfer

 Source Server         : 本地服务器
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : localhost:3306
 Source Schema         : data-view

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 13/07/2021 20:18:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for data_source
-- ----------------------------
DROP TABLE IF EXISTS `data_source`;
CREATE TABLE `data_source`  (
  `data_source_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `business_id` bigint(20) NULL DEFAULT NULL COMMENT '企业ID',
  `data_source_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源名称',
  `data_source_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源类型',
  `data_source_database_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源的数据库名称',
  `data_source_ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源的IP地址',
  `data_source_port` int(5) NULL DEFAULT NULL COMMENT '数据源的端口号',
  `data_source_username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源的账户名称',
  `data_source_password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源的账户密码',
  `create_id` bigint(20) NULL DEFAULT NULL COMMENT '创建者ID',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_id` bigint(20) NULL DEFAULT NULL COMMENT '更新者ID',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `del_flag` tinyint(4) NULL DEFAULT NULL COMMENT '删除标识（-1：不存在；1：存在）',
  PRIMARY KEY (`data_source_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image`  (
  `image_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `image_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片名称',
  `image_path` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片位置',
  `image_size` bigint(20) NULL DEFAULT NULL COMMENT '图片大小',
  `image_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片类型',
  `business_id` bigint(20) NULL DEFAULT NULL COMMENT '企业ID',
  `create_id` bigint(20) NULL DEFAULT NULL COMMENT '创建者ID',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_id` bigint(20) NULL DEFAULT NULL COMMENT '更新者ID',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `del_flag` tinyint(4) NULL DEFAULT NULL COMMENT '删除标识（-1：不存在；1：存在）',
  PRIMARY KEY (`image_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for view_chart_item
-- ----------------------------
DROP TABLE IF EXISTS `view_chart_item`;
CREATE TABLE `view_chart_item`  (
  `item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '数据库ID自增',
  `instance_id` bigint(20) NULL DEFAULT NULL COMMENT '大屏实例ID',
  `item_el_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'HTML元素ID',
  `item_index` int(11) NULL DEFAULT NULL COMMENT '索引（用于分层）',
  `item_chart_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `item_chart_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型',
  `item_chart_version` bigint(20) NULL DEFAULT NULL COMMENT '版本',
  `item_chart_data` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '数据配置',
  `item_data` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '数据',
  `item_rotate` int(11) NULL DEFAULT NULL COMMENT '旋转',
  `item_show` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '显示',
  `item_lock` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '锁定',
  `item_choose` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '选中',
  `item_interval` bigint(20) NULL DEFAULT NULL COMMENT '数据刷新',
  `item_option` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '配置项',
  `item_refresh` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '刷新',
  `item_width` bigint(20) NULL DEFAULT NULL COMMENT '宽度',
  `item_height` bigint(20) NULL DEFAULT NULL COMMENT '高度',
  `item_x` bigint(20) NULL DEFAULT NULL COMMENT 'x坐标',
  `item_y` bigint(20) NULL DEFAULT NULL COMMENT 'y坐标',
  `item_version` bigint(20) NULL DEFAULT NULL COMMENT '实例版本',
  PRIMARY KEY (`item_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3264 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for view_instance
-- ----------------------------
DROP TABLE IF EXISTS `view_instance`;
CREATE TABLE `view_instance`  (
  `instance_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '实例ID',
  `business_id` bigint(20) NULL DEFAULT NULL COMMENT '企业ID',
  `instance_background_color` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '背景颜色',
  `instance_background_img` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '背景图',
  `instance_width` bigint(20) NULL DEFAULT NULL COMMENT '宽',
  `instance_height` bigint(20) NULL DEFAULT NULL COMMENT '高',
  `instance_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `instance_view_thumbnail` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '缩略图',
  `instance_version` bigint(20) NULL DEFAULT NULL COMMENT '版本号',
  `create_id` bigint(20) NULL DEFAULT NULL COMMENT '创建者ID',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_id` bigint(20) NULL DEFAULT NULL COMMENT '更新者ID',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `del_flag` tinyint(4) NULL DEFAULT NULL COMMENT '删除标识（-1：不存在；1：存在）',
  PRIMARY KEY (`instance_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 92 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
