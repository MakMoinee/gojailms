/*
 Navicat Premium Data Transfer

 Source Server         : LOCAL
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : jaildb

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 30/11/2022 03:36:12
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (11, 'mak', '$2a$14$qYx8PpZmypE5Du4Ynw0anehwk40RRJrjju4XerCOJ7IqyGnxsK/1K', 2);
INSERT INTO `users` VALUES (12, 'ken', '$2a$14$wNnCpZhpipOQ8IQa77OSo.s.8WzEq0C03IDliI3aWBH5lglQQCtMm', 2);
INSERT INTO `users` VALUES (13, 'sheen', '$2a$14$xkHvYqlUbhSQYRAXRcmVKO/TDgRshJOBjyOgNQEld/yqWB/D.bxqy', 2);
INSERT INTO `users` VALUES (14, 'admin', '$2a$14$gVwmYuaMGBPH9xsRx8y5iePZQ9RcrMtBkNXn5f4kHuCRjYWyGfB4O', 1);
INSERT INTO `users` VALUES (15, 'sheenie', '$2a$14$lg77uAmF8jrBEsCzEqT.YuyhnSBqmYOy0qb0XIqLGzKcn5fFkslkS', 2);
INSERT INTO `users` VALUES (16, 'mak', '$2a$14$TtMEXMY54ankWb/MjN48Se1U2RnCHqar5l8LxLvFjwtCOmSW369Au', 2);
INSERT INTO `users` VALUES (17, 'sample', '$2a$14$vTMLQJon5bs.65ZbQ3v9Te5a90q/HLxgVWSy5XIDLO2AqvDOMPs7q', 2);
INSERT INTO `users` VALUES (18, 'admin', '$2a$14$NKT9r6172lWglVqJoKKok.h9YbaI7ouXV4FPboN3DI091UOwQLMRy', 1);

-- ----------------------------
-- Table structure for visitorhistory
-- ----------------------------
DROP TABLE IF EXISTS `visitorhistory`;
CREATE TABLE `visitorhistory`  (
  `visitorHistoryID` int NOT NULL AUTO_INCREMENT,
  `visitorID` int NOT NULL,
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `visitedDateTime` datetime NOT NULL,
  PRIMARY KEY (`visitorHistoryID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitorhistory
-- ----------------------------
INSERT INTO `visitorhistory` VALUES (2, 9, 'she\'s wearing black vest', '2022-11-30 03:34:34');

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
  `birthDate` date NOT NULL,
  `contactNumber` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `lastModifiedDate` datetime NOT NULL,
  `createdDate` datetime NOT NULL,
  PRIMARY KEY (`visitorID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of visitors
-- ----------------------------
INSERT INTO `visitors` VALUES (4, 11, 'mak', 'mak', 'mak', 'Door 10 Rahmann San Jose Extension, Cebu City', 'New Bataan Compostella Valley', '1998-10-13', '09365861683', '2022-08-25 23:13:00', '2022-08-25 23:13:00');
INSERT INTO `visitors` VALUES (5, 12, 'Kennen', 'Borbon', 'Comaling', 'Purok 4A Poblacion Valencia City, Bukidnon', 'New Bataan Compostella Valley', '1998-10-13', '09269440075', '2022-08-25 23:57:38', '2022-08-25 23:57:38');
INSERT INTO `visitors` VALUES (6, 13, 'Sheenie', 'Borbon', 'Ucab', 'Door 10 Rahmann Extension, Kamputhaw, Cebu City', 'Purok 10 Poblacion Valencia City Bukidnon', '2003-12-22', '09269440075', '2022-08-29 22:46:05', '2022-08-29 22:46:05');
INSERT INTO `visitors` VALUES (7, 15, 'Sheenie', 'Borbon', 'x', 'Door 10 Rahmann Extension ', 'purok 10 Poblacion ', '2003-12-22', '09365861683', '2022-10-11 01:23:03', '2022-10-11 01:23:03');
INSERT INTO `visitors` VALUES (8, 11, 'mak', 'mak', 'mak', 'mak', 'mak', '1998-10-13', '09365861683', '2022-10-13 14:46:44', '2022-10-13 14:46:44');
INSERT INTO `visitors` VALUES (9, 17, 'sample', 'sample', 'sample', 'sample', 'sample', '2022-10-13', '09090464399', '2022-10-13 17:27:19', '2022-10-13 17:27:19');

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
