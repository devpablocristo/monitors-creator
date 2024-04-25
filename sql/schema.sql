-- monitors.business_hours definition

CREATE TABLE `business_hours` (
  `id` int NOT NULL AUTO_INCREMENT,
  `day` int DEFAULT NULL,
  `open` time DEFAULT NULL,
  `close` time DEFAULT NULL,
  `monitor_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.business_unit definition

CREATE TABLE `business_unit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_unit` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.collector definition

CREATE TABLE `collector` (
  `id` int NOT NULL AUTO_INCREMENT,
  `collector_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pdv_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.monitor definition

CREATE TABLE `monitor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `business_unit` int NOT NULL,
  `product` int NOT NULL,
  `brand_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `flow` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `site` int NOT NULL,
  `description` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tpv` float NOT NULL,
  `kam` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `seller_contact` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `technical_contact` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.pdv definition

CREATE TABLE `pdv` (
  `id` int NOT NULL AUTO_INCREMENT,
  `application_client_id` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `brand_id` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `monitor_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.platform definition

CREATE TABLE `platform` (
  `id` int NOT NULL AUTO_INCREMENT,
  `marketplace_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `platform_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sponsor_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `monitor_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.product definition

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.site definition

CREATE TABLE `site` (
  `id` int NOT NULL AUTO_INCREMENT,
  `site` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- monitors.`day` definition

CREATE TABLE `day` (
  `id` int NOT NULL AUTO_INCREMENT,
  `day` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;