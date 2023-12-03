CREATE TABLE IF NOT EXISTS "posts" (
			"ID"	INT,
			"Text"	TEXT,
			"User"	TEXT,
      "Filter" TEXT,
			PRIMARY KEY("ID")
		);

INSERT INTO posts
  (ID, Text, User,Filter)
VALUES
  ('1', 'test','en mode useeeer',''),
  ('2', 'test','use pour moi ✔','usage flow'),
  ('3', 'test','skibidibop',''),
  ('4', 'test','yes','');