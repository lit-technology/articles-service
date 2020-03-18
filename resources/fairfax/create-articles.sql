CREATE TABLE IF NOT EXISTS article (
	id INT PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	article_date DATE NOT NULL,
	body VARCHAR(8191) NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS article_tag (
	tag VARCHAR(50),
	article_date DATE,
	article_id INT,
	PRIMARY KEY (tag, article_date, article_id),
	FOREIGN KEY (article_id)
		REFERENCES article(id)
		ON DELETE CASCADE
);
