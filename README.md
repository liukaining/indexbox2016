# IndexBox 平台总体规划

## 背景 

随着互联网行业开源与分享的理念越来越深入人心，每天在微信公众号、专业技术网站、技术会议等各种渠道产出的技术文章越来越多，技术领域纷繁复杂，内容深度参差不齐，然而却没有一个统一的技术文章入口，将每天产出的这些文章能够分门别类的进行整理，将优质内容沉淀起来，让更多的人获得收益。 IndexBox 的初衷就是将每天产出的技术文章能够做好分类，通过一个统一的入口展示给所有用户。平台本身不会产出任何内容，所有的内容都来自于互联网已经发表的技术文章。

## 平台主要功能

- 用户提交技术文章网址
- 平台接受到网址，分析内容，通过相关分类策略，给文章打上一度分类
- 平台对已经一度分类的文章存库
- 平台将已经分类的文章展现到统一入口
- 用户通过统一入口看到文章信息，如果发现文章的分类与平台分类不同，则可以重新选择分类，或者添加一个新的分类，从而给文章打上二度分类
- 分类可以有多级分类，顶级分类最少，每一个顶级分类都有若干二级分类
- 用户可以给文章进行 顶  操作
- 记录每一篇文章通过统一入口跳转的次数，显示到页面上

## 平台技术选型

- 后端框架：golang的 beego 框架 https://beego.me/
- 存储 ：  MySQL 、Redis 
- 部署分发：Docker （待定）
- 服务器：阿里云
- 静态资源：七牛云
- 代码托管：github  `git@github.com:liukaining/indexbox2016.git`


## API  设计

`/indexbox`  首页

`/add_article`  添加文章

`/list_article/[int:category_id]` 按照分类检索文章

`/search/[string:title_abstract]`  文章题目或者摘要搜索

`/del_article/[int:article_id]` 删除文章


----------


`/add_category` 添加分类

`/list_category_all` 所有分类列表

`/list_category_sub/[int:category_id]` 某个分类的子分类

`/list_category_level/[int:level_id]` 某个等级下的分类

`/del_category/[int:category_id]` 删除某个分类，只能删除没有子分类的分类


----------

`/add_record/[int:article_id]/[int:category_id]` 给某一个文章添加一个分类

`/update_record/[int:article_id]/[int:category_id]` 给某一个文章修改分类

`/get_record/[int:article_id]` 查看某一个文章的分类修改记录

`/get_record_all` 查看所有的记录


----------

`/count/[int:article_id]`  统计某个文章从 indexbox 平台跳转的次数
----------

## 部署方式

前提条件：安装好 bee 和 beego 


git clone xxx


bee run 
