CREATE TABLE IF NOT EXISTS user(
		userId  INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password text, 	
		email text
		-- token TEXT ,
		-- expiresAt DATETIME
);
CREATE TABLE IF NOT EXISTS user_sessions(
  token TEXT PRIMARY KEY,
  expiresAt TEXT,
  userId INTEGER,
  FOREIGN KEY (userId) REFERENCES user(userId)
);
CREATE TABLE IF NOT EXISTS posts(
		postId INTEGER PRIMARY KEY AUTOINCREMENT,
		author text,
		title text,
		content text,
		creationDate TEXT
		-- create foreign key by userID
);
CREATE TABLE IF NOT EXISTS posts_category(
    	postCategoryId INTEGER,
		category TEXT,
		FOREIGN KEY (postCategoryId) REFERENCES posts(postId) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comments(
    	commentsId INTEGER PRIMARY KEY AUTOINCREMENT,
    	postId INTEGER,
    	author TEXT,
    	content TEXT,
    	likes INT DEFAULT 0,
    	dislikes INT DEFAULT 0,
    	FOREIGN KEY (postId)  REFERENCES posts(postId)
);
CREATE TABLE IF NOT EXISTS likes(
    	likeId INTEGER PRIMARY KEY AUTOINCREMENT,
    	username TEXT,
    	postId INTEGER DEFAULT NULL,
    	commentsId INTEGER DEFAULT NULL,
		likeis BOOLEAN,
    	FOREIGN KEY (postId) REFERENCES posts(postId) ON DELETE CASCADE,
    	FOREIGN KEY (commentsId) REFERENCES comments(commentsId) ON DELETE CASCADE
);