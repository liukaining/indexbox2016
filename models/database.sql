CREATE DATABASE IF NOT EXISTS `indexbox`;
USE `indexbox`
SET CHARACTER_SET_CLIENT = utf8;
CREATE TABLE article_info(
  `id` int(10) unsigned NOT NULL  AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `abstract` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '摘要',
  `title` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文章标题',
  `url` VARCHAR(1000) NOT NULL DEFAULT '' COMMENT '文章URL',
  `insert_time` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '入库时间',
  `insert_user` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '添加人',
  `is_del` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '是否删除',
  `ext` VARCHAR(1000) NOT NULL DEFAULT 'ext' COMMENT '扩展字段'
 );

CREATE TABLE category_base_info(
  `id` int(10) unsigned NOT NULL  AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `level` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '级别',
  `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '分类名字',
  `father_category` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '父分类',
  `insert_time` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '入库时间',
  `insert_user` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '添加人'
);

CREATE TABLE category_article_record(
  `id` int(10) unsigned NOT NULL  AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `article_id` int(10) unsigned NOT NULL  DEFAULT 0  COMMENT '文章ID',
  `category_id` int(10) unsigned NOT NULL DEFAULT 0  COMMENT '分类ID',
  `insert_time` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '入库时间',
  `is_effective` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '是否生效'
);