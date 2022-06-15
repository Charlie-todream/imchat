// 多模板渲染
go get github.com/gin-contrib/multitemplate 

// 配置
go get github.com/spf13/cast
go get github.com/spf13/viper

// 日志 lumberjack 是一套滚动日志的实现方案，帮助我们管理日志文件。
go get go.uber.org/zap
go get gopkg.in/natefinch/lumberjack.v2


CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` char(11) NOT NULL DEFAULT '0',
  `password` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) DEFAULT NULL,
  `sex` varchar(2) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `online` tinyint(2) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `memo` varchar(255) DEFAULT NULL,
  `createat` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `contacts`;
CREATE TABLE `contacts`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ownerid` int(11) NOT NULL DEFAULT 0 COMMENT '// 记录是谁的',
  `dstobj` int(11) NOT NULL COMMENT '对端信息',
  `cate` int(11) NOT NULL COMMENT '什么类型',
  `memo` varchar(255) NOT NULL DEFAULT 0 COMMENT '备注',
  `createat` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of contacts
-- ----------------------------
INSERT INTO `contacts` VALUES (1, 1, 2, 1, 11, '2022-06-15 13:40:48');
INSERT INTO `contacts` VALUES (2, 1, 3, 1, 11, '2022-06-14 13:42:18');
INSERT INTO `contacts` VALUES (3, 1, 4, 1, 22, '2022-06-14 13:43:28');

SET FOREIGN_KEY_CHECKS = 1;


CREATE TABLE `communities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `ownerid` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(255) DEFAULT NULL,
  `cate` int(11) DEFAULT NULL,
  `memo` varchar(255) DEFAULT NULL,
  `createat` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;