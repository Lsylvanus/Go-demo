/*
Navicat MySQL Data Transfer

Source Server         : CatchChu
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : express_delivery

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2017-08-26 11:01:19
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `cross_list`
-- ----------------------------
DROP TABLE IF EXISTS `cross_list`;
CREATE TABLE `cross_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `distance` varchar(16) DEFAULT NULL,
  `direction` varchar(24) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `weight` varchar(8) DEFAULT NULL,
  `level` varchar(32) DEFAULT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `cross_id` varchar(64) DEFAULT NULL,
  `width` varchar(8) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of cross_list
-- ----------------------------
INSERT INTO `cross_list` VALUES ('1', '2', '129.321', 'East', '公园路--康岗路', '120', '45000, 45000', '112.8902142', '0757F49F0100401352--0757F49F010040346', '8, 8', '23.17178111');
INSERT INTO `cross_list` VALUES ('2', '3', '32.6315', 'West', '迎宾大道中--院士路', '140', '44000, 44000', '113.0859572', '0750F49F017041172--0750F49F017041254', '20, 16', '22.60075528');
INSERT INTO `cross_list` VALUES ('3', '3', '276.471', 'West', '迎宾大道中--丰裕路', '130', '44000, 45000', '113.0882417', '0750F49F017041172--0750F49F0170415128', '20, 8', '22.60135583');
INSERT INTO `cross_list` VALUES ('4', '3', '307.063', 'SouthWest', '乐怡路--丰盛路', '120', '45000, 45000', '113.0870906', '0750F49F0170412084--0750F49F017041263', '8, 8', '22.6031');
INSERT INTO `cross_list` VALUES ('5', '4', '122.83', 'South', '天佑北路--南新四路', '120', '45000, 45000', '113.1435872', '0757F49F0120421477--0757F49F012042789', '8, 8', '23.02997472');
INSERT INTO `cross_list` VALUES ('6', '4', '150.837', 'NorthEast', '南海大道北--南新三路', '140', '51000, 44000', '113.1425556', '0757F49F012042325--0757F49F012042803', '28, 16', '23.02773472');
INSERT INTO `cross_list` VALUES ('7', '4', '154.767', 'NorthWest', '天佑南路--南新三路', '130', '45000, 44000', '113.1442139', '0757F49F01204211681--0757F49F012042803', '8, 16', '23.02778417');
INSERT INTO `cross_list` VALUES ('8', '5', '57.6677', 'North', '泽汇路--宝定路', '120', '45000, 45000', '111.5698575', '0766F49F01502930--0766F49F0150293288', '8, 8', '22.76809083');
INSERT INTO `cross_list` VALUES ('9', '5', '88.6502', 'NorthWest', '宝定路--大新北路', '120', '45000, 45000', '111.5707594', '0766F49F0150293288--0766F49F015029600866', '8, 8', '22.7681925');
INSERT INTO `cross_list` VALUES ('10', '5', '204.583', 'North', '人民中路--泽汇路', '120', '45000, 45000', '111.5695903', '0766F49F01502929--0766F49F01502930', '8, 8', '22.76679139');
INSERT INTO `cross_list` VALUES ('11', '6', '357.067', 'West', '迎宾三路--兴华三路', '180', '42000, 42000', '111.5877792', '0766F49F01602948--0766F49F016029732', '12, 16', '22.742245');
INSERT INTO `cross_list` VALUES ('12', '6', '397.462', 'West', '人民南路--兴华三路', '150', '45000, 42000', '111.5881042', '0766F49F015029205--0766F49F016029732', '8, 16', '22.74338194');
INSERT INTO `cross_list` VALUES ('13', '6', '445.369', 'South', '迎宾三路--龙都西路', '150', '42000, 45000', '111.5833342', '0766F49F01602948--0766F49F0160295', '12, 4', '22.74651056');

-- ----------------------------
-- Table structure for `delivery`
-- ----------------------------
DROP TABLE IF EXISTS `delivery`;
CREATE TABLE `delivery` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `message` varchar(255) DEFAULT NULL,
  `nu` varchar(32) DEFAULT NULL,
  `ischeck` int(1) DEFAULT NULL,
  `condition` varchar(12) DEFAULT NULL,
  `com` varchar(8) NOT NULL DEFAULT 'shunfeng',
  `status` int(3) DEFAULT NULL,
  `state` int(1) DEFAULT NULL,
  `created` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `version` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of delivery
-- ----------------------------
INSERT INTO `delivery` VALUES ('1', 'ok', '687102589887', '1', 'F00', 'shunfeng', '200', '3', '2017-08-25 17:13:16', '2017-08-25 17:13:16', '1');

-- ----------------------------
-- Table structure for `entrances`
-- ----------------------------
DROP TABLE IF EXISTS `entrances`;
CREATE TABLE `entrances` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `poi_list_id` int(11) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of entrances
-- ----------------------------
INSERT INTO `entrances` VALUES ('1', '3', '3', '22.600062', '113.085777');
INSERT INTO `entrances` VALUES ('2', '3', '4', '22.600062', '113.085777');
INSERT INTO `entrances` VALUES ('3', '3', '5', '22.599768', '113.086652');
INSERT INTO `entrances` VALUES ('4', '4', '6', '23.028787', '113.142645');
INSERT INTO `entrances` VALUES ('5', '4', '7', '23.028787', '113.142645');
INSERT INTO `entrances` VALUES ('6', '4', '8', '23.028787', '113.142645');
INSERT INTO `entrances` VALUES ('7', '6', '12', '22.743221', '111.58655');

-- ----------------------------
-- Table structure for `express`
-- ----------------------------
DROP TABLE IF EXISTS `express`;
CREATE TABLE `express` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `nu` varchar(32) NOT NULL,
  `time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `ftime` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `context` text,
  `location` varchar(255) DEFAULT NULL,
  `created` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `version` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of express
-- ----------------------------
INSERT INTO `express` VALUES ('1', '687102589887', '2017-06-16 11:51:12', '2017-06-16 11:51:12', '已签收,感谢使用顺丰,期待再次为您服务', '', '2017-08-25 17:13:16', '2017-08-25 17:13:16', '1');
INSERT INTO `express` VALUES ('2', '687102589887', '2017-06-16 09:12:40', '2017-06-16 09:12:40', '快件交给吴灿辉，正在派送途中（联系电话：18566891248）', '', '2017-08-25 17:13:16', '2017-08-25 17:13:16', '1');
INSERT INTO `express` VALUES ('3', '687102589887', '2017-06-16 07:06:34', '2017-06-16 07:06:34', '快件到达 【佛山市三水区兴达路营业点】', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('4', '687102589887', '2017-06-16 00:32:12', '2017-06-16 00:32:12', '快件在【佛山顺德集散中心】已装车，准备发往 【佛山市三水区兴达路营业点】', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('5', '687102589887', '2017-06-16 00:32:12', '2017-06-16 00:32:12', '快件到达 【佛山顺德集散中心】', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('6', '687102589887', '2017-06-15 21:31:38', '2017-06-15 21:31:38', '快件在【江门江海集散中心】已装车，准备发往 【佛山顺德集散中心】', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('7', '687102589887', '2017-06-15 21:13:47', '2017-06-15 21:13:47', '快件到达 【江门江海集散中心】', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('8', '687102589887', '2017-06-15 20:02:20', '2017-06-15 20:02:20', '快件在【江门蓬江泰和广场营业部】已装车，准备发往下一站', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');
INSERT INTO `express` VALUES ('9', '687102589887', '2017-06-15 16:46:11', '2017-06-15 16:46:11', '顺丰速运 已收取快件', '', '2017-08-25 17:13:17', '2017-08-25 17:13:17', '1');

-- ----------------------------
-- Table structure for `loc`
-- ----------------------------
DROP TABLE IF EXISTS `loc`;
CREATE TABLE `loc` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `province` varchar(32) DEFAULT NULL,
  `code` int(11) DEFAULT NULL,
  `tel` varchar(8) DEFAULT NULL,
  `city_ad_code` varchar(8) DEFAULT NULL,
  `area_code` varchar(8) DEFAULT NULL,
  `time_stamp` varchar(32) DEFAULT NULL,
  `pos` varchar(255) DEFAULT NULL,
  `result` varchar(5) DEFAULT NULL,
  `message` varchar(16) DEFAULT NULL,
  `desc` varchar(64) DEFAULT NULL,
  `city` varchar(32) DEFAULT NULL,
  `district_ad_code` varchar(8) DEFAULT NULL,
  `district` varchar(32) DEFAULT NULL,
  `country` varchar(64) DEFAULT NULL,
  `province_ad_code` varchar(8) DEFAULT NULL,
  `version` varchar(32) DEFAULT NULL,
  `ad_code` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of loc
-- ----------------------------
INSERT INTO `loc` VALUES ('2', '广东省', '1', '0757', '440600', '0757', '1503713986.97', '在国美电器(佛山三水店)附近, 在公园路旁边, 靠近康岗路--张边路路口', 'true', 'Successful.', '广东省,佛山市,三水区', '佛山市', '440607', '三水区', '中国', '440000', '2.0-3.0.7236.1851', '440607');
INSERT INTO `loc` VALUES ('3', '广东省', '1', '0750', '440700', '0750', '1503714285.33', '在五邑大学环境科学与工程研究所附近, 在丰裕路旁边, 靠近迎宾大道中--院士路路口', 'true', 'Successful.', '广东省,江门市,蓬江区', '江门市', '440703', '蓬江区', '中国', '440000', '2.0-3.0.7236.1851', '440703');
INSERT INTO `loc` VALUES ('4', '广东省', '1', '0757', '440600', '0757', '1503714658.18', '在佛山市南海区人民政府附近, 在南新四路旁边, 靠近南海大道北--南新三路路口', 'true', 'Successful.', '广东省,佛山市,南海区', '佛山市', '440605', '南海区', '中国', '440000', '2.0-3.0.7236.1701', '440605');
INSERT INTO `loc` VALUES ('5', '广东省', '1', '0766', '445300', '0766', '1503714734.26', '在罗定市政协附近, 在泽汇路旁边, 靠近泽汇路--宝定路路口', 'true', 'Successful.', '广东省,云浮市,罗定市', '云浮市', '445381', '罗定市', '中国', '440000', '2.0-3.0.7236.1701', '445381');
INSERT INTO `loc` VALUES ('6', '广东省', '1', '0766', '445300', '0766', '1503714805.31', '在罗定实验中学附近, 在人民南路旁边, 靠近迎宾三路--兴华三路路口', 'true', 'Successful.', '广东省,云浮市,罗定市', '云浮市', '445381', '罗定市', '中国', '440000', '2.0-3.0.7236.1701', '445381');

-- ----------------------------
-- Table structure for `poi_list`
-- ----------------------------
DROP TABLE IF EXISTS `poi_list`;
CREATE TABLE `poi_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `distance` varchar(16) DEFAULT NULL,
  `direction` varchar(24) DEFAULT NULL,
  `tel` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `weight` varchar(8) DEFAULT NULL,
  `type_code` varchar(8) NOT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `address` varchar(128) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  `type` varchar(128) DEFAULT NULL,
  `poi_id` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`,`type_code`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of poi_list
-- ----------------------------
INSERT INTO `poi_list` VALUES ('1', '3', '152', 'NorthWest', '', '益华购物广场(迎宾大道中)', '0.0', '060101', '113.084977', '迎宾大道中118号2幢益华百货地下1层', '22.601902', '购物服务;商场;购物中心', 'B02EE0NYZG');
INSERT INTO `poi_list` VALUES ('2', '3', '188', 'North', '0750-3919340', '中国建设银行(江门迎宾支行)', '0.0', '160106', '113.085108', '迎宾大道中116号101(益华购物广场)', '22.602303', '金融保险服务;银行;中国建设银行', 'B02EE02EI9');
INSERT INTO `poi_list` VALUES ('3', '3', '73', 'South', '', '五邑大学环境科学与工程研究所', '0.0', '141300', '113.085608', '迎宾大道中99号附近', '22.600021', '科教文化服务;科研机构;科研机构', 'B02EE01S2A');
INSERT INTO `poi_list` VALUES ('4', '3', '74', 'South', '', '广东高校轻化工清洁生产工程技术研究中心', '0.0', '141300', '113.085603', '迎宾大道中', '22.600015', '科教文化服务;科研机构;科研机构', 'B0FFF9WSNE');
INSERT INTO `poi_list` VALUES ('5', '3', '160', 'SouthEast', '', '香港台山商会大楼', '0.0', '120201', '113.086901', '迎宾大道中附近', '22.599819', '商务住宅;楼宇;商务写字楼', 'B0FFGDEEWE');
INSERT INTO `poi_list` VALUES ('6', '4', '14', 'NorthEast', '0757-86332422', '佛山市南海区人民政府', '0.0', '130104', '113.143441', '桂城街道南海大道88号', '23.028956', '政府机构及社会团体;政府机关;区县级政府及事业单位', 'B02F5076FX');
INSERT INTO `poi_list` VALUES ('7', '4', '14', 'NorthEast', '0757-86332820', '南海区人大常委会', '0.0', '130104', '113.143441', '南海大道北88号', '23.028956', '政府机构及社会团体;政府机关;区县级政府及事业单位', 'B02F5054A4');
INSERT INTO `poi_list` VALUES ('8', '4', '14', 'NorthEast', '0757-86337337', '南海区人事局', '0.0', '130104', '113.143441', '南海大道北80号', '23.028956', '政府机构及社会团体;政府机关;区县级政府及事业单位', 'B02F5073V2');
INSERT INTO `poi_list` VALUES ('9', '6', '0', 'West', '0766-3925710', '罗定实验中学', '0.0', '141202', '111.584325', '罗城镇迎宾路', '22.742610', '科教文化服务;学校;中学', 'B02FE00TYM');
INSERT INTO `poi_list` VALUES ('10', '6', '19', 'NorthWest', '', '罗定市公安局素龙派出所罗定实验中学警务站', '0.0', '130500', '111.584225', '罗城镇迎宾路罗定实验中学附近', '22.742765', '政府机构及社会团体;公检法机构;公检法机关', 'B0FFFZRR3J');
INSERT INTO `poi_list` VALUES ('11', '6', '213', 'NorthEast', '0766-3866998', '普庆行贸易有限公司', '0.0', '060602', '111.585700', '迎宾三路霞垌108号', '22.744043', '购物服务;家居建材市场;家具城', 'B0FFG9ZCBY');
INSERT INTO `poi_list` VALUES ('12', '6', '250', 'East', '', '一方名轩', '0.0', '120302', '111.586623', '龙城中路旧公路366号', '22.743354', '商务住宅;住宅区;住宅小区', 'B0FFFYTR20');
INSERT INTO `poi_list` VALUES ('13', '6', '308.939', 'East', '', '皇朝家具广场', '0.140276', '060602', '111.587294', '兴华三路与迎宾三路交叉口西南150米', '22.742167', '购物服务;家居建材市场;家具城', 'B02FE0NPDB');

-- ----------------------------
-- Table structure for `road_list`
-- ----------------------------
DROP TABLE IF EXISTS `road_list`;
CREATE TABLE `road_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `distance` varchar(16) DEFAULT NULL,
  `direction` varchar(24) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `level` varchar(8) DEFAULT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `width` varchar(8) DEFAULT NULL,
  `road_id` varchar(64) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of road_list
-- ----------------------------
INSERT INTO `road_list` VALUES ('1', '3', '22', 'East', '迎宾大道中', '4', '113.085', '20', '0750F49F017041172', '22.6006');
INSERT INTO `road_list` VALUES ('2', '3', '32', 'West', '院士路', '4', '113.086', '16', '0750F49F017041254', '22.6008');
INSERT INTO `road_list` VALUES ('3', '3', '276', 'West', '丰裕路', '5', '113.088', '8', '0750F49F0170415128', '22.6014');
INSERT INTO `road_list` VALUES ('4', '4', '81', 'East', '南海大道北', '4', '113.143', '28', '0757F49F012042325', '23.0287');
INSERT INTO `road_list` VALUES ('5', '4', '118', 'South', '南新四路', '5', '113.143', '8', '0757F49F012042789', '23.03');
INSERT INTO `road_list` VALUES ('6', '4', '122', 'South', '天佑北路', '5', '113.144', '8', '0757F49F0120421477', '23.03');
INSERT INTO `road_list` VALUES ('7', '5', '53', 'NorthEast', '泽汇路', '5', '111.57', '8', '0766F49F01502930', '22.7683');
INSERT INTO `road_list` VALUES ('8', '5', '72', 'NorthWest', '宝定路', '5', '111.571', '8', '0766F49F0150293288', '22.7682');
INSERT INTO `road_list` VALUES ('9', '5', '88', 'NorthWest', '大新北路', '5', '111.571', '8', '0766F49F015029600866', '22.7682');
INSERT INTO `road_list` VALUES ('10', '6', '188', 'SouthWest', '迎宾三路', '2', '111.586', '12', '0766F49F01602948', '22.7437');
INSERT INTO `road_list` VALUES ('11', '6', '344', 'West', '兴华三路', '2', '111.588', '16', '0766F49F016029732', '22.7427');
INSERT INTO `road_list` VALUES ('12', '6', '391', 'West', '人民南路', '5', '111.588', '8', '0766F49F015029205', '22.744');

-- ----------------------------
-- Table structure for `sea_area`
-- ----------------------------
DROP TABLE IF EXISTS `sea_area`;
CREATE TABLE `sea_area` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `ad_code` varchar(8) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sea_area
-- ----------------------------
INSERT INTO `sea_area` VALUES ('1', '2', '', '');
INSERT INTO `sea_area` VALUES ('2', '3', '', '');
INSERT INTO `sea_area` VALUES ('3', '4', '', '');
INSERT INTO `sea_area` VALUES ('4', '5', '', '');
INSERT INTO `sea_area` VALUES ('5', '6', '', '');
