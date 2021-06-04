CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) ,
  `last_name` varchar(255) ,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `address1` varchar(255) ,
  `address2` varchar(255) ,
  `phone_number` int(10),
  `Decription` varchar(255),
  `token` varchar(255) DEFAULT NULL,
  `token_expried_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;