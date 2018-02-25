CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(32),
  `password` varchar(32),
  `name` varchar(100),
  CONSTRAINT `i_unique_username` UNIQUE(`username`)
)