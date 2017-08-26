/*
Navicat MySQL Data Transfer

Source Server         : CatchChu
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : express_delivery

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2017-08-26 16:39:19
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `cross_list`
-- ----------------------------
DROP TABLE IF EXISTS `cross_list`;
CREATE TABLE `cross_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `cross_id` varchar(86) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `distance` varchar(16) DEFAULT NULL,
  `direction` varchar(24) DEFAULT NULL,
  `weight` varchar(8) DEFAULT NULL,
  `width` varchar(8) DEFAULT NULL,
  `level` varchar(32) DEFAULT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of cross_list
-- ----------------------------
INSERT INTO `cross_list` VALUES ('1', '2', '0757F49F0100401352--0757F49F010040346', '公园路--康岗路', '129.321', 'East', '120', '8, 8', '45000, 45000', '112.8902142', '23.17178111');
INSERT INTO `cross_list` VALUES ('2', '3', '0750F49F017041172--0750F49F017041254', '迎宾大道中--院士路', '32.6315', 'West', '140', '20, 16', '44000, 44000', '113.0859572', '22.60075528');
INSERT INTO `cross_list` VALUES ('3', '3', '0750F49F017041172--0750F49F0170415128', '迎宾大道中--丰裕路', '276.471', 'West', '130', '20, 8', '44000, 45000', '113.0882417', '22.60135583');
INSERT INTO `cross_list` VALUES ('4', '3', '0750F49F0170412084--0750F49F017041263', '乐怡路--丰盛路', '307.063', 'SouthWest', '120', '8, 8', '45000, 45000', '113.0870906', '22.6031');
INSERT INTO `cross_list` VALUES ('5', '4', '0757F49F0120421477--0757F49F012042789', '天佑北路--南新四路', '122.83', 'South', '120', '8, 8', '45000, 45000', '113.1435872', '23.02997472');
INSERT INTO `cross_list` VALUES ('6', '4', '0757F49F012042325--0757F49F012042803', '南海大道北--南新三路', '150.837', 'NorthEast', '140', '28, 16', '51000, 44000', '113.1425556', '23.02773472');
INSERT INTO `cross_list` VALUES ('7', '4', '0757F49F01204211681--0757F49F012042803', '天佑南路--南新三路', '154.767', 'NorthWest', '130', '8, 16', '45000, 44000', '113.1442139', '23.02778417');
INSERT INTO `cross_list` VALUES ('8', '5', '0766F49F01502930--0766F49F0150293288', '泽汇路--宝定路', '57.6677', 'North', '120', '8, 8', '45000, 45000', '111.5698575', '22.76809083');
INSERT INTO `cross_list` VALUES ('9', '5', '0766F49F0150293288--0766F49F015029600866', '宝定路--大新北路', '88.6502', 'NorthWest', '120', '8, 8', '45000, 45000', '111.5707594', '22.7681925');
INSERT INTO `cross_list` VALUES ('10', '5', '0766F49F01502929--0766F49F01502930', '人民中路--泽汇路', '204.583', 'North', '120', '8, 8', '45000, 45000', '111.5695903', '22.76679139');
INSERT INTO `cross_list` VALUES ('11', '6', '0766F49F01602948--0766F49F016029732', '迎宾三路--兴华三路', '357.067', 'West', '180', '12, 16', '42000, 42000', '111.5877792', '22.742245');
INSERT INTO `cross_list` VALUES ('12', '6', '0766F49F015029205--0766F49F016029732', '人民南路--兴华三路', '397.462', 'West', '150', '8, 16', '45000, 42000', '111.5881042', '22.74338194');
INSERT INTO `cross_list` VALUES ('13', '6', '0766F49F01602948--0766F49F0160295', '迎宾三路--龙都西路', '445.369', 'South', '150', '12, 4', '42000, 45000', '111.5833342', '22.74651056');
INSERT INTO `cross_list` VALUES ('14', '7', '0758F49F012036266--0758F49F012036731', '大桥路--鸿苑东街', '262.354', 'West', '130', '16, 8', '51000, 45000', '112.4462314', '23.05605472');
INSERT INTO `cross_list` VALUES ('15', '7', '0758F49F012036266--0758F49F012036874', '大桥路--端州七路', '317.412', 'NorthWest', '140', '16, 20', '51000, 44000', '112.4462933', '23.05428194');
INSERT INTO `cross_list` VALUES ('16', '7', '0758F49F012036266--0758F49F0120363774', '大桥路--黄塘路', '348.134', 'SouthWest', '140', '16, 20', '51000, 44000', '112.4459053', '23.05818833');
INSERT INTO `cross_list` VALUES ('17', '8', '0753G50F0450177022--0753G50F04501870', '梅江大道--新中路', '82.0227', 'West', '140', '8, 8', '44000, 44000', '116.1230272', '24.28848333');
INSERT INTO `cross_list` VALUES ('18', '8', '0753G50F0450174139--0753G50F04501870', '梅江三路--新中路', '82.0227', 'West', '130', '12, 8', '45000, 44000', '116.1230272', '24.28848333');
INSERT INTO `cross_list` VALUES ('19', '8', '0753G50F0450174139--0753G50F0450177022', '梅江三路--梅江大道', '82.0227', 'West', '130', '12, 8', '45000, 44000', '116.1230272', '24.28848333');
INSERT INTO `cross_list` VALUES ('20', '9', '0750F49F0210394050--0750F49F0210394051', '中山三巷--中山路', '47.2389', 'NorthEast', '120', '8, 8', '45000, 45000', '112.7938753', '22.25153111');
INSERT INTO `cross_list` VALUES ('21', '9', '0750F49F0210394--0750F49F0210394050', '学宫路--中山三巷', '63.9344', 'NorthWest', '120', '8, 8', '45000, 45000', '112.7944858', '22.25150139');
INSERT INTO `cross_list` VALUES ('22', '9', '0750F49F0210394--0750F49F0210399', '学宫路--环城北路', '80.6645', 'SouthWest', '120', '8, 4', '45000, 45000', '112.7946153', '22.25243194');
INSERT INTO `cross_list` VALUES ('23', '10', '0750F49F026036360--0750F49F026036361', '西康大道--宴都路一巷', '262.517', 'SouthEast', '120', '4, 4', '45000, 45000', '112.4958703', '21.85730472');
INSERT INTO `cross_list` VALUES ('24', '10', '0750F49F026036360--0750F49F026038114', '西康大道--365省道', '262.517', 'SouthEast', '130', '4, 12', '45000, 51000', '112.4958703', '21.85730472');
INSERT INTO `cross_list` VALUES ('25', '10', '0750F49F026036236--0750F49F026036361', '宴都路--宴都路一巷', '262.517', 'SouthEast', '130', '8, 4', '51000, 45000', '112.4958703', '21.85730472');
INSERT INTO `cross_list` VALUES ('26', '11', '020F49F0100434109--020F49F01004375', '飞云东街--机场路', '76.2888', 'South', '130', '8, 20', '45000, 51000', '113.2596494', '23.17836194');
INSERT INTO `cross_list` VALUES ('27', '11', '020F49F01004321913--020F49F01004375', '飞云西街--机场路', '85.1886', 'SouthEast', '110', '4, 20', '49, 51000', '113.2593614', '23.17834139');
INSERT INTO `cross_list` VALUES ('28', '11', '020F49F01004314148--020F49F01004375', '松云街--机场路', '96.9773', 'NorthEast', '130', '8, 20', '45000, 51000', '113.2593775', '23.17690917');
INSERT INTO `cross_list` VALUES ('29', '12', '0758F49F012036104--0758F49F012036203', '府前大街--上岸路', '81.2017', 'SouthEast', '130', '12, 4', '51000, 45000', '112.4572111', '23.02591833');
INSERT INTO `cross_list` VALUES ('30', '12', '0758F49F012036104--0758F49F0120366106', '府前大街--南亭路', '175.44', 'East', '130', '12, 8', '51000, 45000', '112.4561106', '23.02579583');
INSERT INTO `cross_list` VALUES ('31', '12', '0758F49F012036104--0758F49F012036179', '府前大街--新城路', '188.327', 'West', '130', '12, 8', '51000, 45000', '112.4595133', '23.02594222');

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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4;

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
INSERT INTO `entrances` VALUES ('8', '7', '15', '23.055946', '112.443196');
INSERT INTO `entrances` VALUES ('9', '7', '18', '23.054073', '112.446746');
INSERT INTO `entrances` VALUES ('10', '8', '19', '24.288225', '116.122709');
INSERT INTO `entrances` VALUES ('11', '8', '20', '24.289193', '116.122907');
INSERT INTO `entrances` VALUES ('12', '8', '22', '24.288225', '116.122709');
INSERT INTO `entrances` VALUES ('13', '8', '23', '24.288223', '116.122712');
INSERT INTO `entrances` VALUES ('14', '9', '24', '22.251392', '112.793889');
INSERT INTO `entrances` VALUES ('15', '9', '25', '22.251537', '112.794056');
INSERT INTO `entrances` VALUES ('16', '9', '26', '22.251537', '112.794056');
INSERT INTO `entrances` VALUES ('17', '9', '27', '22.251537', '112.794056');
INSERT INTO `entrances` VALUES ('18', '9', '28', '22.251321', '112.794475');
INSERT INTO `entrances` VALUES ('19', '11', '36', '23.178339', '113.259813');
INSERT INTO `entrances` VALUES ('20', '11', '37', '23.177105', '113.260619');
INSERT INTO `entrances` VALUES ('21', '12', '39', '23.025858', '112.457910');
INSERT INTO `entrances` VALUES ('22', '12', '40', '23.025858', '112.457910');
INSERT INTO `entrances` VALUES ('23', '12', '41', '23.025858', '112.457910');

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
-- Table structure for `forecast`
-- ----------------------------
DROP TABLE IF EXISTS `forecast`;
CREATE TABLE `forecast` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city_id` int(11) DEFAULT NULL,
  `date` varchar(32) DEFAULT NULL,
  `sunrise` varchar(16) DEFAULT NULL,
  `sunset` varchar(16) DEFAULT NULL,
  `high` varchar(32) DEFAULT NULL,
  `low` varchar(32) DEFAULT NULL,
  `aqi` float DEFAULT NULL,
  `fx` varchar(32) DEFAULT NULL,
  `fl` varchar(32) DEFAULT NULL,
  `type` varchar(32) DEFAULT NULL,
  `notice` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of forecast
-- ----------------------------
INSERT INTO `forecast` VALUES ('1', '1', '26日星期六', '06:07', '18:51', '高温 34.0℃', '低温 25.0℃', '51', '无持续风向', '<3级', '阵雨', '今日有短时阵雨，外出请携带雨具');
INSERT INTO `forecast` VALUES ('2', '1', '27日星期日', '06:08', '18:50', '高温 30.0℃', '低温 25.0℃', '38', '东风', '6-7级', '暴雨', '尽量减少户外活动，防止意外发生');
INSERT INTO `forecast` VALUES ('3', '1', '28日星期一', '06:08', '18:49', '高温 27.0℃', '低温 24.0℃', '28', '无持续风向', '<3级', '暴雨', '好好休息一天吧，是个能好好看书的日子');
INSERT INTO `forecast` VALUES ('4', '1', '29日星期二', '06:08', '18:48', '高温 29.0℃', '低温 25.0℃', '27', '无持续风向', '<3级', '中雨', '现正下雨,地湿路滑,请小心开车');
INSERT INTO `forecast` VALUES ('5', '1', '30日星期三', '06:09', '18:47', '高温 32.0℃', '低温 26.0℃', '24', '无持续风向', '<3级', '雷阵雨', '雷雨闪电时，应切断电器电源，以免损坏');

-- ----------------------------
-- Table structure for `loc`
-- ----------------------------
DROP TABLE IF EXISTS `loc`;
CREATE TABLE `loc` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `country` varchar(64) DEFAULT NULL,
  `code` int(11) DEFAULT NULL,
  `desc` varchar(64) DEFAULT NULL,
  `province` varchar(32) DEFAULT NULL,
  `province_ad_code` varchar(8) DEFAULT NULL,
  `city` varchar(64) DEFAULT NULL,
  `city_ad_code` varchar(8) DEFAULT NULL,
  `district` varchar(64) DEFAULT NULL,
  `district_ad_code` varchar(8) DEFAULT NULL,
  `gc` varchar(64) DEFAULT NULL,
  `ad_code` varchar(8) DEFAULT NULL,
  `pos` varchar(255) DEFAULT NULL,
  `area_code` varchar(8) DEFAULT NULL,
  `tel` varchar(32) DEFAULT NULL,
  `time_stamp` varchar(32) DEFAULT NULL,
  `result` varchar(5) DEFAULT NULL,
  `message` varchar(32) DEFAULT NULL,
  `version` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of loc
-- ----------------------------
INSERT INTO `loc` VALUES ('2', '中国', '1', '广东省,佛山市,三水区', '广东省', '440000', '佛山市', '440600', '三水区', '440607', '佛山市三水区西南公园', '440607', '在国美电器(佛山三水店)附近, 在公园路旁边, 靠近康岗路--张边路路口', '0757', '0757', '1503713986.97', 'true', 'Successful.', '2.0-3.0.7236.1851');
INSERT INTO `loc` VALUES ('3', '中国', '1', '广东省,江门市,蓬江区', '广东省', '440000', '江门市', '440700', '蓬江区', '440703', '江门市蓬江区五邑大学', '440703', '在五邑大学环境科学与工程研究所附近, 在丰裕路旁边, 靠近迎宾大道中--院士路路口', '0750', '0750', '1503714285.33', 'true', 'Successful.', '2.0-3.0.7236.1851');
INSERT INTO `loc` VALUES ('4', '中国', '1', '广东省,佛山市,南海区', '广东省', '440000', '佛山市', '440600', '南海区', '440605', '佛山市南海区夏北幼儿园', '440605', '在佛山市南海区人民政府附近, 在南新四路旁边, 靠近南海大道北--南新三路路口', '0757', '0757', '1503714658.18', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('5', '中国', '1', '广东省,云浮市,罗定市', '广东省', '440000', '云浮市', '445300', '罗定市', '445381', '罗定市大润发', '445381', '在罗定市政协附近, 在泽汇路旁边, 靠近泽汇路--宝定路路口', '0766', '0766', '1503714734.26', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('6', '中国', '1', '广东省,云浮市,罗定市', '广东省', '440000', '云浮市', '445300', '罗定市', '445381', '罗定实验中学', '445381', '在罗定实验中学附近, 在人民南路旁边, 靠近迎宾三路--兴华三路路口', '0766', '0766', '1503714805.31', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('7', '中国', '1', '广东省,肇庆市,端州区', '广东省', '440000', '肇庆市', '441200', '端州区', '441202', '肇庆桥西客运站', '441202', '在沙街桥西综合商场附近, 在鸿苑东街旁边, 靠近大桥路--端州七路路口', '0758', '0758', '1503717768.97', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('8', '中国', '1', '广东省,梅州市,梅江区', '广东省', '440000', '梅州市', '441400', '梅江区', '441402', '梅州茶山', '441402', '在梅州市人民政府附近, 在梅江三路旁边, 靠近梅江大道--新中路路口', '0753', '0753', '1503718267.15', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('9', '中国', '1', '广东省,江门市,台山市', '广东省', '440000', '江门市', '440700', '台山市', '440781', '江门市台山', '440781', '在台山市人民政府附近, 在中山路旁边, 靠近中山三巷--中山路路口', '0750', '0750', '1503718752.32', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('10', '中国', '1', '广东省,江门市,台山市', '广东省', '440000', '江门市', '440700', '台山市', '440781', '江门市台山汶村镇', '440781', '在台山市公安局交通警察大队汶村中队附近, 在宴都路一巷旁边, 靠近西康大道--365省道路口', '0750', '0750', '1503718782.31', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('11', '中国', '1', '广东省,广州市,白云区', '广东省', '440000', '广州市', '440100', '白云区', '440111', '广州市白云机场', '440111', '在云港大厦(飞云东街)附近, 在飞云东街旁边, 靠近飞云东街--机场路路口', '020', '020', '1503720181.5', 'true', 'Successful.', '2.0-3.0.7236.1701');
INSERT INTO `loc` VALUES ('12', '中国', '1', '广东省,肇庆市,高要区', '广东省', '440000', '肇庆市', '441200', '高要区', '441204', '肇庆市高要市湖西路湖西苑街', '441204', '在高要区规划展示馆附近, 在上岸路旁边, 靠近府前大街--上岸路路口', '0758', '0758', '1503722442.26', 'true', 'Successful.', '2.0-3.0.7236.1701');

-- ----------------------------
-- Table structure for `poi_list`
-- ----------------------------
DROP TABLE IF EXISTS `poi_list`;
CREATE TABLE `poi_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `poi_id` varchar(32) DEFAULT NULL,
  `name` varchar(86) DEFAULT NULL,
  `address` varchar(128) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `distance` varchar(16) DEFAULT NULL,
  `direction` varchar(24) DEFAULT NULL,
  `tel` varchar(64) DEFAULT NULL,
  `weight` varchar(8) DEFAULT NULL,
  `type_code` varchar(8) NOT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`,`type_code`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of poi_list
-- ----------------------------
INSERT INTO `poi_list` VALUES ('1', '3', 'B02EE0NYZG', '益华购物广场(迎宾大道中)', '迎宾大道中118号2幢益华百货地下1层', '购物服务;商场;购物中心', '152', 'NorthWest', '', '0.0', '060101', '113.084977', '22.601902');
INSERT INTO `poi_list` VALUES ('2', '3', 'B02EE02EI9', '中国建设银行(江门迎宾支行)', '迎宾大道中116号101(益华购物广场)', '金融保险服务;银行;中国建设银行', '188', 'North', '0750-3919340', '0.0', '160106', '113.085108', '22.602303');
INSERT INTO `poi_list` VALUES ('3', '3', 'B02EE01S2A', '五邑大学环境科学与工程研究所', '迎宾大道中99号附近', '科教文化服务;科研机构;科研机构', '73', 'South', '', '0.0', '141300', '113.085608', '22.600021');
INSERT INTO `poi_list` VALUES ('4', '3', 'B0FFF9WSNE', '广东高校轻化工清洁生产工程技术研究中心', '迎宾大道中', '科教文化服务;科研机构;科研机构', '74', 'South', '', '0.0', '141300', '113.085603', '22.600015');
INSERT INTO `poi_list` VALUES ('5', '3', 'B0FFGDEEWE', '香港台山商会大楼', '迎宾大道中附近', '商务住宅;楼宇;商务写字楼', '160', 'SouthEast', '', '0.0', '120201', '113.086901', '22.599819');
INSERT INTO `poi_list` VALUES ('6', '4', 'B02F5076FX', '佛山市南海区人民政府', '桂城街道南海大道88号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '14', 'NorthEast', '0757-86332422', '0.0', '130104', '113.143441', '23.028956');
INSERT INTO `poi_list` VALUES ('7', '4', 'B02F5054A4', '南海区人大常委会', '南海大道北88号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '14', 'NorthEast', '0757-86332820', '0.0', '130104', '113.143441', '23.028956');
INSERT INTO `poi_list` VALUES ('8', '4', 'B02F5073V2', '南海区人事局', '南海大道北80号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '14', 'NorthEast', '0757-86337337', '0.0', '130104', '113.143441', '23.028956');
INSERT INTO `poi_list` VALUES ('9', '6', 'B02FE00TYM', '罗定实验中学', '罗城镇迎宾路', '科教文化服务;学校;中学', '0', 'West', '0766-3925710', '0.0', '141202', '111.584325', '22.742610');
INSERT INTO `poi_list` VALUES ('10', '6', 'B0FFFZRR3J', '罗定市公安局素龙派出所罗定实验中学警务站', '罗城镇迎宾路罗定实验中学附近', '政府机构及社会团体;公检法机构;公检法机关', '19', 'NorthWest', '', '0.0', '130500', '111.584225', '22.742765');
INSERT INTO `poi_list` VALUES ('11', '6', 'B0FFG9ZCBY', '普庆行贸易有限公司', '迎宾三路霞垌108号', '购物服务;家居建材市场;家具城', '213', 'NorthEast', '0766-3866998', '0.0', '060602', '111.585700', '22.744043');
INSERT INTO `poi_list` VALUES ('12', '6', 'B0FFFYTR20', '一方名轩', '龙城中路旧公路366号', '商务住宅;住宅区;住宅小区', '250', 'East', '', '0.0', '120302', '111.586623', '22.743354');
INSERT INTO `poi_list` VALUES ('13', '6', 'B02FE0NPDB', '皇朝家具广场', '兴华三路与迎宾三路交叉口西南150米', '购物服务;家居建材市场;家具城', '308.939', 'East', '', '0.140276', '060602', '111.587294', '22.742167');
INSERT INTO `poi_list` VALUES ('14', '7', 'B02F60Q441', '上海城商住小区2期', '端州七路东50米', '商务住宅;住宅区;住宅小区', '209', 'NorthWest', '', '0.0', '120302', '112.442048', '23.056963');
INSERT INTO `poi_list` VALUES ('15', '7', 'B02F60PZ78', '沙街桥西综合商场', '端州七路西50米', '购物服务;商场;普通商场', '50', 'West', '', '0.0', '060102', '112.443184', '23.055850');
INSERT INTO `poi_list` VALUES ('16', '7', 'B02F60PZ2X', '端州农村商业银行(三村支行)', '端州七路上海城一期商铺1-5卡', '金融保险服务;银行;农村商业银行', '345.787', 'NorthWest', '0758-2871420', '0.238167', '160121', '112.440836', '23.057500');
INSERT INTO `poi_list` VALUES ('17', '7', 'B02F60PZ6S', '新龙城酒店', '端州七路36号(上海城对面)', '住宿服务;宾馆酒店;四星级宾馆', '385.095', 'NorthWest', '0758-2326888', '0.48', '100103', '112.440256', '23.057258');
INSERT INTO `poi_list` VALUES ('18', '7', 'B0FFG7QSJ7', '肇庆海关', '端州七路13号之一', '政府机构及社会团体;政府机关;地市级政府及事业单位', '376.016', 'East', '0758-2966888', '0.215613', '130103', '112.447157', '23.054725');
INSERT INTO `poi_list` VALUES ('19', '8', 'B02F102B51', '梅州市人民政府', '新中路政府大楼', '政府机构及社会团体;政府机关;地市级政府及事业单位', '25', 'NorthWest', '0753-2250962', '0.0', '130103', '116.122022', '24.288742');
INSERT INTO `poi_list` VALUES ('20', '8', 'B02F102PVM', '梅州市人大常委会', '梅江大道88号', '政府机构及社会团体;政府机关;地市级政府及事业单位', '65', 'NorthEast', '0753-2188319', '0.0', '130103', '116.122499', '24.289146');
INSERT INTO `poi_list` VALUES ('21', '8', 'B02F10S0S1', '梅州人大', '梅江大道88号', '政府机构及社会团体;政府机关;地市级政府及事业单位', '85', 'NorthEast', '0753-2188319', '0.0', '130103', '116.122794', '24.289185');
INSERT INTO `poi_list` VALUES ('22', '8', 'B02F104ORZ', '梅州市法制局', '新中路', '政府机构及社会团体;政府机关;地市级政府及事业单位', '29', 'East', '0753-2248224', '0.0', '130103', '116.122523', '24.288578');
INSERT INTO `poi_list` VALUES ('23', '8', 'B02F104OS7', '梅州市发展和改革局', '新中路', '政府机构及社会团体;政府机关;地市级政府及事业单位', '30', 'East', '0753-2250266', '0.0', '130103', '116.122526', '24.288576');
INSERT INTO `poi_list` VALUES ('24', '9', 'B02EE0OGTR', '台山市档案馆', '中山路21号', '科教文化服务;档案馆;档案馆', '67', 'SouthWest', '', '0.0', '140900', '112.793725', '22.251400');
INSERT INTO `poi_list` VALUES ('25', '9', 'B02EE00AT5', '台山市人民政府', '中山路23号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '0', 'NorthEast', '', '0.0', '130104', '112.794065', '22.251924');
INSERT INTO `poi_list` VALUES ('26', '9', 'B02EE0O97Y', '中共台山市委员会', '中山路23号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '2', 'North', '0750-5529465', '0.0', '130104', '112.794056', '22.251939');
INSERT INTO `poi_list` VALUES ('27', '9', 'B02EE0OLS4', '中共台山市纪律检查委员会', '中山路23号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '2', 'NorthWest', '0750-5525869', '0.0', '130104', '112.794044', '22.251937');
INSERT INTO `poi_list` VALUES ('28', '9', 'B02EE0OB63', '溯源月刊社', '中山路24号后座', '科教文化服务;传媒机构;报社', '71', 'South', '', '0.0', '141103', '112.794303', '22.251315');
INSERT INTO `poi_list` VALUES ('29', '10', 'B02EE01PMA', '中国邮政储蓄银行(汶村营业部)', '汶村镇宴都路一巷2号', '金融保险服务;银行;中国邮政储蓄银行', '210', 'NorthWest', '', '0.0', '160139', '112.495933', '21.856817');
INSERT INTO `poi_list` VALUES ('30', '10', 'B02EE028FG', '台山市公安局交通警察大队汶村中队', '汶村镇双龙路', '政府机构及社会团体;交通车辆管理;交通管理机构', '199', 'NorthEast', '0750-5712337', '0.0', '130601', '112.497577', '21.856795');
INSERT INTO `poi_list` VALUES ('31', '10', 'B02EE028FE', '万利楼', '宴都路与宴都路一巷交叉口西100米', '商务住宅;楼宇;商务写字楼', '304.47', 'NorthWest', '', '0.204635', '120201', '112.494963', '21.857227');
INSERT INTO `poi_list` VALUES ('32', '10', 'B02EE01GRI', '西康', '台山市', '地名地址信息;普通地名;村庄级地名', '201', 'NorthWest', '', '0.0', '190108', '112.495544', '21.856469');
INSERT INTO `poi_list` VALUES ('33', '10', 'B02EE01PM8', '西康大楼', '365省道北200米', '商务住宅;楼宇;商务写字楼', '384.206', 'North', '', '0.082422', '120201', '112.497876', '21.858454');
INSERT INTO `poi_list` VALUES ('34', '11', 'B00141JL2A', '中国工商银行(广州机场支行)', '机场路云港大厦A座1层', '金融保险服务;银行;中国工商银行', '47', 'NorthEast', '020-86126537', '0.0', '160105', '113.260099', '23.178012');
INSERT INTO `poi_list` VALUES ('35', '11', 'B00140NTZW', '中国民生银行(广州白云支行)', '机场路284号机场综合楼', '金融保险服务;银行;中国民生银行', '116', 'North', '020-86122085', '0.0', '160112', '113.259789', '23.178734');
INSERT INTO `poi_list` VALUES ('36', '11', 'B00140UGBS', '云港大厦(飞云东街)', '机场路282号', '商务住宅;楼宇;商务写字楼', '32', 'NorthEast', '', '0.0', '120201', '113.260029', '23.177887');
INSERT INTO `poi_list` VALUES ('37', '11', 'B0FFFGV9KL', '中国南方航空培训部', '飞云东街8号', '科教文化服务;培训机构;培训机构', '118', 'SouthEast', '', '0.0', '141400', '113.260847', '23.177238');
INSERT INTO `poi_list` VALUES ('38', '11', 'B00140NU05', '中国民用航空中南地区管理局公安局', '机场路288号', '政府机构及社会团体;公检法机构;公安警察', '184', 'North', '020-86635171', '0.0', '130501', '113.260040', '23.179336');
INSERT INTO `poi_list` VALUES ('39', '12', 'B02F60Q3DB', '高要区规划展示馆', '府前大街25号', '科教文化服务;展览馆;展览馆', '24', 'SouthEast', '', '0.0', '140200', '112.457982', '23.025301');
INSERT INTO `poi_list` VALUES ('40', '12', 'B02F6014E1', '中共高要区纪律检查委员会', '南岸街道府前大街25号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '24', 'SouthEast', '0758-8392645', '0.0', '130104', '112.457982', '23.025301');
INSERT INTO `poi_list` VALUES ('41', '12', 'B02F601JO8', '肇庆市高要区人大常委会', '南岸街道府前大街25号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '24', 'SouthEast', '0758-8392123', '0.0', '130104', '112.457982', '23.025301');
INSERT INTO `poi_list` VALUES ('42', '12', 'B02F6014E3', '肇庆市高要区人民政府', '府前大街25号', '政府机构及社会团体;政府机关;区县级政府及事业单位', '49', 'SouthEast', '0758-8392959', '0.0', '130104', '112.457957', '23.024991');
INSERT INTO `poi_list` VALUES ('43', '12', 'B02F6012A4', '中国建设银行(高要支行)', '南岸府前大街116号1-2层', '金融保险服务;银行;中国建设银行', '85', 'North', '0758-8392875', '0.0', '160106', '112.457515', '23.026133');

-- ----------------------------
-- Table structure for `road_list`
-- ----------------------------
DROP TABLE IF EXISTS `road_list`;
CREATE TABLE `road_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `loc_id` int(11) DEFAULT NULL,
  `road_id` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `distance` varchar(24) DEFAULT NULL,
  `direction` varchar(32) DEFAULT NULL,
  `width` varchar(8) DEFAULT NULL,
  `level` varchar(8) DEFAULT NULL,
  `longitude` varchar(32) DEFAULT NULL,
  `latitude` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of road_list
-- ----------------------------
INSERT INTO `road_list` VALUES ('1', '3', '0750F49F017041172', '迎宾大道中', '22', 'East', '20', '4', '113.085', '22.6006');
INSERT INTO `road_list` VALUES ('2', '3', '0750F49F017041254', '院士路', '32', 'West', '16', '4', '113.086', '22.6008');
INSERT INTO `road_list` VALUES ('3', '3', '0750F49F0170415128', '丰裕路', '276', 'West', '8', '5', '113.088', '22.6014');
INSERT INTO `road_list` VALUES ('4', '4', '0757F49F012042325', '南海大道北', '81', 'East', '28', '4', '113.143', '23.0287');
INSERT INTO `road_list` VALUES ('5', '4', '0757F49F012042789', '南新四路', '118', 'South', '8', '5', '113.143', '23.03');
INSERT INTO `road_list` VALUES ('6', '4', '0757F49F0120421477', '天佑北路', '122', 'South', '8', '5', '113.144', '23.03');
INSERT INTO `road_list` VALUES ('7', '5', '0766F49F01502930', '泽汇路', '53', 'NorthEast', '8', '5', '111.57', '22.7683');
INSERT INTO `road_list` VALUES ('8', '5', '0766F49F0150293288', '宝定路', '72', 'NorthWest', '8', '5', '111.571', '22.7682');
INSERT INTO `road_list` VALUES ('9', '5', '0766F49F015029600866', '大新北路', '88', 'NorthWest', '8', '5', '111.571', '22.7682');
INSERT INTO `road_list` VALUES ('10', '6', '0766F49F01602948', '迎宾三路', '188', 'SouthWest', '12', '2', '111.586', '22.7437');
INSERT INTO `road_list` VALUES ('11', '6', '0766F49F016029732', '兴华三路', '344', 'West', '16', '2', '111.588', '22.7427');
INSERT INTO `road_list` VALUES ('12', '6', '0766F49F015029205', '人民南路', '391', 'West', '8', '5', '111.588', '22.744');
INSERT INTO `road_list` VALUES ('13', '7', '0758F49F012036874', '端州七路', '8', 'NorthEast', '20', '4', '112.444', '23.0558');
INSERT INTO `road_list` VALUES ('14', '7', '0758F49F012036266', '大桥路', '254', 'West', '16', '4', '112.446', '23.0557');
INSERT INTO `road_list` VALUES ('15', '7', '0758F49F012036731', '鸿苑东街', '262', 'West', '8', '5', '112.446', '23.0561');
INSERT INTO `road_list` VALUES ('16', '8', '0753G50F04501870', '新中路', '44', 'NorthWest', '8', '4', '116.122', '24.2882');
INSERT INTO `road_list` VALUES ('17', '8', '0753G50F0450174139', '梅江三路', '81', 'NorthWest', '12', '5', '116.123', '24.2883');
INSERT INTO `road_list` VALUES ('18', '8', '0753G50F045017343', '梅新路', '113', 'East', '8', '5', '116.121', '24.2882');
INSERT INTO `road_list` VALUES ('19', '9', '0750F49F0210394051', '中山路', '47', 'NorthEast', '8', '5', '112.794', '22.2515');
INSERT INTO `road_list` VALUES ('20', '9', '0750F49F0210394050', '中山三巷', '47', 'NorthEast', '8', '5', '112.794', '22.2515');
INSERT INTO `road_list` VALUES ('21', '9', '0750F49F0210394', '学宫路', '48', 'West', '8', '5', '112.795', '22.2518');
INSERT INTO `road_list` VALUES ('22', '10', '0750F49F026036361', '宴都路一巷', '207', 'SouthEast', '4', '5', '112.496', '21.8566');
INSERT INTO `road_list` VALUES ('23', '10', '0750F49F026038114', '365省道', '217', 'South', '12', '4', '112.497', '21.857');
INSERT INTO `road_list` VALUES ('24', '10', '0750F49F026036360', '西康大道', '262', 'SouthEast', '4', '5', '112.496', '21.8573');
INSERT INTO `road_list` VALUES ('25', '11', '020F49F01004375', '机场路', '23', 'SouthEast', '20', '4', '113.26', '23.1778');
INSERT INTO `road_list` VALUES ('26', '11', '020F49F010043583', 'S41机场高速', '33', 'East', '16', '1', '113.259', '23.1776');
INSERT INTO `road_list` VALUES ('27', '11', '020F49F0100434109', '飞云东街', '74', 'South', '8', '5', '113.26', '23.1784');
INSERT INTO `road_list` VALUES ('28', '12', '0758F49F012036104', '府前大街', '68', 'SouthWest', '12', '4', '112.458', '23.0259');
INSERT INTO `road_list` VALUES ('29', '12', '0758F49F012036203', '上岸路', '75', 'SouthEast', '4', '5', '112.457', '23.0258');
INSERT INTO `road_list` VALUES ('30', '12', '0758F49F0120366106', '南亭路', '162', 'East', '8', '5', '112.456', '23.0254');

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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sea_area
-- ----------------------------
INSERT INTO `sea_area` VALUES ('1', '2', '', '');
INSERT INTO `sea_area` VALUES ('2', '3', '', '');
INSERT INTO `sea_area` VALUES ('3', '4', '', '');
INSERT INTO `sea_area` VALUES ('4', '5', '', '');
INSERT INTO `sea_area` VALUES ('5', '6', '', '');
INSERT INTO `sea_area` VALUES ('6', '7', '', '');
INSERT INTO `sea_area` VALUES ('7', '8', '', '');
INSERT INTO `sea_area` VALUES ('8', '9', '', '');
INSERT INTO `sea_area` VALUES ('9', '10', '', '');
INSERT INTO `sea_area` VALUES ('10', '11', '', '');
INSERT INTO `sea_area` VALUES ('11', '12', '', '');

-- ----------------------------
-- Table structure for `weather`
-- ----------------------------
DROP TABLE IF EXISTS `weather`;
CREATE TABLE `weather` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city` varchar(32) DEFAULT NULL,
  `shidu` varchar(5) DEFAULT NULL,
  `pm25` float DEFAULT NULL,
  `pm10` float DEFAULT NULL,
  `quality` varchar(16) DEFAULT NULL,
  `wendu` varchar(5) DEFAULT NULL,
  `ganmao` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of weather
-- ----------------------------
INSERT INTO `weather` VALUES ('1', '江门', '59%', '37', '65', '轻度污染', '33', '儿童、老年人及心脏、呼吸系统疾病患者人群应减少长时间或高强度户外锻炼');

-- ----------------------------
-- Table structure for `yesterday`
-- ----------------------------
DROP TABLE IF EXISTS `yesterday`;
CREATE TABLE `yesterday` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city_id` int(11) DEFAULT NULL,
  `date` varchar(32) DEFAULT NULL,
  `sunrise` varchar(16) DEFAULT NULL,
  `sunset` varchar(16) DEFAULT NULL,
  `high` varchar(32) DEFAULT NULL,
  `low` varchar(32) DEFAULT NULL,
  `aqi` float DEFAULT NULL,
  `fx` varchar(32) DEFAULT NULL,
  `fl` varchar(32) DEFAULT NULL,
  `type` varchar(32) DEFAULT NULL,
  `notice` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of yesterday
-- ----------------------------
INSERT INTO `yesterday` VALUES ('1', '1', '25日星期五', '06:07', '18:52', '高温 33.0℃', '低温 26.0℃', '49', '无持续风向', '<3级', '阵雨', '今日有短时阵雨，外出请携带雨具');
