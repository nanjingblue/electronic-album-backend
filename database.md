# DATABASE
DBMS: MySql
## SHOW CREATE TABLE
由于使用的 GORM 框架, 表与表之间的关联被抽象为了`逻辑外键`。
所以在实际的数据库中并不存在`物理外键`。   

下面是手动添加外键后的结果，实际画图，以下方为准。
## user (用户表)
```sql
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) NOT NULL,
  `password_digest` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT 'active',
  `sex` varchar(255) DEFAULT NULL,
  `age` int unsigned DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_user_deleted_at` (`deleted_at`)
);
```
### album (相册表)
```sql
CREATE TABLE `album` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `album_name` varchar(255) NOT NULL,
  `user_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  UNIQUE KEY `album_name` (`album_name`),
  KEY `idx_album_deleted_at` (`deleted_at`)
);
```
### picture (图片表)
```sql
CREATE TABLE `picture` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `picture_name` varchar(255) NOT NULL,
  `link` varchar(255) NOT NULL,
  `album_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`album_id`) REFERENCES album(`id`),
  UNIQUE KEY `picture_name` (`picture_name`),
  KEY `idx_picture_deleted_at` (`deleted_at`)
);
```
### friend (朋友表)
```sql
CREATE TABLE `friend` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `friend_id` int unsigned NOT NULL,
  `relationship` varchar(255) NOT NULL DEFAULT 'unfollow',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  FOREIGN KEY (`friend_id`) REFERENCES user(`id`),
  KEY `idx_friend_deleted_at` (`deleted_at`)
);
```
### post (贴子表)
```sql
CREATE TABLE `post` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `image_one` varchar(255) DEFAULT NULL,
  `image_two` varchar(255) DEFAULT NULL,
  `image_three` varchar(255) DEFAULT NULL,
  `image_four` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  KEY `idx_post_deleted_at` (`deleted_at`)
)
```
### comment (评论表)
```sql
CREATE TABLE `comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `post_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`post_id`) REFERENCES post(`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  KEY `idx_comment_deleted_at` (`deleted_at`)
)
```