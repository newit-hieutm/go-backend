CREATE TABLE authors (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);

CREATE TABLE books (
  id        BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  author_id BIGINT  NOT NULL,
  title     text    NOT NULL,
  FOREIGN KEY (author_id) REFERENCES authors (id)
);