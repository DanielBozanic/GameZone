CREATE DATABASE IF NOT EXISTS `product-db`;
CREATE DATABASE IF NOT EXISTS `comment-and-rating-db`;
CREATE DATABASE IF NOT EXISTS `contact-and-report-db`;
CREATE DATABASE IF NOT EXISTS `news-db`;
CREATE DATABASE IF NOT EXISTS `users-db`;

CREATE USER IF NOT EXISTS 'testuser'@'localhost' IDENTIFIED BY 'PassWord_FOR_NTP!';
GRANT ALL on *.*  to 'testuser'@'%';