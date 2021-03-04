/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 8.0.23 : Database - test
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `test`;

/*Table structure for table `post` */

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userid` int unsigned NOT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `UpdateAt` timestamp NULL DEFAULT NULL,
  `title` varchar(30) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  `status` tinyint DEFAULT NULL COMMENT '正常1,精华9',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `post` */

insert  into `post`(`id`,`userid`,`CreateAt`,`UpdateAt`,`title`,`content`,`deleted`,`status`) values 
(24,3,'2021-02-19 10:38:48','2021-02-19 10:38:48','sb123123','说什么啊说什么啊说什么啊',0,NULL),
(25,4,'2021-02-19 10:42:17','2021-02-19 10:52:30','sb1231231','说什么啊说什么啊说什么啊',1,NULL),
(26,5,'2021-02-19 10:58:49','2021-02-19 10:58:49','2021年2月19日10:58:42','2021年2月19日10:58:42',0,NULL),
(27,6,'2021-02-19 15:23:44','2021-02-19 15:23:44','111111','气得我风热特特热乎软饭如果天天也让他要给他人还有益禾堂就会同意还挺近还有他今天黄景瑜突进',0,NULL),
(28,6,'2021-02-19 15:29:05','2021-02-19 15:29:05','       ','气得我风热特特热乎软饭如果天天也让他要给他人还有益禾堂就会同意还挺近还有他今天黄景瑜突进',0,NULL),
(29,6,'2021-02-19 15:33:51','2021-02-19 15:33:51','我的压岁钱_800字','   \"同学们，过年你们最开心的是什么时候呢？没错，当然就是收到压岁钱的时候啦！关于压岁钱，还有一个传说呢，你们知道吗？\n   相传在很久以前，有个叫做“祟”的妖怪，经常在大年三十的夜晚悄悄溜到人们家里，摸正在熟睡的小孩子的头，被它摸过的小孩儿，会生出各种奇怪的病。人们无意间发现“祟”很怕光，因此，大年三十的晚上，人们把屋里点上蜡烛，大人们都要守着孩子到半夜，希望能把“祟”吓跑。',0,NULL),
(30,6,'2021-02-19 15:36:20','2021-02-19 15:36:20','我的压岁钱_800字11111111111111','   \"同学们，过年你们最开心的是什么时候呢？没错，当然就是收到压岁钱的时候啦！关于压岁钱，还有一个传说呢，你们知道吗？\n   相传在很久以前，有个叫做“祟”的妖怪，经常在大年三十的夜晚悄悄溜到人们家里，摸正在熟睡的小孩子的头，被它摸过的小孩儿，会生出各种奇怪的病。人们无意间发现“祟”很怕光，因此，大年三十的晚上，人们把屋里点上蜡烛，大人们都要守着孩子到半夜，希望能把“祟”吓跑。',0,NULL),
(31,5,'2021-02-19 15:45:40','2021-02-19 15:45:40','          ','2021年2月19日10:58:42',0,NULL),
(32,5,'2021-02-19 15:55:19','2021-02-19 15:55:19',' 1111111111','2021年2月19日10:58:42',0,NULL),
(33,5,'2021-02-19 16:01:43','2021-02-19 16:01:43','1 1 1 1 1 1 1 ','2021年2月19日10:58:42',0,NULL),
(34,5,'2021-02-19 16:01:54','2021-02-19 16:01:54','   1                       ','2021年2月19日10:58:42',0,NULL),
(35,5,'2021-02-19 16:06:11','2021-02-19 16:06:11',' 11111111111111','2021年2月19日10:58:42',0,NULL),
(36,5,'2021-02-19 16:07:55','2021-02-19 16:07:55','1 1 1 1 1 1 1 1 ','2021年2月19日10:58:421',0,NULL),
(37,6,'2021-02-21 03:07:36','2021-02-21 03:07:36','我的压岁钱_800字11111111111111111','   \"同学们，过年你们最开心的是什么时候呢？没错，当然就是收到压岁钱的时候啦！关于压岁钱，还有一个传说呢，你们知道吗？\n   相传在很久以前，有个叫做“祟”的妖怪，经常在大年三十的夜晚悄悄溜到人们家里，摸正在熟睡的小孩子的头，被它摸过的小孩儿，会生出各种奇怪的病。人们无意间发现“祟”很怕光，因此，大年三十的晚上，人们把屋里点上蜡烛，大人们都要守着孩子到半夜，希望能把“祟”吓跑。',0,NULL),
(38,0,'2021-02-23 10:35:38','2021-02-23 10:35:38','2021年2月23日10:35:35','2021年2月23日10:35:35',0,NULL),
(39,0,'2021-02-23 10:38:07','2021-02-23 10:38:07','2021年2月23日10:38:05','2021年2月23日10:35:35',0,NULL);

/*Table structure for table `postvote` */

DROP TABLE IF EXISTS `postvote`;

CREATE TABLE `postvote` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL,
  `userid` int unsigned NOT NULL,
  `vote` tinyint DEFAULT NULL COMMENT '点赞1,踩0',
  `CreateAt` timestamp NULL DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `postvote` */

/*Table structure for table `reply` */

DROP TABLE IF EXISTS `reply`;

CREATE TABLE `reply` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL,
  `userid` int unsigned NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `reply` */

insert  into `reply`(`id`,`pid`,`userid`,`content`,`CreateAt`,`deleted`) values 
(1,24,3,'乱说什么!','2021-02-20 10:02:12',0),
(2,24,5,'啛啛喳喳错错错错错错错错错错错错错错错错错错','2021-02-20 11:04:39',0),
(3,24,5,'啛啛喳喳错错错错错错错错错错错错错错错错错错','2021-02-20 11:04:49',0),
(6,24,5,'啧啧啧啧啧啧做做做做做做做做做做做做','2021-02-20 11:14:30',1),
(7,26,5,'啧啧啧啧啧啧做做做做做做做做做做做做1','2021-02-20 11:26:53',0),
(8,26,5,'咔咔咔咔咔咔扩扩扩扩扩扩扩扩','2021-02-20 11:27:16',0),
(9,26,0,'咔咔咔咔咔咔扩扩扩扩扩扩扩扩','2021-02-20 13:59:16',0),
(10,24,3,'123123123123','2021-02-22 08:49:01',0),
(11,24,3,'12312312312312','2021-02-22 08:52:29',0),
(12,24,0,'121231231233123','2021-02-22 08:53:24',0),
(13,24,0,'1212312312331asd23','2021-02-22 08:53:35',0);

/*Table structure for table `replyvote` */

DROP TABLE IF EXISTS `replyvote`;

CREATE TABLE `replyvote` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `replyid` int unsigned DEFAULT NULL,
  `userid` int unsigned DEFAULT NULL,
  `vote` tinyint DEFAULT NULL COMMENT '点赞1,踩0',
  `CreateAt` timestamp NULL DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `replyvote` */

insert  into `replyvote`(`id`,`replyid`,`userid`,`vote`,`CreateAt`,`deleted`) values 
(6,9,3,1,'2021-02-20 16:58:47',0),
(7,1,3,1,'2021-02-20 16:59:03',0);

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `Passport` varchar(30) DEFAULT NULL,
  `Password` varchar(45) DEFAULT NULL,
  `Nickname` varchar(30) DEFAULT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `UpdateAt` timestamp NULL DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '正常1,锁定2,拉黑-1,未激活0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user` */

insert  into `user`(`id`,`Passport`,`Password`,`Nickname`,`CreateAt`,`UpdateAt`,`status`) values 
(3,'lqclqc','4297f44b13955235245b2497399d7a93','lqclqc','2021-02-19 10:27:12','2021-02-19 10:27:12',NULL),
(4,'lqclqclqc','4297f44b13955235245b2497399d7a93','lqclqclqc','2021-02-19 10:41:32','2021-02-19 10:41:32',NULL),
(5,'asdasd','4297f44b13955235245b2497399d7a93','asdasd','2021-02-19 10:57:32','2021-02-19 10:57:32',NULL),
(6,'quest1','66f23f70d8aa8680017410c7295f0f35','er','2021-02-19 14:58:38','2021-02-19 14:58:38',NULL),
(7,'quest12','66f23f70d8aa8680017410c7295f0f35','   ','2021-02-19 15:05:48','2021-02-19 15:05:48',NULL),
(8,'ppdppd','4297f44b13955235245b2497399d7a93','ppdppd','2021-02-23 10:32:27','2021-02-23 10:32:27',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
