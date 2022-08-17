-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: localhost    Database: users-db
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
-- Dumping data for table `alembic_version`
--

LOCK TABLES `alembic_version` WRITE;
/*!40000 ALTER TABLE `alembic_version` DISABLE KEYS */;
INSERT INTO `alembic_version` VALUES ('e5f93d395664');
/*!40000 ALTER TABLE `alembic_version` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_verification`
--

LOCK TABLES `user_verification` WRITE;
/*!40000 ALTER TABLE `user_verification` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_verification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','123','admin@gmail.com','admin','admin','ROLE_ADMIN',1),(3,'markuza','123','markuza@gmail.com','Petar','Markovic','ROLE_EMPLOYEE',1),(4,'dan','123','regohim693@wnpop.com','Daniel','Bozanic','ROLE_USER',1),(5,'uros','123','pipoli3562@yubua.com','Uros','Tomic','ROLE_USER',1),(6,'mare','123','mare@gmail.com','Marko','Markovic','ROLE_USER',1),(7,'zika','123','zika@gmail.com','Zika','Zivkovic','ROLE_USER',1),(8,'rane','123','rane@gmail.com','Ranko','Tomic','ROLE_USER',1),(9,'dragan','123','dragan@gmail.com','Dragan','Zivkovic','ROLE_USER',1),(10,'una','123','una@hotmail.com','Una','Unic','ROLE_USER',1),(11,'tara','123','tara@yahoo.com','Tamara','Jovanovic','ROLE_USER',1),(12,'petruza','123','petruza@hotmail.com','Petar','Petrovic','ROLE_USER',1),(13,'alex','123','alex@gmail.com','Aleksandar','Rankovic','ROLE_USER',1),(14,'neca','123','neca@gmail.com','Nemanja','Pavlovic','ROLE_USER',1),(33,'admin1','123','admin1@hotmail.com','admin1','admin1','ROLE_ADMIN',1),(34,'wumpa','123','jawexi5480@yasiok.com','Wumpa','Wumpic','ROLE_USER',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-17 21:20:59
