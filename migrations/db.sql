# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.13-MariaDB-1:10.4.13+maria~focal)
# Database: db_member_01
# Generation Time: 2023-02-03 05:53:03 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table tb_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_user`;

CREATE TABLE `tb_user` (
                           `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                           `upw` varchar(40) DEFAULT NULL COMMENT 'password',
                           `cell` varchar(20) DEFAULT NULL COMMENT '手机',
                           `sex` tinyint(3) DEFAULT '0' COMMENT '0-初始化  1-男  2-女',
                           `nickname` varchar(190) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                           `location` varchar(32) DEFAULT NULL COMMENT '位置',
                           `headimg` varchar(512) DEFAULT NULL COMMENT '头像',
                           `headimg_web` varchar(512) DEFAULT NULL COMMENT 'WEB站头像',
                           `birthday` date DEFAULT '1990-01-01',
                           `year_protected` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0：显示 1：隐藏',
                           `desc` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '个人简介',
                           `reg_date` datetime NOT NULL COMMENT '注册日期',
                           `ent_date` datetime DEFAULT NULL COMMENT '最后一次登录时间',
                           `ent_ip` varchar(32) DEFAULT NULL COMMENT '最后一次登录IP',
                           `dv` varchar(32) DEFAULT NULL COMMENT '登录方式',
                           `lvl` int(11) NOT NULL DEFAULT '1' COMMENT '等级',
                           `exp` int(11) DEFAULT '0' COMMENT '经验',
                           `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `state` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态(0:正常 1:销号)',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='USER TABLE';



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
