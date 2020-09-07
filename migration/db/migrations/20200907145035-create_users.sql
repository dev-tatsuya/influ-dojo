
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `screen_name` varchar(255) NOT NULL DEFAULT '',
  `profile_image` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;
