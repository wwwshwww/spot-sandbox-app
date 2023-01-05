CREATE DATABASE IF NOT EXISTS `yeah`;
USE `yeah`;

CREATE TABLE IF NOT EXISTS `spot`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `postal_code` VARCHAR(255) NOT NULL,
    `address_representation` VARCHAR(255) NOT NULL,
    `lat` FLOAT NOT NULL,
    `lng` FLOAT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`lat`, `lng`)
);

CREATE TABLE IF NOT EXISTS `dbscan_profile`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `distance_type` VARCHAR(255) NOT NULL,
    `min_count` INT NOT NULL,
    `max_count` INT,
    `meter_threshold` INT,
    `duration_threshold` BIGINT,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `spots_profile`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `spots_profile_spot`(
    `spots_profile_id` INT NOT NULL,
    `spot_id` INT NOT NULL,
    PRIMARY KEY (`spots_profile_id`, `spot_id`)
);