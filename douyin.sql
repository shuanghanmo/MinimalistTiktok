/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 02/02/2023 20:24:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pass_word` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES (10, 'zhangsan', '031edf66974191f2a22999f36fa39079', '2023-02-02 18:00:42', '2023-02-02 18:00:42');
INSERT INTO `tb_user` VALUES (11, 'lisi', '272420f6f0c0a818f050909b8ef709d5', '2023-02-02 18:23:50', '2023-02-02 18:23:50');
INSERT INTO `tb_user` VALUES (12, '李四', '031edf66974191f2a22999f36fa39079', '2023-02-02 18:31:04', '2023-02-02 18:31:04');

-- ----------------------------
-- Table structure for tb_user_info
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_info`;
CREATE TABLE `tb_user_info`  (
  `id` bigint(20) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `follow_count` bigint(20) NULL DEFAULT 0,
  `follower_count` bigint(255) NULL DEFAULT 0,
  `is_follow` tinyint(1) NULL DEFAULT 0 COMMENT '0:未关注 1:已关注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user_info
-- ----------------------------
INSERT INTO `tb_user_info` VALUES (10, 'douyin_simple_2023:02:02_zhangsan', 0, 0, 0);
INSERT INTO `tb_user_info` VALUES (11, 'douyin_simple_2023:02:02_lisi', 0, 0, 0);
INSERT INTO `tb_user_info` VALUES (12, 'douyin_simple_2023:02:02_李四', 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
