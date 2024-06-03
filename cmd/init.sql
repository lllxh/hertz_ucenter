CREATE TABLE `user` (
                        `username` longtext COLLATE utf8mb4_unicode_ci,
                        `userAccount` longtext COLLATE utf8mb4_unicode_ci,
                        `avatarUrl` longtext COLLATE utf8mb4_unicode_ci,
                        `gender` tinyint(4) DEFAULT NULL,
                        `userPassword` longtext COLLATE utf8mb4_unicode_ci,
                        `phone` longtext COLLATE utf8mb4_unicode_ci,
                        `email` longtext COLLATE utf8mb4_unicode_ci,
                        `userStatus` bigint(20) DEFAULT NULL,
                        `userRole` bigint(20) DEFAULT NULL,
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `created_at` datetime(3) DEFAULT NULL,
                        `updated_at` datetime(3) DEFAULT NULL,
                        `deleted_at` datetime(3) DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

