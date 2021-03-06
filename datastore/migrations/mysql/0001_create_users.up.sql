CREATE TABLE IF NOT EXISTS Users (
	ID INTEGER AUTO_INCREMENT,
	CreationTime DATETIME,
	LastUpdatedTime DATETIME,
	DeletionTime DATETIME,
	Status INTEGER,
	Username VARCHAR(128) NOT NULL,
	Password CHAR(60) NOT NULL,
	Email VARCHAR(254),
	CONSTRAINT pk_UsersID PRIMARY KEY (ID),
	CONSTRAINT uk_UsersUsername UNIQUE (Username)
);

CREATE INDEX ix_UsersUsername ON Users(Username);