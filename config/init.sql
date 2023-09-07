-- MariaDB dump 10.19-11.0.2-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: minitok
-- ------------------------------------------------------
-- Server version	11.0.2-MariaDB

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

CREATE DATABASE IF NOT EXISTS minitok;
USE minitok;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `video_id` bigint(20) unsigned DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,1,'测试评论1-1'),
                              (2,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,1,'测试评论1-1'),
                              (3,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,3,1,'测试评论3-1'),
                              (4,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,4,1,'测试评论4-1'),
                              (5,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,5,1,'测试评论5-1'),
                              (6,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,2,'测试评论1-2'),
                              (7,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,3,'测试评论2-3'),
                              (8,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,3,'测试评论1-3'),
                              (9,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,4,'测试评论1-4'),
                              (10,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,5,'测试评论1-5'),
                              (11,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,6,'测试评论2-6'),
                              (12,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,7,'测试评论1-7'),
                              (13,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,6,8,'测试评论6-8'),
                              (14,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,7,8,'测试评论7-8'),
                              (15,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,5,9,'测试评论5-9'),
                              (16,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,4,9,'测试评论4-9'),
                              (17,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,9,9,'测试评论9-9'),
                              (18,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,10,'测试评论2-10'),
                              (19,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,11,'测试评论2-11'),
                              (20,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,3,11,'测试评论3-11'),
                              (21,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,3,12,'测试评论3-12'),
                              (22,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,6,12,'测试评论6-12'),
                              (23,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,6,12,'测试评论6-12'),
                              (24,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,7,13,'测试评论7-13'),
                              (25,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,8,13,'测试评论8-13'),
                              (26,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,13,'测试评论1-13'),
                              (27,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,13,'测试评论2-13'),
                              (28,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,5,14,'测试评论5-14'),
                              (29,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,3,15,'测试评论3-15'),
                              (30,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,2,16,'测试评论2-16'),
                              (31,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,3,16,'测试评论3-16');
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `favorites`
--

