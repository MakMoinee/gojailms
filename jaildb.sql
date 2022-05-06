/*
 Navicat Premium Data Transfer

 Source Server         : Local DB
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : jaildb

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 06/05/2022 13:43:47
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of inmates
-- ----------------------------
INSERT INTO `inmates` VALUES (1, 1, 'Sample', 'Sample', 'X', 'Purok4A Poblacion Valencia City, Bukidnon', 'Zamboanga City', '1998-02-01 00:00:00', '2022-03-19 20:42:46', '2022-03-19 20:42:46');
INSERT INTO `inmates` VALUES (2, 1, 'Sample', 'Sample', 'X', 'Purok4A Poblacion Valencia City, Bukidnon', 'Zamboanga City', '1998-02-01 00:00:00', '2022-03-19 22:05:53', '2022-03-19 22:05:53');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `userID` int NOT NULL AUTO_INCREMENT,
  `userName` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `userType` int NOT NULL,
  PRIMARY KEY (`userID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (8, 'ken', '$2a$14$DNxpB0cl.Mo2WPJ2az248e2W5sOotenpMWcoBAyL2VBxPGOc6wh1i', 2);
INSERT INTO `users` VALUES (14, 'mak12', '$2a$14$ZFdWuIFIoAvUImivraVMPeTIbE0DN2VwTlJO5mT8HULwm7ZkoFpaq', 2);
INSERT INTO `users` VALUES (18, 'mak12', '$2a$14$suA/RimIfvJg.Zrh./bjTOgEs5ibBoSNnuCEd2KpCuYxPsU8EBn.a', 2);
INSERT INTO `users` VALUES (38, 'sam123', '$2a$14$jUM/TIzvKadzcyYZPOR5COqvyZD4YmHzGpY0xNjIRtPtOyoVVTypy', 2);
INSERT INTO `users` VALUES (39, 'hello', '$2a$14$J8AFbEPdIiLxgnYksilmneksv4.9phJshJlllLwzcdhN9sgdtNgwu', 2);
INSERT INTO `users` VALUES (40, 'hello', '$2a$14$6gdm1gqkWB5PuU2n8F3DEuXCUgdNXTWxMuZjF2vQhj9WmX6oNITFG', 2);
INSERT INTO `users` VALUES (41, 'admin', '$2a$14$SIqCm4uqCg3JiaxjCCCRjOoVdqFlG4jG1Pe8BGQaKcjT2TuyueKdW', 1);

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
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of visitors
-- ----------------------------
INSERT INTO `visitors` VALUES (1, 2, 'admin', 'admin', 'X', 'None', 'none', '1998-10-13 11:45:18', '2022-03-19 17:49:01', '2022-03-19 17:49:01');
INSERT INTO `visitors` VALUES (5, 14, 'Kennen', 'Borbon', 'Comaling', 'Purok 4A Poblacion, Valencia City, Bukidnon', 'New Bataan, Compostella Valley', '1998-10-13 11:45:18', '2022-05-02 23:36:41', '2022-05-02 23:36:41');
INSERT INTO `visitors` VALUES (15, 38, 'Kennen', 'Borbon', 'Comaling', 'Purok 4A Poblacion, Valencia City, Bukidnon', 'New Bataan, Compostella Valley', '1998-10-13 11:45:18', '2022-05-03 18:46:07', '2022-05-03 18:46:07');
INSERT INTO `visitors` VALUES (16, 39, 'Hello', 'Hello', 'X', 'NA', 'NA', '2022-05-03 00:00:00', '2022-05-03 19:09:18', '2022-05-03 19:09:18');
INSERT INTO `visitors` VALUES (17, 39, 'Hello', 'Hello', 'X', 'NA', 'NA', '2022-05-03 00:00:00', '2022-05-03 19:09:21', '2022-05-03 19:09:21');

-- ----------------------------
-- View structure for vwgetallinmates
-- ----------------------------
DROP VIEW IF EXISTS `vwgetallinmates`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `vwgetallinmates` AS select `inmates`.`inmateID` AS `inmateID`,`inmates`.`crimeID` AS `crimeID`,`inmates`.`firstName` AS `firstName`,`inmates`.`lastName` AS `lastName`,`inmates`.`middleName` AS `middleName`,`inmates`.`address` AS `address`,`inmates`.`birthPlace` AS `birthPlace`,`inmates`.`birthDate` AS `birthDate`,`crime`.`crimeDescription` AS `crimeDescription`,`crime`.`sentence` AS `sentence`,`inmates`.`lastModifiedDate` AS `lastModifiedDate`,`inmates`.`createdDate` AS `createdDate` from (`inmates` join `crime` on((`inmates`.`crimeID` = `crime`.`crimeID`)));

SET FOREIGN_KEY_CHECKS = 1;
