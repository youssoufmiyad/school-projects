CREATE TABLE IF NOT EXISTS "users" (
			"ID"	INT,
            "email" TEXT UNIQUE,
			"username"	TEXT,
			"password"	TEXT,
            PRIMARY KEY (ID)
		);
INSERT INTO users
  (ID,email, username, password)
VALUES
  (1,'addtest1@piridotcom', 'en mode useeeer','1234'),
  (2,'addtest2@piridotcom', 'use pour moi ✔','1234')