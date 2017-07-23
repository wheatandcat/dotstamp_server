-- MySQL dump 10.16  Distrib 10.2.7-MariaDB, for osx10.12 (x86_64)
--
-- Host: localhost    Database: stamp_test
-- ------------------------------------------------------
-- Server version	10.2.7-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `user_contributions`
--

LOCK TABLES `user_contributions` WRITE;
/*!40000 ALTER TABLE `user_contributions` DISABLE KEYS */;
INSERT INTO `user_contributions` VALUES (1,1,'foo',1,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
INSERT INTO `user_contributions` VALUES (2,1,'bar',1,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
/*!40000 ALTER TABLE `user_contributions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_contribution_details`
--

LOCK TABLES `user_contribution_details` WRITE;
/*!40000 ALTER TABLE `user_contribution_details` DISABLE KEYS */;
INSERT INTO `user_contribution_details` VALUES (1,1,'[{"priority":1,"body":"foo","iconType":1,"iconFace":1,"directionType":1,"talkType":1,"edit":false,"character":{"id":1,"fileName":"1.jpg","voiceType":1}}]','2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
INSERT INTO `user_contribution_details` VALUES (2,1,'[{"priority":1,"body":"foo","iconType":1,"iconFace":1,"directionType":1,"talkType":1,"edit":false,"character":{"id":1,"fileName":"1.jpg","voiceType":1}}]','2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
/*!40000 ALTER TABLE `user_contribution_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_contribution_tags`
--

LOCK TABLES `user_contribution_tags` WRITE;
/*!40000 ALTER TABLE `user_contribution_tags` DISABLE KEYS */;
INSERT INTO `user_contribution_tags` VALUES (1,1,'foo,bar,foobar','2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
INSERT INTO `user_contribution_tags` VALUES (2,1,'foo,bar,foobar','2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
/*!40000 ALTER TABLE `user_contribution_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `log_user_contributions`
--

LOCK TABLES `log_user_contributions` WRITE;
/*!40000 ALTER TABLE `log_user_contributions` DISABLE KEYS */;
INSERT INTO `log_user_contributions` VALUES (1,1,1,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
INSERT INTO `log_user_contributions` VALUES (2,2,1,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);
/*!40000 ALTER TABLE `log_user_contributions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `user_contribution_searches`
--

INSERT INTO `user_contribution_searches` VALUES (1,1,'foo,bar',1,'2017-07-23 15:12:52','2017-07-23 15:12:52');
INSERT INTO `user_contribution_searches` VALUES (2,1,'bar',1,'2017-07-23 15:12:52','2017-07-23 15:12:52');
--
-- Dumping data for table `user_character_images`
--

INSERT INTO `user_character_images` VALUES (1,1,0,0,1,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);

--
-- Dumping data for table `user_contribution_follows`
--

INSERT INTO `user_contribution_follows` VALUES (1,1,2,'2017-07-23 15:12:52','2017-07-23 15:12:52',NULL);

--
-- Dumping data for table `user_forget_passwords`
--

INSERT INTO `user_forget_passwords` VALUES (1,"test@tedt.com","abcdef",'2015-01-01 10:00:00','2015-01-01 10:00:00',NULL);
INSERT INTO `user_forget_passwords` VALUES (2,"dotstamplocaltest@gmail.com","XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDaFpLSjFbcXoEFfRsWxP",'2015-01-01 10:00:00','2015-01-01 10:00:00',NULL);

--
-- Dumping data for table `user_contribution_sound_lengths`
--

INSERT INTO `user_contribution_sound_lengths` VALUES (1,1,15,30,'2015-01-01 10:00:00','2015-01-01 10:00:00',NULL);


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-07-23 20:07:22
