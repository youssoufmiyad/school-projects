CREATE TABLE IF NOT EXISTS "comment" (
			"ID"	INT,
            "CommentID" INT,
			"Text"	TEXT,
			"User"	TEXT,
      PRIMARY KEY("CommentID")
		);
INSERT INTO comment
  (ID,CommentID, Text, User)
VALUES
  (1,1, 'testcom','user1'),
  (1,2, 'testcom','user2'),
  (2,3, 'testcom','user3'),
  (1,4, 'testcom','user4');