DROP TABLE IF EXISTS `favorites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `favorites` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `video_id` longtext DEFAULT NULL,
  `user_id` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_favorites_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `favorites`
--

LOCK TABLES `favorites` WRITE;
/*!40000 ALTER TABLE `favorites` DISABLE KEYS */;
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:20:56.193', '2023-09-05 04:20:56.193', NULL, '1', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:25.048', '2023-09-05 04:21:25.048', NULL, '2', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:28.536', '2023-09-05 04:21:28.536', NULL, '3', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:31.771', '2023-09-05 04:21:31.771', NULL, '4', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:35.869', '2023-09-05 04:21:35.869', NULL, '5', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:39.345', '2023-09-05 04:21:39.345', NULL, '6', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:42.450', '2023-09-05 04:21:42.450', NULL, '7', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:50.740', '2023-09-05 04:21:50.740', NULL, '8', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:53.895', '2023-09-05 04:21:53.895', NULL, '9', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:21:57.609', '2023-09-05 04:21:57.609', NULL, '10', '1');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:29:57.149', '2023-09-05 04:29:57.149', NULL, '1', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:04.912', '2023-09-05 04:30:04.912', NULL, '2', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:07.733', '2023-09-05 04:30:07.733', NULL, '3', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:10.436', '2023-09-05 04:30:10.436', NULL, '4', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:13.201', '2023-09-05 04:30:13.201', NULL, '5', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:22.284', '2023-09-05 04:30:22.284', NULL, '6', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:25.919', '2023-09-05 04:30:25.919', NULL, '7', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:28.314', '2023-09-05 04:30:28.314', NULL, '8', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:30.698', '2023-09-05 04:30:30.698', NULL, '9', '2');
INSERT INTO `favorites`(`created_at`, `updated_at`, `deleted_at`, `video_id`, `user_id`) VALUES ('2023-09-05 04:30:33.595', '2023-09-05 04:30:33.595', NULL, '10', '2');
/*!40000 ALTER TABLE `favorites` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `avatar` longtext DEFAULT NULL,
  `background_image` longtext DEFAULT NULL,
  `signature` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2023-09-05 21:00:56.193','2023-09-05 21:00:56.193',NULL,'xxhhy','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (2,'2023-09-05 21:01:56.193','2023-09-05 21:01:56.193',NULL,'Nefelibata1','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (3,'2023-09-05 21:02:56.193','2023-09-05 21:02:56.193',NULL,'Nefelibata2','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (4,'2023-09-05 21:03:56.193','2023-09-05 21:03:56.193',NULL,'Nefelibata4','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (5,'2023-09-05 21:04:56.193','2023-09-05 21:04:56.193',NULL,'Nefelibata5','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (6,'2023-09-05 21:05:56.193','2023-09-05 21:05:56.193',NULL,'Nefelibata6','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (7,'2023-09-05 21:06:56.193','2023-09-05 21:06:56.193',NULL,'Nefelibata7','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (8,'2023-09-05 21:07:56.193','2023-09-05 21:07:56.193',NULL,'Nefelibata8','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (9,'2023-09-05 21:08:56.193','2023-09-05 21:08:56.193',NULL,'Nefelibata9','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (10,'2023-09-05 21:09:56.193','2023-09-05 21:09:56.193',NULL,'Nefelibata10','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (11,'2023-09-05 21:10:56.193','2023-09-05 21:10:56.193',NULL,'Nefelibata11','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (12,'2023-09-05 21:11:56.193','2023-09-05 21:11:56.193',NULL,'Nefelibata12','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (13,'2023-09-05 21:12:56.193','2023-09-05 21:12:56.193',NULL,'Nefelibata13','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (14,'2023-09-05 21:13:56.193','2023-09-05 21:13:56.193',NULL,'Nefelibata14','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (15,'2023-09-05 21:14:56.193','2023-09-05 21:14:56.193',NULL,'Nefelibata15','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (16,'2023-09-05 21:15:56.193','2023-09-05 21:15:56.193',NULL,'Nefelibata16','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (17,'2023-09-05 21:16:56.193','2023-09-05 21:16:56.193',NULL,'Nefelibata17','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (18,'2023-09-05 21:17:56.193','2023-09-05 21:17:56.193',NULL,'Nefelibata18','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (19,'2023-09-05 21:18:56.193','2023-09-05 21:18:56.193',NULL,'Nefelibata19','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (20,'2023-09-05 21:19:56.193','2023-09-05 21:19:56.193',NULL,'Nefelibata20','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (21,'2023-09-05 21:20:56.193','2023-09-05 21:20:56.193',NULL,'Nefelibata21','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (22,'2023-09-05 21:21:56.193','2023-09-05 21:21:56.193',NULL,'Nefelibata22','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (23,'2023-09-05 21:22:56.193','2023-09-05 21:22:56.193',NULL,'Nefelibata23','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (24,'2023-09-05 21:23:56.193','2023-09-05 21:23:56.193',NULL,'Nefelibata24','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (25,'2023-09-05 21:24:56.193','2023-09-05 21:24:56.193',NULL,'Nefelibata25','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (26,'2023-09-05 21:25:56.193','2023-09-05 21:25:56.193',NULL,'Nefelibata26','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (27,'2023-09-05 21:26:56.193','2023-09-05 21:26:56.193',NULL,'Nefelibata27','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (28,'2023-09-05 21:27:56.193','2023-09-05 21:27:56.193',NULL,'Nefelibata28','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (29,'2023-09-05 21:28:56.193','2023-09-05 21:28:56.193',NULL,'Nefelibata29','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试'),
                            (30,'2023-09-05 21:29:56.193','2023-09-05 21:29:56.193',NULL,'Nefelibata30','313233343536220a331c624e79f344c2a907ea0e3b065186d9a5c11f03eac84c74c3c3767185','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg','https://minitok-video.oss-cn-shanghai.aliyuncs.com/default_background.png', '测试');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `videos`
--

DROP TABLE IF EXISTS `videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `videos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `author_id` bigint(20) unsigned DEFAULT NULL,
  `play_url` longtext DEFAULT NULL,
  `cover_url` longtext DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_videos_deleted_at` (`deleted_at`),
  KEY `idx_videos_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (1,'2023-08-22 22:04:36.669','2023-08-22 22:04:36.669',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/3b423475-a920-493f-b503-65f8df58c3d1.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/e042d8c7-15cc-4a79-8548-1cb9bf668a52.png','测试用'),
                            (2,'2023-08-22 22:04:37.421','2023-08-22 22:04:37.421',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/22db5c57-8d48-4e8d-9b3e-0ca2a28ba3bd.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/1da28046-08f6-4eda-b405-f8962d881460.png','测试用'),
                            (3,'2023-08-22 22:04:39.882','2023-08-22 22:04:39.882',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/a8646211-ad44-49cb-a86e-a95d0f0ad25c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/b095a504-3a4b-4c21-929f-4a340cdbfc73.png','测试用'),
                            (4,'2023-08-22 22:15:41.049','2023-08-22 22:15:41.049',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/df38d388-d795-4a9d-97c1-8cddde4046bd.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/a7059f11-3e04-42c5-abf5-a8063a0d3458.png','测试用'),
                            (5,'2023-08-22 22:15:41.154','2023-08-22 22:15:41.154',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/00739732-9d6e-43ec-a3f0-be6b552de4b8.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/d4ebdb35-dd01-4017-9b42-6d29b6d5b636.png','测试用'),
                            (6,'2023-08-22 22:17:37.713','2023-08-22 22:17:37.713',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/7dbe2bb2-6a89-4267-b9dc-b7e21f8f4b2c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/04db80a9-9225-4740-b523-9abf9bcb145b.png','测试用'),
                            (7,'2023-08-22 22:17:37.712','2023-08-22 22:17:37.712',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/31cd51ed-734d-4960-9640-fcdf7704c2fa.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/c3644c7b-13b0-4390-b0bd-89b2288e53ab.png','测试用'),
                            (8,'2023-08-22 22:17:39.439','2023-08-22 22:17:39.439',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/e74cafb3-ab0e-48e0-8552-d9b76bbca21c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/875db781-cc3e-476a-b3f8-1bd057c623ac.png','测试用'),
                            (9,'2023-08-22 22:18:37.670','2023-08-22 22:18:37.670',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/b694c95a-6854-4b02-bac0-dedf30d82d71.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/e63c78f2-83f8-4cd1-96d5-3883dcdef79a.png','测试用'),
                            (10,'2023-08-22 22:18:45.227','2023-08-22 22:18:45.227',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/f7b1e246-8557-427d-91f6-9db84d8ed197.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/f70d06fa-08d4-4a87-8a5a-2526aecf4cc2.png','测试用'),
                            (11,'2023-08-22 22:18:51.184','2023-08-22 22:18:51.184',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/cd0da9db-171a-456a-a21b-43a9a0ba7e8a.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/498b66e5-6476-454c-95df-6468737ae716.png','测试用'),
                            (12,'2023-08-22 22:18:54.902','2023-08-22 22:18:54.902',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/8eb41aff-c376-4346-8c0e-2da4a4647434.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/8468c9b2-6764-4a53-a856-25ca319d37c0.png','测试用'),
                            (13,'2023-08-22 22:18:56.581','2023-08-22 22:18:56.581',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/b2a7c33c-b51a-4ea2-86e4-9effc8d9d44c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/4d80009c-4a1f-4eaa-984f-489f855dd000.png','测试用'),
                            (14,'2023-08-22 22:22:02.580','2023-08-22 22:22:02.580',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/6257f6ee-5558-486f-b7d5-df44f9132374.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/6ed724fa-c431-4a8e-9092-9888b22a9f24.png','测试用'),
                            (15,'2023-08-22 22:22:05.502','2023-08-22 22:22:05.502',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/906f17b1-6e49-431f-a323-5b99e8699772.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/5966aa44-cada-4132-8630-c5a52d5dfb35.png','测试用'),
                            (16,'2023-08-22 22:22:08.017','2023-08-22 22:22:08.017',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/31cf1de5-839d-47b0-82a1-0df74b8b7548.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/62a2e792-4375-4692-add6-2c1a6e800b62.png','测试用'),
                            (17,'2023-08-22 22:24:20.452','2023-08-22 22:24:20.452',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/4b7a831c-66c7-463d-a4c9-fb909b738658.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/d84179a0-4fdb-41f6-8c2f-007221230966.png','测试用'),
                            (18,'2023-08-22 22:25:00.627','2023-08-22 22:25:00.627',NULL,1,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/acb5124c-3351-474b-aafb-28b069875a6c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/9ee43793-f63f-4a7b-b320-376c915baee6.png','测试用');

INSERT INTO `videos` VALUES (19,'2023-08-22 22:50:01.295','2023-08-22 22:50:01.295',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/c8e78318-73ce-4b85-bd34-50194c8255cb.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/8cc61b34-448c-4b9a-a9ab-7524bac9131b.png','测试用1'),
                            (20,'2023-08-22 22:50:03.115','2023-08-22 22:50:03.115',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/ff47b274-8d89-43bb-a01b-59f33d457721.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/84f4276e-a959-43cd-83bb-5e9490ef74bc.png','测试用1'),
                            (21,'2023-08-22 22:50:04.511','2023-08-22 22:50:04.511',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/9f314d1c-5e26-4eb9-aa9f-ab86c905e135.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/3dc2459c-2869-46ab-82cc-283bb296fd04.png','测试用1'),
                            (22,'2023-08-22 22:50:06.218','2023-08-22 22:50:06.218',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/15b39c87-51a0-44b5-ba0b-bf27b11551db.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/bd7fa1d9-5fba-4574-a416-52dda7998616.png','测试用1'),
                            (23,'2023-08-22 22:50:07.745','2023-08-22 22:50:07.745',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/656c71e3-31f2-4b81-93ff-c9a8d0b2c26e.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/06b43c52-c76d-4bdd-8a4e-dc945aca3cfa.png','测试用1'),
                            (24,'2023-08-22 22:50:09.330','2023-08-22 22:50:09.330',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/827049c3-dbc7-41a9-a407-3d646f75366c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/995b737a-2904-441b-bff1-f0ad946f4a68.png','测试用1'),
                            (25,'2023-08-22 22:50:10.836','2023-08-22 22:50:10.836',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/ca48d293-d0c0-40b5-82cb-77eba2a8de23.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/78fc3e05-e31f-4d78-81a5-2dd97cef9712.png','测试用1'),
                            (26,'2023-08-23 00:09:03.071','2023-08-23 00:09:03.071',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/0d999211-6609-4fdc-ab89-7bf4e39868ca.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/00185808-8988-4710-98af-9b605bbc6cc6.png','测试用1'),
                            (27,'2023-08-23 00:09:03.174','2023-08-23 00:09:03.174',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/976fd4ac-89d8-41f7-9caa-7c60dc5daecd.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/ad2e234f-6d02-4e19-abe0-b72167405bf0.png','测试用1'),
                            (28,'2023-08-23 00:15:20.344','2023-08-23 00:15:20.344',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/5ca37ebd-120b-4293-be70-b016648147bb.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/ff78aac9-b711-446f-b39b-10d4f3a2a275.png','测试用1'),
                            (29,'2023-08-23 00:19:59.312','2023-08-23 00:19:59.312',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/59ae7d9c-875c-42c6-9935-e51d098da757.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/54a66779-9353-45f4-9942-6ad9bc6c9880.png','测试用1'),
                            (30,'2023-08-23 00:20:39.650','2023-08-23 00:20:39.650',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/a1874382-e137-4dfd-804a-5076195c77da.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/ef92f3be-7843-4456-a2c4-7da33494deaa.png','测试用1'),
                            (31,'2023-08-23 01:14:23.986','2023-08-23 01:14:23.986',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/70ee35a5-6fea-40a2-8732-c7cf3a86cd62.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/e95bb15f-4553-4bbb-afb3-f389dd609713.png','测试用1'),
                            (32,'2023-08-23 01:14:24.108','2023-08-23 01:14:24.108',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/c01d0888-f87a-431a-8125-386a7d720f9c.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/12e21dfb-d186-4ba1-9438-e453e4e14146.png','测试用1'),
                            (33,'2023-08-23 01:14:24.009','2023-08-23 01:14:24.009',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/5bb16da8-6a03-48fc-bc85-acb87f2d2b82.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/7526d2a8-eb5c-4035-abe1-312834b5cd4e.png','测试用1'),
                            (34,'2023-08-23 03:19:13.286','2023-08-23 03:19:13.286',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/a6273d53-e18e-452d-82d9-ae8aa0870227.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/b58a6562-4caf-4c59-a5b1-e7c56382e88d.png','测试用1'),
                            (35,'2023-08-23 03:20:27.770','2023-08-23 03:20:27.770',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/0b6db380-18d5-4f5c-8a0f-d66bb6d097b5.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/118e45ab-0b49-4c63-b013-15156ef01994.png','测试用2'),
                            (36,'2023-08-23 03:23:26.539','2023-08-23 03:23:26.539',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/989c47ab-9dd2-4a31-a18f-616b6c252ce9.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/d76381fb-8611-4ec9-9e32-362ebd4463d3.png','测试用2'),
                            (37,'2023-08-23 03:23:27.602','2023-08-23 03:23:27.602',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/0153d60d-e3ab-4a6f-892c-98bc0a53843f.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/7b05cc70-1ed2-4860-be77-6b7ef21c7fc1.png','测试用2'),
                            (38,'2023-08-23 03:23:48.161','2023-08-23 03:23:48.161',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/f9a32b61-330a-4b00-976c-6ec0d2759f38.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/b2038c9b-767e-4053-8ff2-f78f960b4dd6.png','测试用2'),
                            (39,'2023-08-23 03:43:10.040','2023-08-23 03:43:10.040',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/404d1d7d-2ef7-4cd3-bb9b-5c8262924283.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/4b737542-d44c-4e9d-ad88-ee6fea93a8c2.png','测试用2'),
                            (40,'2023-08-24 04:00:13.513','2023-08-24 04:00:13.513',NULL,2,'https://minitok-video.oss-cn-shanghai.aliyuncs.com/7ceedfdc-8477-4ad6-90f9-d905a0ee1ea9.mp4','https://minitok-video.oss-cn-shanghai.aliyuncs.com/76340b76-bc1e-4409-bd84-6615503e49d0.png','测试用3');
/*!40000 ALTER TABLE `videos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-12  2:41:47
