
DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userid` int unsigned NOT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `UpdateAt` timestamp NULL DEFAULT NULL,
  `title` varchar(30) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `reply` */

DROP TABLE IF EXISTS `reply`;

CREATE TABLE `reply` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL,
  `userid` int unsigned NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除为1',
  `num` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `Passport` varchar(30) DEFAULT NULL,
  `Password` varchar(45) DEFAULT NULL,
  `Nickname` varchar(30) DEFAULT NULL,
  `CreateAt` timestamp NULL DEFAULT NULL,
  `UpdateAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
