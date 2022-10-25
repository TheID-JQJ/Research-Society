CREATE TABLE `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sub` varchar(100) NOT NULL,
  `obj` varchar(100) NOT NULL,
  `act` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_finance_kinds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `feedbacks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `content` longtext NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_feedbacks_deleted_at` (`deleted_at`),
  KEY `fk_feedbacks_user_information` (`user_id`),
  CONSTRAINT `fk_feedbacks_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `finance_kinds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_finance_kinds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `finances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `kind_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `time` datetime(3) NOT NULL,
  `money` bigint NOT NULL,
  `details` longtext NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_finances_deleted_at` (`deleted_at`),
  KEY `fk_finances_finance_kind` (`kind_id`),
  KEY `fk_finances_user_information` (`user_id`),
  CONSTRAINT `fk_finances_finance_kind` FOREIGN KEY (`kind_id`) REFERENCES `finance_kinds` (`id`),
  CONSTRAINT `fk_finances_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `goods_informations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `picture` varchar(255) NOT NULL,
  `name` varchar(32) NOT NULL,
  `price` double NOT NULL,
  `integral_price` bigint NOT NULL,
  `monthly_sales` bigint NOT NULL DEFAULT '0',
  `inventory` bigint NOT NULL DEFAULT '0',
  `kind_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_goods_informations_deleted_at` (`deleted_at`),
  KEY `fk_goods_informations_goods_kind` (`kind_id`),
  CONSTRAINT `fk_goods_informations_goods_kind` FOREIGN KEY (`kind_id`) REFERENCES `goods_kinds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `goods_kinds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `description` longtext NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_goods_kinds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_number` varchar(10) NOT NULL,
  `avatar` longtext NOT NULL,
  `group_name` varchar(32) NOT NULL,
  `owner_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `group_number` (`group_number`),
  KEY `idx_groups_deleted_at` (`deleted_at`),
  KEY `fk_groups_user_information` (`owner_id`),
  CONSTRAINT `fk_groups_user_information` FOREIGN KEY (`owner_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `notices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `content` longtext NOT NULL,
  `time` datetime(3) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_notices_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `public_resource_kinds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_public_resource_kinds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `public_resources` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `picture` longtext NOT NULL,
  `name` varchar(32) NOT NULL,
  `price` double NOT NULL,
  `integral_price` bigint NOT NULL,
  `rent_amount` bigint NOT NULL DEFAULT '0',
  `kind_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_public_resources_deleted_at` (`deleted_at`),
  KEY `fk_public_resources_public_resource_kind` (`kind_id`),
  CONSTRAINT `fk_public_resources_public_resource_kind` FOREIGN KEY (`kind_id`) REFERENCES `public_resource_kinds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `seat_kinds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_seat_kinds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `seats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `picture` longtext NOT NULL,
  `kind_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_seats_deleted_at` (`deleted_at`),
  KEY `fk_seats_seat_kind` (`kind_id`),
  CONSTRAINT `fk_seats_seat_kind` FOREIGN KEY (`kind_id`) REFERENCES `seat_kinds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `target_types` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_target_types_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `residual` double NOT NULL,
  `integral` bigint NOT NULL,
  `duration` bigint NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `idx_user_accounts_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_user_accounts_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_appoints` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `seat_id` bigint unsigned NOT NULL,
  `start_time` datetime(3) NOT NULL,
  `end_t_ime` datetime(3) NOT NULL,
  `status` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_appoints_deleted_at` (`deleted_at`),
  KEY `fk_user_appoints_user_information` (`user_id`),
  KEY `fk_user_appoints_seat` (`seat_id`),
  CONSTRAINT `fk_user_appoints_seat` FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`),
  CONSTRAINT `fk_user_appoints_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_donates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `public_resource_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `public_resource_id` (`public_resource_id`),
  KEY `idx_user_donates_deleted_at` (`deleted_at`),
  KEY `fk_user_donates_user_information` (`user_id`),
  CONSTRAINT `fk_user_donates_public_resource` FOREIGN KEY (`public_resource_id`) REFERENCES `public_resources` (`id`),
  CONSTRAINT `fk_user_donates_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_friends` (
  `user_id` bigint unsigned NOT NULL,
  `friend_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  KEY `fk_user_friends_user_information` (`user_id`),
  KEY `fk_user_friends_friend_information` (`friend_id`),
  CONSTRAINT `fk_user_friends_friend_information` FOREIGN KEY (`friend_id`) REFERENCES `user_informations` (`id`),
  CONSTRAINT `fk_user_friends_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `group_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_groups_deleted_at` (`deleted_at`),
  KEY `fk_user_groups_user_information` (`user_id`),
  KEY `fk_user_groups_group` (`group_id`),
  CONSTRAINT `fk_user_groups_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`),
  CONSTRAINT `fk_user_groups_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `goods_id` bigint unsigned NOT NULL,
  `history_time` datetime(3) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_histories_deleted_at` (`deleted_at`),
  KEY `fk_user_histories_user_information` (`user_id`),
  KEY `fk_user_histories_goods_information` (`goods_id`),
  CONSTRAINT `fk_user_histories_goods_information` FOREIGN KEY (`goods_id`) REFERENCES `goods_informations` (`id`),
  CONSTRAINT `fk_user_histories_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_informations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_number` varchar(10) NOT NULL,
  `avatar` longtext NOT NULL,
  `username` varchar(32) NOT NULL,
  `sex` varchar(2) NOT NULL,
  `introduction` longtext NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone_number` varchar(11) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_number` (`user_number`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone_number` (`phone_number`),
  KEY `idx_user_informations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_operate_informations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `content` longtext NOT NULL,
  `time` datetime(3) NOT NULL,
  `result` tinyint(1) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_operate_informations_deleted_at` (`deleted_at`),
  KEY `fk_user_operate_informations_user_information` (`user_id`),
  CONSTRAINT `fk_user_operate_informations_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `goods_id` bigint unsigned NOT NULL,
  `order_number` varchar(18) NOT NULL,
  `quantity` bigint NOT NULL,
  `real_price` double NOT NULL,
  `real_integral_price` bigint NOT NULL,
  `validity` datetime(3) DEFAULT NULL,
  `create_time` datetime(3) NOT NULL,
  `pay_time` datetime(3) DEFAULT NULL,
  `consume_time` datetime(3) DEFAULT NULL,
  `status` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_number` (`order_number`),
  KEY `idx_user_orders_deleted_at` (`deleted_at`),
  KEY `fk_user_orders_user_information` (`user_id`),
  KEY `fk_user_orders_goods_information` (`goods_id`),
  CONSTRAINT `fk_user_orders_goods_information` FOREIGN KEY (`goods_id`) REFERENCES `goods_informations` (`id`),
  CONSTRAINT `fk_user_orders_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_passwords` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `password` longtext NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `idx_user_passwords_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_user_passwords_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_real_names` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `name` varchar(32) NOT NULL,
  `id_number` varchar(18) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  UNIQUE KEY `id_number` (`id_number`),
  KEY `idx_user_real_names_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_user_real_names_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_rents` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `public_resource_id` bigint unsigned NOT NULL,
  `rent_price` double NOT NULL,
  `rent_integral_price` bigint NOT NULL,
  `deposit` double NOT NULL,
  `rent_time` datetime(3) NOT NULL,
  `back_time` datetime(3) DEFAULT NULL,
  `status` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_rents_deleted_at` (`deleted_at`),
  KEY `fk_user_rents_user_information` (`user_id`),
  KEY `fk_user_rents_public_resource` (`public_resource_id`),
  CONSTRAINT `fk_user_rents_public_resource` FOREIGN KEY (`public_resource_id`) REFERENCES `public_resources` (`id`),
  CONSTRAINT `fk_user_rents_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `role` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_roles_deleted_at` (`deleted_at`),
  KEY `fk_user_roles_user_information` (`user_id`),
  CONSTRAINT `fk_user_roles_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_seats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `seat_id` bigint unsigned NOT NULL,
  `start_time` datetime(3) NOT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `status` varchar(32) NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_seats_deleted_at` (`deleted_at`),
  KEY `fk_user_seats_user_information` (`user_id`),
  KEY `fk_user_seats_seat` (`seat_id`),
  CONSTRAINT `fk_user_seats_seat` FOREIGN KEY (`seat_id`) REFERENCES `seats` (`id`),
  CONSTRAINT `fk_user_seats_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_shopping_carts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `goods_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_shopping_carts_deleted_at` (`deleted_at`),
  KEY `fk_user_shopping_carts_user_information` (`user_id`),
  KEY `fk_user_shopping_carts_goods_information` (`goods_id`),
  CONSTRAINT `fk_user_shopping_carts_goods_information` FOREIGN KEY (`goods_id`) REFERENCES `goods_informations` (`id`),
  CONSTRAINT `fk_user_shopping_carts_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_stars` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `goods_id` bigint unsigned NOT NULL,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_stars_deleted_at` (`deleted_at`),
  KEY `fk_user_stars_user_information` (`user_id`),
  KEY `fk_user_stars_goods_information` (`goods_id`),
  CONSTRAINT `fk_user_stars_goods_information` FOREIGN KEY (`goods_id`) REFERENCES `goods_informations` (`id`),
  CONSTRAINT `fk_user_stars_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_targets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `type_id` bigint unsigned NOT NULL,
  `details` longtext,
  `state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `idx_user_targets_deleted_at` (`deleted_at`),
  KEY `fk_user_targets_target_type` (`type_id`),
  CONSTRAINT `fk_user_targets_target_type` FOREIGN KEY (`type_id`) REFERENCES `target_types` (`id`),
  CONSTRAINT `fk_user_targets_user_information` FOREIGN KEY (`user_id`) REFERENCES `user_informations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci