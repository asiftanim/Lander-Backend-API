-- MariaDB dump 10.18  Distrib 10.4.17-MariaDB, for Win64 (AMD64)
--
-- Host: 127.0.0.1    Database: lander
-- ------------------------------------------------------
-- Server version	10.5.10-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `brokers`
--

DROP TABLE IF EXISTS `brokers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `brokers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) CHARACTER SET latin1 NOT NULL,
  `email` varchar(256) CHARACTER SET latin1 NOT NULL,
  `password` varchar(256) CHARACTER SET latin1 NOT NULL,
  `image_path` varchar(256) CHARACTER SET latin1 DEFAULT NULL,
  `is_active` tinyint(4) NOT NULL,
  `is_delete` tinyint(4) NOT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  `modified_by` int(11) DEFAULT NULL,
  `modified_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chat_mail_candidates`
--

DROP TABLE IF EXISTS `chat_mail_candidates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chat_mail_candidates` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chat_id` int(11) NOT NULL,
  `mail_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
  `is_active` tinyint(4) DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_MAIL_CANDIDATE_CHAT_idx` (`chat_id`),
  CONSTRAINT `FK_MAIL_CANDIDATE_CHAT` FOREIGN KEY (`chat_id`) REFERENCES `chats` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chats`
--

DROP TABLE IF EXISTS `chats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` varchar(256) CHARACTER SET latin1 NOT NULL,
  `sender_id` int(11) NOT NULL,
  `receiver_id` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `user_type` int(11) NOT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `domains`
--

DROP TABLE IF EXISTS `domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `domains` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) NOT NULL,
  `name` varchar(256) COLLATE utf8mb4_bin DEFAULT NULL,
  `price_eur` decimal(9,2) DEFAULT NULL,
  `price_usd` decimal(9,2) DEFAULT NULL,
  `is_purshable` tinyint(4) DEFAULT NULL,
  `modified_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `prospect_domain_queries`
--

DROP TABLE IF EXISTS `prospect_domain_queries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `prospect_domain_queries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `prospect_id` int(11) NOT NULL,
  `broker_id` int(11) DEFAULT NULL,
  `domain_id` int(11) NOT NULL,
  `offer_type` int(11) DEFAULT NULL,
  `offered_price` decimal(9,2) DEFAULT NULL,
  `asking_price` decimal(9,2) DEFAULT NULL,
  `currency` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `status` int(11) NOT NULL,
  `is_price_seen` tinyint(4) DEFAULT NULL,
  `price_seen_time` timestamp(6) NULL DEFAULT NULL,
  `is_interacted` tinyint(4) DEFAULT NULL,
  `mail_sent_status` int(11) DEFAULT NULL,
  `mail_sent_time` timestamp(6) NULL DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  `modified_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_DOMAIN_QUERY_BROKER_idx` (`broker_id`),
  KEY `FK_DOMAIN_QUERY_PROSPECT_idx` (`prospect_id`),
  CONSTRAINT `FK_DOMAIN_QUERY_BROKER` FOREIGN KEY (`broker_id`) REFERENCES `brokers` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_DOMAIN_QUERY_PROSPECT` FOREIGN KEY (`prospect_id`) REFERENCES `prospects` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `prospects`
--

DROP TABLE IF EXISTS `prospects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `prospects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(256) CHARACTER SET latin1 DEFAULT NULL,
  `last_name` varchar(256) CHARACTER SET latin1 DEFAULT NULL,
  `email` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `cooki_secrect` varchar(256) CHARACTER SET latin1 DEFAULT NULL,
  `company` varchar(256) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_active` tinyint(4) NOT NULL,
  `is_delete` tinyint(4) NOT NULL,
  `is_email_verified` tinyint(4) DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  `modified_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `public_domain_prices`
--

DROP TABLE IF EXISTS `public_domain_prices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `public_domain_prices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) NOT NULL,
  `price` decimal(9,2) NOT NULL,
  `currency` varchar(45) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `prospect_id` int(11) NOT NULL,
  `prospect_domain_query_id` int(11) DEFAULT NULL,
  `price` decimal(9,2) NOT NULL,
  `payment_method` int(11) NOT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  `modified_by` int(11) DEFAULT NULL,
  `modified_at` timestamp(6) NULL DEFAULT NULL,
  `is_active` tinyint(4) NOT NULL,
  `is_delete` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_TRANSACTION_PROSPECT_idx` (`prospect_id`),
  KEY `FK_TRANSACTION_PROSPECT_DOMAIN_QUERY_idx` (`prospect_domain_query_id`),
  CONSTRAINT `FK_TRANSACTION_PROSPECT` FOREIGN KEY (`prospect_id`) REFERENCES `prospects` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_TRANSACTION_PROSPECT_DOMAIN_QUERY` FOREIGN KEY (`prospect_domain_query_id`) REFERENCES `prospect_domain_queries` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-23 19:23:23
