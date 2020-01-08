CREATE Table users(id int NOT NULL AUTO_INCREMENT, login varchar(50), PRIMARY KEY (id));
INSERT INTO users (id, login) VALUES (1, 'nulpatrol');
INSERT INTO users (id, login) VALUES (2, 'alonat');

CREATE TABLE languages (id int NOT NULL AUTO_INCREMENT, code varchar(50), PRIMARY KEY (id))
INSERT INTO languages (id, code) VALUES (1, 'en');
INSERT INTO languages (id, code) VALUES (2, 'de');

CREATE TABLE words (id int NOT NULL AUTO_INCREMENT, word varchar(50), PRIMARY KEY (id))
INSERT INTO words (id, word) VALUES (1, 'to be');

CREATE TABLE words_forms (
	id int NOT NULL AUTO_INCREMENT,
	word_id int NOT NULL,
	form varchar(50),
	PRIMARY KEY (id),
	FOREIGN KEY (word_id) REFERENCES words(id)
);
INSERT INTO words_forms (id, word_id, form) VALUES (1, 1, 'am');
INSERT INTO words_forms (id, word_id, form) VALUES (2, 1, 'are');
INSERT INTO words_forms (id, word_id, form) VALUES (3, 1, 'is');
