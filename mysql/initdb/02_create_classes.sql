DROP TABLE IF EXISTS classes;

CREATE TABLE classes (
  `id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255),
  `owner_id` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);