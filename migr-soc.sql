CREATE TABLE `Rounds` (
	`address` varchar(255) NOT NULL,
	`deposit` INT(255) NOT NULL,
	`bsnap` INT(255),
	`psnap` INT(255),
	`Spos` INT(255),
	`Sneg` INT(255),
	`reserve` INT(255)
);

CREATE TABLE `PendingPlayers` (
	`sender` varchar(255) NOT NULL,
	`roundAddress` varchar(255) NOT NULL
);

CREATE TABLE `Players` (
	`address` varchar(255) NOT NULL,
	`roundAddress` varchar(255) NOT NULL,
	`balance` INT(255),
	`nwin` INT(255),
	`n` INT(255),
	`spos` INT(255),
	`sneg` INT(255)
);

CREATE TABLE `Lots` (
	`address` varchar(255) NOT NULL,
	`roundAddress` varchar(255) NOT NULL,
	`owner` varchar(255),
	`timeFirst` INT(255),
	`timeSecond` INT(255),
	`value` INT(255),
	`price` INT(255),
	`receiveTokens` INT(255),
	`snapshot` INT(255)
);

ALTER TABLE `PendingPlayers` ADD CONSTRAINT `PendingPlayers_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Players` ADD CONSTRAINT `Players_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Lots` ADD CONSTRAINT `Lots_fk0` FOREIGN KEY (`roundAddress`) REFERENCES `Rounds`(`address`);

ALTER TABLE `Lots` ADD CONSTRAINT `Lots_fk1` FOREIGN KEY (`owner`) REFERENCES `Players`(`address`);





