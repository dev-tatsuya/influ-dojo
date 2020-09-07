
-- +migrate Up
CREATE TABLE IF NOT EXISTS `daily_works` (
  `user_id` varchar(255) NOT NULL DEFAULT '',
  `tweets_count` int(11) DEFAULT NULL,
  `increase_tweets_count` int(11) DEFAULT NULL,
  `my_tweets_count` int(11) DEFAULT NULL,
  `replies_count` int(11) DEFAULT NULL,
  `favorites_count` int(11) DEFAULT NULL,
  `increase_favorites_count` int(11) DEFAULT NULL,
  `point` double DEFAULT NULL,
  `ranking` int(11) DEFAULT NULL,
  `last_ranking` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `daily_works`;
