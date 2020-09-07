
-- +migrate Up
CREATE TABLE IF NOT EXISTS `daily_results` (
  `user_id` varchar(255) NOT NULL DEFAULT '',
  `followers_count` int(11) DEFAULT NULL,
  `increase_followers_count` int(11) DEFAULT NULL,
  `point` int(11) DEFAULT NULL,
  `ranking` int(11) DEFAULT NULL,
  `last_ranking` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `daily_results`;
