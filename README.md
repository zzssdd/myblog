# myblog

#### 介绍
本项目是一个个人博客系统，能够实现后台写文章、发动态、发布项目、文章分类管理、评论管理，以及前台文章的分页展示、分类展示、动态的展示以及评价，还有项目的展示，以及留言和个人简历的展示。

#### 软件架构
1.前端采用了vue框架，前台页面继承了vuesax组件库、后台继承了element-ui组件库，对后端发送axios请求获取数据（前台界面参考于别人）
2.后端采用了golang的gin框架，采用gorm框架对mysql进行操作，采用了日志库和日志分割作为gin的中间件来记录日志，利用cors进行跨域访问，采用jwt进行后台登陆的鉴权
3.将项目部署在docker容器中


#### 安装教程

docker下的安装（无需go环境，需要安装docker，并在docker中拉取启动mysql:5.6并创建数据库myblog）：
1. 修改config下的config.ini文件
2. 进入文件所在的文件夹，执行docker build -f Dockerfile -t myblog .
3. 执行 docker run -p 3000:3000 -d myblog
4. 前台网页访问localhost:3000，后台网页访问localhost:3000/admin

直接安装访问（需要配置go环境，并安装mysql：5.6，创建数据库myblog）：
1. 修改config.ini文件
2. 执行命令go mod download下载go依赖包
3. go run main.go执行
4. 前台网页访问localhost:3000，后台网页访问localhost:3000/admin

#### 使用说明

1.  需要在windows下或者docker下安装mysql
2.  如果主机名不是localhost，需要修改config/config.ini并修改前端axios的baseURL后重新打包

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
