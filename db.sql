CREATE TABLE `userinfo` (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) DEFAULT NULL,
  `password` VARCHAR(64) DEFAULT NULL,
  `createtime` DATE DEFAULT NULL,
  `address` VARCHAR(64) DEFAULT NULL,
  `age` INT(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
