CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `Passport` varchar(30) DEFAULT NULL,
  `Password` varchar(45) DEFAULT NULL,
  `Nickname` varchar(30) DEFAULT NULL,
  `CreateAt` time DEFAULT NULL,
  `UpdateAt` time DEFAULT NULL,
  PRIMARY KEY (`id`)
)