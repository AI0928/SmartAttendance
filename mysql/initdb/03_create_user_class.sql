DROP TABLE IF EXISTS user_class;

CREATE TABLE user_class (
  `user_id` VARCHAR(255) NOT NULL,
  `owner_id` VARCHAR(255) NOT NULL,
  `times` VARCHAR(255) NOT NULL,
  `flag` VARCHAR(255),
  PRIMARY KEY (`user_id`, `owner_id`, `times`)
);