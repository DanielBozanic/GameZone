-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: localhost    Database: contact-and-report-db
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `bans`
--

LOCK TABLES `bans` WRITE;
/*!40000 ALTER TABLE `bans` DISABLE KEYS */;
INSERT INTO `bans` VALUES (1,5,'Offensive language',NULL,'2022-09-16 18:54:00.000');
/*!40000 ALTER TABLE `bans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `contact_messages`
--

LOCK TABLES `contact_messages` WRITE;
/*!40000 ALTER TABLE `contact_messages` DISABLE KEYS */;
INSERT INTO `contact_messages` VALUES (1,4,'dan','Nintendo switch lite availability','Hi,\n\nI would like to purchase the nintendo switch lite before halloween for my son, so I was wondering when will it be available for purchase.\n\nSincerely,\nDaniel',NULL,'2022-08-16 20:31:31.025'),(2,4,'dan','Working hours during holidays','Hi,\nWill GameZone be open on the 5th of January?\n\nSincerely,\nDaniel','Hi Daniel,\n\nGameZone will be working on the 5th of January from 8:00 - 15:00.\n\nHave a good day.\nJack','2022-01-02 10:15:00.025'),(3,4,'dan','Elden Ring PC Deluxe Edition','Hi,\n\nI\'ve seen that you have only the standard version of Elden Ring for purchase, so I\'m interested to know when or if at all you will have Elden Ring PC Deluxe Edition available for purchase.\n\nDaniel','Hi Daniel,\n\nYou can purchase Elden Ring PC Deluxe Edition starting 30th of February.\n','2022-02-26 14:32:53.635'),(5,4,'dan','Halloween product discounts','Hi, \n\nAre there gonna be any Halloween product discounts?\n\nDaniel','Hi Daniel, \n\nAll products from 31st of October to 5th of November will be 30% off.','2021-10-27 12:10:53.635');
/*!40000 ALTER TABLE `contact_messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `reports`
--

LOCK TABLES `reports` WRITE;
/*!40000 ALTER TABLE `reports` DISABLE KEYS */;
INSERT INTO `reports` VALUES (1,5,'Inappropriate language','Criticism could have been formulated more professionally.','2022-08-15 18:31:35.437'),(2,5,'Offensive language','This person is offensive on every single post','2022-08-15 20:41:31.686');
/*!40000 ALTER TABLE `reports` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-17 14:53:23
