-- MySQL dump 10.13  Distrib 5.7.44, for Linux (x86_64)
--
-- Host: localhost    Database: co_store
-- ------------------------------------------------------
-- Server version	5.7.44-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `feed_back`
--

DROP TABLE IF EXISTS `feed_back`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `feed_back` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8_bin NOT NULL,
  `email` varchar(255) COLLATE utf8_bin NOT NULL,
  `content` varchar(255) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='反馈';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feed_back`
--

LOCK TABLES `feed_back` WRITE;
/*!40000 ALTER TABLE `feed_back` DISABLE KEYS */;
INSERT INTO `feed_back` VALUES (1,'张杰','15506531157','15506531157@163.com','模块缺失！');
/*!40000 ALTER TABLE `feed_back` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location_goods`
--

DROP TABLE IF EXISTS `location_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location_goods` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '路径ID',
  `order_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `position` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '节点地址',
  `position_load` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='发货路径节点表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location_goods`
--

LOCK TABLES `location_goods` WRITE;
/*!40000 ALTER TABLE `location_goods` DISABLE KEYS */;
/*!40000 ALTER TABLE `location_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_form`
--

DROP TABLE IF EXISTS `order_form`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_form` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `product_id` int(10) unsigned NOT NULL,
  `to_address` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `remarks` longtext COLLATE utf8mb4_bin NOT NULL,
  `user_uid` longtext COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_form`
--

LOCK TABLES `order_form` WRITE;
/*!40000 ALTER TABLE `order_form` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_form` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `type_id` int(10) unsigned NOT NULL,
  `icon` longtext COLLATE utf8mb4_bin NOT NULL,
  `name` longtext COLLATE utf8mb4_bin NOT NULL,
  `content` longtext COLLATE utf8mb4_bin NOT NULL,
  `recommend` tinyint(3) unsigned NOT NULL,
  `not_recommended` tinyint(3) unsigned NOT NULL,
  `collect` int(10) unsigned DEFAULT NULL,
  `uid` longtext COLLATE utf8mb4_bin NOT NULL,
  `price` int(10) unsigned NOT NULL,
  `user_uid` longtext COLLATE utf8mb4_bin NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1,1,'http://localhost:89/product/76b3398c-9ac0-4811-8ca0-52c582a021ca.jpg','惠普暗影精灵7','无与伦比的性能：搭载最新一代处理器和显卡，让你畅享流畅的游戏体验，无论是挑战最高难度的游戏还是创造性的任务。',0,0,0,'6169ba87-bcd5-44f2-8b9c-68a4b3fe9785',9586,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(2,1,'http://118.31.43.239:89/product/fdd8e5ab-3e79-45d0-bc86-f5dabc16bba1.jpg','联想拯救者Y9000P','在战场的边缘，唯有真正的勇者才能立于不败之地。现在，轮到你挺身而出，展现真正的英雄风范！联想拯救者系列笔记本，不仅仅是一台电脑，更是你征服游戏世界的武器！',0,0,0,'ddb5f939-c814-4fff-a68d-207a185f8d4c',9299,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(3,1,'http://118.31.43.239:89/product/4715d7bd-ab6e-45b3-998a-c988dcad41df.jpg','苹果 MacBook Pro','苹果公司的旗舰笔记本电脑，拥有高性能处理器、视网膜显示屏和全新的蝴蝶键盘。',0,0,0,'32ba6656-4a30-43a6-9c67-827022bc55cd',10249,'d2bcfba3-68ce-44ca-a16d-b0c1cba0e7fb',1),(4,1,'http://118.31.43.239:89/product/a1e107d8-dbb4-4c73-ac3f-043bf2290f81.jpg','戴尔 XPS 13','戴尔的旗舰超极本，采用极窄边框设计，拥有高分辨率显示屏和长续航电池。',0,0,0,'82a2d375-2da4-4864-ac09-2cb7ff85b481',6946,'d2bcfba3-68ce-44ca-a16d-b0c1cba0e7fb',1),(5,1,'http://118.31.43.239:89/product/292d3c64-c344-4a24-9025-038bbe517ca8.jpg','华硕 ROG Strix GL12','华硕的游戏台式机系列，配置高性能处理器和显卡，适合游戏和多媒体内容创作。',0,0,0,'092c47fc-c0d8-43f4-8aa0-51438b856b9b',6784,'d2bcfba3-68ce-44ca-a16d-b0c1cba0e7fb',1),(12,5,'http://118.31.43.239:89/product/56c52f45-a906-4bd5-b21c-f0625b9b4ac4.png','li','test',0,0,0,'dafb1f95-e6fb-46f1-9a20-226d0a50f9a1',12,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(7,3,'http://118.31.43.239:89/product/11bd9ff1-b222-4de7-b6d6-c6932be1ccf6.png','华为Meta60Pro','华为Mate 60凭借其卓越的性能与创新设计，成为市场上的佼佼者。它搭载了最新的麒麟处理器，提供超强的计算能力和卓越的能效表现，确保了用户在多任务处理和游戏娱乐中的流畅体验。Mate 60配备了先进的摄影系统，包括高像素主摄像头和多种拍摄模式，使用户能够轻松捕捉高清照片和视频。此外，它的电池续航能力出色，并支持快速充电技术，极大提升了使用便捷性。其精致的外观设计与高质量的材质，也为用户带来了极佳的手感和视觉享受。总之，华为Mate 60在性能、摄影、续航和设计等各方面都有显著优势，成为一款备受推崇的智能手机。',0,0,0,'1a3f57c8-bf0e-47eb-ad06-3b34ac9ec1c6',6999,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(8,2,'http://118.31.43.239:89/product/335256c6-1102-487a-8e00-a6dded5c954a.png','美菱-506WP9BDZ曙光锦','506WP9BDZ曙光锦冰箱具备大容量存储空间，能够满足全家人的食物储存需求；其智能控温功能确保了食物的新鲜度和保存效果；同时，高效节能设计不仅减少了电力消耗，还降低了使用成本，使其成为家庭中可靠又经济的选择。',0,0,0,'98fdda24-7af2-42c8-8ad9-8109ebfc6e34',9999,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(9,8,'http://118.31.43.239:89/product/200b0e9f-ab1b-4e83-86fc-d0c637a601c7.jpg','《三体》','《三体》是一部具有深刻思想内涵和科学性的优秀科幻小说，通过深入探讨人类文明、科学探索、以及宇宙的奥秘，展示了作者刘慈欣丰富的想象力和思考深度。小说不仅探讨了物理学、计算机科学、天文学等科学问题，还深刻反思了人类文明的演化和未来的发展，引导读者思考如何应对未来的挑战。',0,0,0,'d35d09fc-814d-4fc3-a5c7-1c71c919288d',47,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1),(10,4,'http://118.31.43.239:89/product/c1ecb5bf-c8ad-431e-b3da-deb1e8bc9c39.png','罗蒙西装','罗蒙西装以其精湛的裁剪工艺和高品质的面料著称，融合经典与现代设计元素，展现出优雅与时尚的完美平衡。每一件罗蒙西装都经过精细的制作流程，确保良好的版型和舒适的穿着体验，适合各种正式场合和商务需求，是追求品位和品质人士的不二之选。',0,0,0,'73afc58c-eaf9-4ff8-8f51-aa03e8e55350',289,'d2bcfba3-68ce-44ca-a16d-b0c1cba0e7fb',1),(11,9,'http://118.31.43.239:89/product/8b410b1f-23f6-4a15-87f7-e435f3801116.jpg','翡翠手镯','翡翠手镯因其晶莹剔透的质地和丰富多彩的颜色而备受青睐，象征着高贵、吉祥与长寿。它不仅是一种美丽的饰品，更是文化与历史的传承，每一款手镯都独具特色，蕴含着天然的灵气与独特的韵味。佩戴翡翠手镯，不仅能够彰显个人品位，还能带来一份宁静与祥和，是馈赠亲友和珍藏自用的绝佳选择。',0,0,0,'c0e644fc-ef05-405b-983b-64c06e619865',365098,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0',1);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_banner`
--

DROP TABLE IF EXISTS `product_banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_banner` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `banner` longtext COLLATE utf8mb4_bin NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `uid` longtext COLLATE utf8mb4_bin NOT NULL,
  `product_name` longtext COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_banner`
--

LOCK TABLES `product_banner` WRITE;
/*!40000 ALTER TABLE `product_banner` DISABLE KEYS */;
INSERT INTO `product_banner` VALUES (1,'http://118.31.43.239:89/product/4a454109-24b8-4725-bc8e-dd23a1eb64c4.jpg',1,'0c49e912-b7f4-463b-8638-1df4c62b96f7','惠普暗影精灵7'),(2,'http://118.31.43.239:89/product/ae23522f-06b7-4f5a-9a1b-ae00d9166a2d.jpg',2,'bac20fdc-335f-40cb-a9e6-f27f1aec5568','联想拯救者Y9000P'),(3,'http://118.31.43.239:89/product/993751f6-7360-4014-8763-abe6340cfa05.jpg',3,'3a7ea100-a7c3-4355-bc06-a2e0f03ffbb8','苹果 MacBook Pro'),(4,'http://118.31.43.239:89/product/c5392a00-2555-4ad6-9d50-8d20371f55a1.jpg',4,'80cd77c2-1633-4334-aab6-7afe5a69fba9','戴尔 XPS 13'),(5,'http://118.31.43.239:89/product/84394c9e-21c9-4af5-8531-bd667dc5abc2.jpg',5,'0ccde1e8-641d-4dde-9b6e-41f6ee727c49','华硕 ROG Strix GL12'),(12,'http://118.31.43.239:89/product/c4c239a2-9f19-4e76-bf04-e5dd3e952b3d.png',12,'f05709f9-8e47-474c-aaeb-aff87ff0ae0a','li'),(7,'http://118.31.43.239:89/product/4ebd0dbb-29e3-4458-9595-5de0ef4ccd18.png',7,'807bf080-6aca-44ea-87e6-df6bc6acc691','华为Meta60Pro'),(8,'http://118.31.43.239:89/product/79fceb3f-7a08-499d-94c6-db1d412323c4.jpg',8,'9bfc6946-de35-4565-ae6e-d058f72db174','美菱-506WP9BDZ曙光锦'),(9,'http://118.31.43.239:89/product/acfb6db6-4df6-4fe2-9a56-ab419d5d8e56.jpg',9,'29f64f26-b570-4f6b-9b56-5deaeac06c5f','《三体》'),(10,'http://118.31.43.239:89/product/2e71160d-0da3-494a-8a4a-dd6d6eea6ee2.png',10,'eab8d067-c663-4832-b150-4ebe11bb867f','罗蒙西装'),(11,'http://118.31.43.239:89/product/c0f992f8-6651-473d-a924-0bc59b45534b.png',11,'673c116f-e8d9-4772-be90-e591f5e26a5b','翡翠手镯'),(13,'http://118.31.43.239:89/product/8d99d338-4091-4b7c-b541-ce4a7767198d.png',13,'06f4710e-e3a3-4854-b451-95bc67c6d74c','谁');
/*!40000 ALTER TABLE `product_banner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_evaluate`
--

DROP TABLE IF EXISTS `product_evaluate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_evaluate` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '商品评价ID',
  `product_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `content` longtext COLLATE utf8mb4_bin NOT NULL,
  `icon` longtext COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商品评价';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_evaluate`
--

LOCK TABLES `product_evaluate` WRITE;
/*!40000 ALTER TABLE `product_evaluate` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_evaluate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_type`
--

DROP TABLE IF EXISTS `product_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_type` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '商品类型ID',
  `root` int(10) unsigned NOT NULL DEFAULT '0',
  `type` longtext COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商品类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_type`
--

LOCK TABLES `product_type` WRITE;
/*!40000 ALTER TABLE `product_type` DISABLE KEYS */;
INSERT INTO `product_type` VALUES (1,0,'电脑'),(2,0,'家电'),(3,0,'手机'),(4,0,'服装'),(5,0,'食品'),(6,0,'办公'),(7,0,'家具'),(8,0,'图书'),(9,0,'珠宝'),(10,0,'厨具'),(11,0,'玩具'),(12,0,'消耗品');
/*!40000 ALTER TABLE `product_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_limit`
--

DROP TABLE IF EXISTS `sys_limit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_limit` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '权限Id',
  `type` longtext COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户权限';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_limit`
--

LOCK TABLES `sys_limit` WRITE;
/*!40000 ALTER TABLE `sys_limit` DISABLE KEYS */;
INSERT INTO `sys_limit` VALUES (1,'超级管理员'),(2,'商家'),(3,'普通用户');
/*!40000 ALTER TABLE `sys_limit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户Id',
  `user_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '密码',
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '邮箱',
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '电话号码',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',
  `limit_type` int(10) NOT NULL COMMENT '权限',
  `avatar` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '用户头像',
  `integral` int(10) NOT NULL DEFAULT '0' COMMENT '积分',
  `balance` int(10) NOT NULL DEFAULT '0' COMMENT '余额',
  `enable_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户是否激活',
  `like_num` int(10) NOT NULL DEFAULT '0' COMMENT '点赞',
  `dont_like` int(10) NOT NULL DEFAULT '0' COMMENT '踩',
  `uid` varchar(36) COLLATE utf8mb4_bin NOT NULL COMMENT '用户唯一标识',
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (3,'Undefined','$2a$10$hv./b8Zz2uni185JgVkVsOE/Ip6ZVQCuH7XDMBizeCLzfTW/gCcEm','15506531157@163.com','15506531157',0,1,'http://118.31.43.239:81/image/icon.png',0,0,1,0,0,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0'),(5,'Undefined6608','$2a$10$npN0C7L6c8sawiXbKDBWJedglOLV9k7KMwTniJY5q3Hmkxqk678Jq','19953008089@163.com','17865937390',1,2,'http://118.31.43.239:89/icon/8ef53985-f445-407b-a905-e652e0923a6b.png',0,0,1,0,0,'d2bcfba3-68ce-44ca-a16d-b0c1cba0e7fb'),(10,'unde0','$2a$10$lwaqTJE5wkO4D2PHC1lLouTzb.NqaBm9mipFgqusPAFDDzYjWI3OG','1105149059@qq.com','18888888888',0,3,'http://118.31.43.239:89/icon/0f8f3789-6a82-45d5-9bad-86985951ba72.png',0,0,1,0,0,'bdd83754-71ea-4ab5-85ac-104232e6c341');
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_address`
--

DROP TABLE IF EXISTS `user_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_address` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '发货地址ID',
  `user_uid` longtext COLLATE utf8mb4_bin NOT NULL,
  `consignee` longtext COLLATE utf8mb4_bin NOT NULL,
  `phone` longtext COLLATE utf8mb4_bin NOT NULL,
  `gender` tinyint(1) NOT NULL DEFAULT '0',
  `address` longtext COLLATE utf8mb4_bin NOT NULL,
  `def` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户收货地址';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_address`
--

LOCK TABLES `user_address` WRITE;
/*!40000 ALTER TABLE `user_address` DISABLE KEYS */;
INSERT INTO `user_address` VALUES (1,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0','张杰','15506531157',0,'历山路',0),(2,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0','张三','17865937390',1,'ccpark',0),(3,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0','李四','18888888888',0,'淄博职业学院',0),(4,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0','王五','19999999999',0,'淄博职业学院人工智能与大数据学院',0),(5,'7e5f1591-8d94-43ef-a080-c44cd0eca8c0','王五','17800908534',0,'淄博职业学院人工智能与大数据学院',0);
/*!40000 ALTER TABLE `user_address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_evaluate`
--

DROP TABLE IF EXISTS `user_evaluate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_evaluate` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `merchant_id` int(10) unsigned NOT NULL,
  `commentator_id` int(10) unsigned NOT NULL,
  `content` longtext COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商户评论';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_evaluate`
--

LOCK TABLES `user_evaluate` WRITE;
/*!40000 ALTER TABLE `user_evaluate` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_evaluate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'co_store'
--

--
-- Dumping routines for database 'co_store'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-28  9:36:58
