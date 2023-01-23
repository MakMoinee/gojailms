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

 Date: 23/01/2023 16:44:40
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of crime
-- ----------------------------
INSERT INTO `crime` VALUES (1, 'murder', NULL, 2, '2022-10-22 00:00:00', '2023-01-21 21:10:08', '2023-01-21 21:10:12');

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
INSERT INTO `inmates` VALUES (1, 1, 'Ricardo', 'Dalisay', 'X', 'Purok4A Poblacion Valencia City, Bukidnon', 'Zamboanga City', '1998-02-01 00:00:00', '2022-03-19 20:42:46', '2022-03-19 20:42:46');

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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (18, 'admin', '$2a$14$NKT9r6172lWglVqJoKKok.h9YbaI7ouXV4FPboN3DI091UOwQLMRy', 1);
INSERT INTO `users` VALUES (19, 'sample', '$2a$14$BJ9ad6H.sTaPtEPV3M/VPuOZNAUogV73InjSAWnPQ9J7S4kVdvEgC', 2);

-- ----------------------------
-- Table structure for visitorhistory
-- ----------------------------
DROP TABLE IF EXISTS `visitorhistory`;
CREATE TABLE `visitorhistory`  (
  `visitorHistoryID` int NOT NULL AUTO_INCREMENT,
  `visitorID` int NOT NULL,
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `visitedDateTime` datetime NOT NULL,
  `visitedOut` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`visitorHistoryID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitorhistory
-- ----------------------------
INSERT INTO `visitorhistory` VALUES (2, 9, 'she\'s wearing black vest', '2022-11-30 03:34:34', NULL);
INSERT INTO `visitorhistory` VALUES (3, 9, 'She\'s wearing black shirt with maong pants.', '2022-11-30 12:39:27', NULL);
INSERT INTO `visitorhistory` VALUES (4, 10, 'White shirt', '2023-01-13 08:39:48', NULL);
INSERT INTO `visitorhistory` VALUES (5, 10, 'blue cap', '2023-01-13 09:32:28', NULL);
INSERT INTO `visitorhistory` VALUES (6, 10, 'white tshirt', '2023-01-23 16:22:32', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of visitors
-- ----------------------------
INSERT INTO `visitors` VALUES (10, 19, 'Sample', 'Sample', 'X', 'Poblacion Valencia City', 'New Bataan', '1998-10-13', '09365861684', '2023-01-13 07:29:16', '2023-01-13 07:29:16');

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

-- ----------------------------
-- View structure for vwvisithistory
-- ----------------------------
DROP VIEW IF EXISTS `vwvisithistory`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `vwvisithistory` AS select `visitorhistory`.`visitorHistoryID` AS `visitorHistoryID`,`visitors`.`userID` AS `userID`,`visitors`.`firstName` AS `firstName`,`visitors`.`lastName` AS `lastName`,`visitors`.`middleName` AS `middleName`,`visitorhistory`.`remarks` AS `remarks`,`visitorhistory`.`visitedDateTime` AS `visitedDateTime`,`visitorhistory`.`visitedOut` AS `visitedOut` from (`visitorhistory` join `visitors` on((`visitorhistory`.`visitorID` = `visitors`.`visitorID`)));

SET FOREIGN_KEY_CHECKS = 1;
