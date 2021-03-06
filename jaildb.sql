/*
 Navicat Premium Data Transfer

 Source Server         : Local
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : jaildb

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 23/05/2022 21:43:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crime
-- ----------------------------
DROP TABLE IF EXISTS `crime`;
CREATE TABLE `crime`  (
  `crimeID` int NOT NULL AUTO_INCREMENT,
  `crimeDescription` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `extendDescription` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sentence` int NOT NULL,
  `sentenceStartDate` datetime NOT NULL,
  `lastModifiedDate` datetime NOT NULL,
  `createdDate` datetime NOT NULL,
  PRIMARY KEY (`crimeID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of crime
-- ----------------------------

-- ----------------------------
-- Table structure for inmates
-- ----------------------------
DROP TABLE IF EXISTS `inmates`;
CREATE TABLE `inmates`  (
  `inmateID` int NOT NULL AUTO_INCREMENT,
  `crimeID` int NOT NULL,
  `firstName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `lastName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `middleName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `birthPlace` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `birthDate` datetime NOT NULL,
  `lastModifiedDate` datetime NOT NULL,
  `createdDate` datetime NOT NULL,
  PRIMARY KEY (`inmateID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of inmates
-- ----------------------------
INSERT INTO `inmates` VALUES (1, 1, 'Sample', 'Sample', 'X', 'Purok4A Poblacion Valencia City, Bukidnon', 'Zamboanga City', '1998-02-01 00:00:00', '2022-03-19 20:42:46', '2022-03-19 20:42:46');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `userID` int NOT NULL AUTO_INCREMENT,
  `userName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `userType` int NOT NULL,
  PRIMARY KEY (`userID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (4, 'admin', '$2a$14$LB/JqF5eYUcogIs2/L0kbu5AzxYeMhIjyYzPNbM0Nrl9DkATYezY.', 1);
INSERT INTO `users` VALUES (5, 'user', '$2a$14$TGT/70/VgVWloX2XxXhGAOJEQFS2SPSxuIZyVIF33BLQnkVXWIRLm', 2);
INSERT INTO `users` VALUES (6, 'sample', '$2a$14$.6kwBoiXe8DT80nbMnvQ2eh9sOJp6xqZKMta9ssuoiKkUu9y5n0Wq', 2);

-- ----------------------------
-- Table structure for visitors
-- ----------------------------
DROP TABLE IF EXISTS `visitors`;
CREATE TABLE `visitors`  (
  `visitorID` int NOT NULL AUTO_INCREMENT,
  `userID` int NOT NULL,
  `firstName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `lastName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `middleName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `birthPlace` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `birthDate` datetime NOT NULL,
  `lastModifiedDate` datetime NOT NULL,
  `createdDate` datetime NOT NULL,
  PRIMARY KEY (`visitorID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of visitors
-- ----------------------------
INSERT INTO `visitors` VALUES (2, 3, 'Kennen', 'Borbon', 'Comaling', 'Purok 4A Poblacion, Valencia City, Bukidnon', 'New Bataan, Compostella Valley', '1998-10-13 11:45:18', '2022-05-06 22:29:45', '2022-05-06 22:29:45');
INSERT INTO `visitors` VALUES (3, 6, 'User', 'X', 'User', ' Valencia City, Bukidnon', 'BPH Kibawer', '1998-10-13 11:45:18', '2022-05-10 18:13:20', '2022-05-10 18:13:20');

-- ----------------------------
-- View structure for vwgetallinmates
-- ----------------------------
DROP VIEW IF EXISTS `vwgetallinmates`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `vwgetallinmates` AS select `inmates`.`inmateID` AS `inmateID`,`inmates`.`crimeID` AS `crimeID`,`inmates`.`firstName` AS `firstName`,`inmates`.`lastName` AS `lastName`,`inmates`.`middleName` AS `middleName`,`inmates`.`address` AS `address`,`inmates`.`birthPlace` AS `birthPlace`,`inmates`.`birthDate` AS `birthDate`,`crime`.`crimeDescription` AS `crimeDescription`,`crime`.`sentence` AS `sentence`,`inmates`.`lastModifiedDate` AS `lastModifiedDate`,`inmates`.`createdDate` AS `createdDate` from (`inmates` join `crime` on((`inmates`.`crimeID` = `crime`.`crimeID`)));

-- ----------------------------
-- View structure for vwuservisitor
-- ----------------------------
DROP VIEW IF EXISTS `vwuservisitor`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `vwuservisitor` AS select `users`.`userID` AS `userID`,`users`.`userName` AS `userName`,`visitors`.`firstName` AS `firstName`,`visitors`.`lastName` AS `lastName`,`visitors`.`middleName` AS `middleName`,`visitors`.`birthPlace` AS `birthPlace` from (`users` join `visitors` on((`users`.`userID` = `visitors`.`userID`)));

SET FOREIGN_KEY_CHECKS = 1;
