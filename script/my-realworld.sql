/*
 Navicat Premium Data Transfer

 Source Server         : MySQL80_win
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : localhost:3306
 Source Schema         : my-realworld

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 21/04/2025 23:50:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for _articletotag
-- ----------------------------
DROP TABLE IF EXISTS `_articletotag`;
CREATE TABLE `_articletotag`  (
  `A` int(0) NOT NULL,
  `B` int(0) NOT NULL,
  INDEX `_ArticleToTag_A_fkey`(`A`) USING BTREE,
  INDEX `_ArticleToTag_B_fkey`(`B`) USING BTREE,
  CONSTRAINT `_ArticleToTag_A_fkey` FOREIGN KEY (`A`) REFERENCES `article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `_ArticleToTag_B_fkey` FOREIGN KEY (`B`) REFERENCES `tag` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of _articletotag
-- ----------------------------

-- ----------------------------
-- Table structure for _userfavorites
-- ----------------------------
DROP TABLE IF EXISTS `_userfavorites`;
CREATE TABLE `_userfavorites`  (
  `A` int(0) NOT NULL,
  `B` int(0) NOT NULL,
  INDEX `_UserFavorites_A_fkey`(`A`) USING BTREE,
  INDEX `_UserFavorites_B_fkey`(`B`) USING BTREE,
  CONSTRAINT `_UserFavorites_A_fkey` FOREIGN KEY (`A`) REFERENCES `article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `_UserFavorites_B_fkey` FOREIGN KEY (`B`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of _userfavorites
-- ----------------------------

-- ----------------------------
-- Table structure for _userfollows
-- ----------------------------
DROP TABLE IF EXISTS `_userfollows`;
CREATE TABLE `_userfollows`  (
  `A` int(0) NOT NULL,
  `B` int(0) NOT NULL,
  INDEX `_UserFollows_A_fkey`(`A`) USING BTREE,
  INDEX `_UserFollows_B_fkey`(`B`) USING BTREE,
  CONSTRAINT `_UserFollows_A_fkey` FOREIGN KEY (`A`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `_UserFollows_B_fkey` FOREIGN KEY (`B`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of _userfollows
-- ----------------------------

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `slug` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updatedAt` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `authorId` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `Article_authorId_fkey`(`authorId`) USING BTREE,
  CONSTRAINT `Article_authorId_fkey` FOREIGN KEY (`authorId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `createdAt` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updatedAt` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `articleId` int(0) NOT NULL,
  `authorId` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `Comment_articleId_fkey`(`articleId`) USING BTREE,
  INDEX `Comment_authorId_fkey`(`authorId`) USING BTREE,
  CONSTRAINT `Comment_articleId_fkey` FOREIGN KEY (`articleId`) REFERENCES `article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Comment_authorId_fkey` FOREIGN KEY (`authorId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'https://api.realworld.io/images/smiley-cyrus.jpeg',
  `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `demo` tinyint(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'haipeng_lin@outlook.com', '林海鹏', '123456789', 'https://i.stack.imgur.com/xHWG8.jpg', 'I love cat', 1);
INSERT INTO `user` VALUES (2, '306372404@qq.com', '鱼', '123456', 'https://api.realworld.io/images/smiley-cyrus.jpeg', NULL, 0);
INSERT INTO `user` VALUES (3, '306372411@qq.com', '鱼123', '12345678', 'https://api.realworld.io/images/smiley-cyrus.jpeg', NULL, 0);

SET FOREIGN_KEY_CHECKS = 1;
