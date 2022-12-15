CREATE TABLE `Rounds` (
	`address` VARCHAR(255) NOT NULL,
	`deposit` VARCHAR(255) NOT NULL,
	`bsnap` VARCHAR(255) DEFAULT '',
	`psnap` VARCHAR(255) DEFAULT '',
	`Spos` VARCHAR(255) DEFAULT '',
	`Sneg` VARCHAR(255) DEFAULT '',
	`reserve` VARCHAR(255) DEFAULT ''
);

CREATE TABLE `PendingPlayers` (
	`sender` VARCHAR(255) NOT NULL,
	`roundAddress` VARCHAR(255) NOT NULL
);

CREATE TABLE `Players` (
	`address` varchar(255) NOT NULL,
	`roundAddress` varchar(255) NOT NULL,
	`balance` INT(255) DEFAULT 0,
	`nwin` INT(255) DEFAULT 0,
	`n` INT(255) DEFAULT 0,
	`spos` VARCHAR(255) DEFAULT '',
	`sneg` VARCHAR(255) DEFAULT ''
);

CREATE TABLE `Lots` (
	`address` varchar(255) NOT NULL,
	`roundAddress` varchar(255) NOT NULL,
	`owner` varchar(255) DEFAULT '',
	`timeFirst` VARCHAR(255) DEFAULT '',
	`timeSecond` VARCHAR(255) DEFAULT '',
	`value` VARCHAR(255) DEFAULT '',
	`price` INT(255) DEFAULT 0,
	`receiveTokens` VARCHAR(255) DEFAULT '',
	`snapshot` VARCHAR(255) DEFAULT '',
	`prevSnapshot` VARCHAR(255) DEFAULT ''
);

ALTER TABLE `PendingPlayers` ADD CONSTRAINT `PendingPlayers_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Players` ADD CONSTRAINT `Players_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Lots` ADD CONSTRAINT `Lots_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Lots` ADD CONSTRAINT `Lots_fk1` FOREIGN KEY (`owner`) REFERENCES `Players`(`address`);





