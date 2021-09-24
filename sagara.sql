# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.35)
# Database: kumparan
# Generation Time: 2021-09-06 00:04:41 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table article
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS sagaracrud;
USE sagaracrud;
CREATE TABLE IF NOT EXISTS product (id int(11) unsigned NOT NULL AUTO_INCREMENT,name varchar(200) DEFAULT NULL,description text,price int(11) DEFAULT NULL,image varchar(150) DEFAULT NULL,created_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,status int(1) DEFAULT '1',PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE user (id int(11) unsigned NOT NULL AUTO_INCREMENT,username varchar(20) DEFAULT NULL,full_name varchar(60) DEFAULT NULL,password varchar(200) DEFAULT NULL,type_user int(1) DEFAULT NULL,PRIMARY KEY (id),UNIQUE KEY username (username)) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `user` (`id`, `username`, `full_name`, `password`, `type_user`)
VALUES
	(6, 'adminasik25', 'babang tamvan', 'dc0ad7ee526a332b093d16a4b84f7cf542e79b4aba30a4e20724f1b5c37e8141', 1);

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
