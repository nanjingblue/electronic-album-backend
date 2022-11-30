# DATABASE
DBMS: MySql
## SHOW CREATE TABLE
由于使用的 GORM 框架, 表与表之间的关联被抽象为了`逻辑外键`。
所以在实际的数据库中并不存在`物理外键`。   

下面是手动添加外键后的结果，实际画图，以下方为准。
## users (用户表)
```sql
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) NOT NULL,
  `nickname` varchar(191) DEFAULT '???',
  `password_digest` longtext NOT NULL,
  `status` varchar(191) NOT NULL DEFAULT 'active',
  `sex` longtext,
  `age` bigint unsigned DEFAULT NULL,
  `avatar` longtext,
  `description` longtext,
  `gender` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_users_deleted_at` (`deleted_at`)
);
```
### galleries (相册表)
```sql
CREATE TABLE `galleries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `gallery_name` longtext NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `cover` longtext,
  `status` varchar(191) NOT NULL DEFAULT 'active',
  `description` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_galleries_user` (`user_id`),
  KEY `idx_galleries_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_galleries_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```
### pictures (图片表)
```sql
 CREATE TABLE `pictures` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `picture_name` longtext NOT NULL,
  `path` longtext NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `gallery_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pictures_user` (`user_id`),
  KEY `fk_pictures_gallery` (`gallery_id`),
  KEY `idx_pictures_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_pictures_gallery` FOREIGN KEY (`gallery_id`) REFERENCES `galleries` (`id`),
  CONSTRAINT `fk_pictures_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```
### friends (朋友表)
```sql
CREATE TABLE `friends` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `friend_id` bigint unsigned NOT NULL,
  `relationship` varchar(191) NOT NULL DEFAULT 'unfollow',
  PRIMARY KEY (`id`),
  KEY `fk_friends_user` (`user_id`),
  KEY `idx_friends_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_friends_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```
### posts (贴子表)
```sql
 CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `content` longtext,
  `image` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_posts_user` (`user_id`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```
### comments (评论表)
```sql
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `post_id` bigint unsigned NOT NULL,
  `content` longtext,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_comments_post` (`post_id`),
  KEY `fk_comments_user` (`user_id`),
  KEY `idx_comments_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```
### user_posts (用户和贴子关系表)
```sql
CREATE TABLE `user_posts` (
  `user_id` bigint unsigned NOT NULL,
  `post_id` bigint unsigned NOT NULL,
  `liked` tinyint(1) DEFAULT '0',
  `collected` tinyint(1) DEFAULT '0',
  `commented` tinyint(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`,`post_id`),
  KEY `idx_user_posts_deleted_at` (`deleted_at`),
  KEY `fk_user_posts_post` (`post_id`),
  CONSTRAINT `fk_user_posts_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  CONSTRAINT `fk_user_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
```