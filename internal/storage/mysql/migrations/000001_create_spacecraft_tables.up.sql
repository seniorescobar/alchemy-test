CREATE TABLE `spacecrafts` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` TEXT NOT NULL,
	`class` TEXT NOT NULL,
	`crew` INT NOT NULL,
	`image` TEXT NOT NULL,
	`value` FLOAT NOT NULL,
	`status` TEXT NOT NULL,
	`armaments` JSON NOT NULL,

	PRIMARY KEY (`id`) 
);
