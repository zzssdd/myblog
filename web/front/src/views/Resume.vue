<template>
      <div class="container">
        <div class="leftBox">
            <div class="head">
                <img src="../assets/avater.jpg" alt="" style="width:136px;height:136px" />
                <h1>张勇健的简历</h1>
            </div>
            <div class="basicInfo">
                <h2 class="leftTitle">基本信息</h2>
                <ul>
                    <li><span>名字：</span>张勇健</li>
                    <li><span>性别：</span>男</li>
                    <li><span>本科毕业院校：</span>烟台大学</li>
                    <li><span>本科就读专业：</span>计算机科学与技术</li>
                    <li><span>毕业时间：</span>2022年7月</li>
                    <li><span>研究生就读院校：</span>武汉理工大学</li>
                    <li><span>研究生就读专业：</span>软件工程专硕</li>
                    <li><span>个人博客：</span><a href="http://www.zhangyongjian.top" target="_blank">www.mangoya.cn</a></li>
                </ul>
            </div>
            <div class="contact">
                <h2 class="leftTitle">联系方式</h2>
                <ul>
                    <li><span>电话：</span>18053801003</li>
                    <li><span>邮箱：</span>zhangyongjian2021@163.com</li>
                    <li><span>微信：</span>I-m_zyj</li>
                    <li><span>QQ：</span>1654622146</li>
                </ul>
            </div>
            <div class="application">
                <h2 class="leftTitle">应聘岗位</h2>
                <p>golang开发工程师</p>
            </div>
            <div class="aboutme">
                我是一个对golang由衷热爱、有趣的工程师。
                希望在开发领域，让我所掌握的技能落实到实处，献上一点绵薄之力！
            </div>
        </div>
        <!-- <div class="rightBox">

            <h3 class="rightTitle"><box-icon type='solid' name='rocket'></box-icon>项目与工作经验</h3>
            <h4><img src="img/tal.png" alt="" />北京世纪好未来科技教育有限公司（2018.04——至今）<a href="http://www.100tal.com/"
                    target="_blank"><i class="fa fa-fw fa-external-link"></i></a></h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>前端埋点管理平台</h5>
                    <p>
                        前端项目埋点录入编辑、sdk上报、数据看板于一体的管理平台，规范埋点和自动生成通用数据看板，并可快速生成自定义看板。包含登录/添加协同用户模块，上传oss/版本管理模块，通用echarts图表和自定义图表模块等.
                    </p>
                    <p>
                        技术栈：vue/koa2/mongodb
                    </p>
                    <p>
                        1、担任产品、设计、交互、前端、后端多个角色，独自完成平台的设计和搭建；
                    </p>
                    <p>
                        2、开发兼容web/node、微信支付宝百度头条等小程序的日志上报sdk；

                    </p>
                    <p>
                        3、通过koa2框架搭建node服务，并实现oauth第三方授权登录，mongodb存储数据等功能
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>可视化编辑平台</h5>
                    <p>
                        通过拖拽搭建h5页面，结合创建、编辑、上线于一体的编辑平台。支持复杂的自定义组件生态，模板定制等功能，并集成用户权限管理、项目数据看板、组件管理等功能。
                    </p>
                    <p>
                        技术栈：react/koa2/mongodb
                    </p>
                    <p>
                        1、平台功能的开发和迭代，包括模板、组件、图片资源管理
                    </p>
                    <p>
                        2、脚手架功能迭代，改造webpack5 federation抽离公共组件和controller;
                    </p>
                    <p>
                        3、业务模板、通用组件的开发，包括拼团模板、支付购买组件等。
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>公益直播讲座</h5>
                    <p>web讲座主要用于免费直播回放视频观看，包括列表、预约、直播、聊天答题和回放功能。原有的项目技术落后，代码可读性差，影响新功能的开发以及后续维护</p>
                    <p>
                        技术栈：vue
                    </p>
                    <p>
                        1、将原有的 jQuery 技术栈替换成 Vue 模块化、组件化开发各个功能 ；
                    </p>
                    <p>
                        2、替换原有的flash播放方式，改换hls模式播放，使用websocket实现聊天、答题等交互；
                    </p>
                    <p>
                        3、增加直播间聊天降级开关配置，接口做缓存处理等降级功能，验证百万用户同时在线流量；
                    </p>
                </li>
            </ul>

            <h4><img src="img/xingyuan.png" alt="" />北京行圆汽车信息技术有限公司（2017.1——至今）<a href="http://www.xingyuanauto.com/"
                    target="_blank"><i class="fa fa-fw fa-external-link"></i></a></h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>商务产品<a href="http://flow.xingyuanauto.com/webMove-h5/" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>
                        负责商业营销类和功能类的H5前端主要研发工作
                    </p>
                    <p>
                        使用 skrollr 、swiper、three.js、create.js等插件和AnimateCC 动画软件来实现前端动画效果，使用less编译css, 解决浏览器兼容性问题
                    </p>
                    <p>
                        使用ajax异步请求数据，完成支付定位、微信平台开发等功能以及前后端数据交互
                    </p>
                    <p>
                        使用vue.js和webpack开发单页面功能性项目
                    </p>
                    <p>
                        持续优化相关产品的质量、性能、用户体验,构建H5高级动效和功能小插件
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>内部项目</h5>
                    <p>
                        大型电商平台系统——<a href="http://www.chitunion.com/" target="_blank">赤兔联盟</a>的前端开发，该项目后端由.net 和java
                        实现，主要负责部分页面的ajax数据开发和功能实现，以及网站公共功能和模块的开发
                    </p>
                    <p>
                        使用layui实现样式，ejs以及jquery实现前端数据的绑定，以及其他jquery插件完成相关组件的开发
                    </p>
                </li>
            </ul>
            <h4><img src="img/zhongshang.png" alt="" />中商情大数据股份有限公司（2014.12——2016.12）<a href="http://www.askci.com/#"
                    target="_blank"><i class="fa fa-fw fa-external-link"></i></a></h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>中商情报网<a href="http://www.askci.com/" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>
                        该项目是中商的新闻门户网站，大量的产业以及新闻数据展示。负责前端开发工作，包括页面布局，数据库渲染，分为 <a href="http://www.askci.com/"
                            target="_blank">PC版</a>和 <a href="http://m.askci.com/" target="_blank">H5移动版</a>
                    </p>
                    <p>
                        使用jquery、swiper等插件，实现页面布局和交互效果，优化网站的性能和用户体验
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>八方查数据网站<a href="http://8.askci.com/" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>
                        该项目主要是中商企业注册等工商信息查询产品，负责canvas动画，响应式样式HTML5/CSS，实现AJAX请求，数据动态加载等前端开发工作，以及多条件嵌套查询等功能
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>千数堂<a href="http://d.askci.com/" target="_blank"><i class="fa fa-fw fa-external-link"></i></a>
                    </h5>
                    <p>该项目为商情数据的数据电商模块业务，商品为产品为各行业的产业数据，包括登录注册、首页、详情、购买、用户中心等页面模块，完整的电商平台</p>
                    <p>
                        负责canvas动画，响应式样式HTML5/CSS，实现多功能的交互和逻辑处理，相关组件的开发和封装
                    </p>
                    <p>
                        使用 jquery easyui 实现组件样式，使用ajax实现动态数据交互
                    </p>
                    <p>
                        <img src="img/qianshutang.png" alt="" />
                    </p>
                </li>
            </ul>
            <h4>个人项目</h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>美食美荟小程序 <a href="https://github.com/Aimee1608/foodEveryDay" target="_blank"><i
                                class="fa fa-fw fa-github"></i></a> <a
                            href="http://www.mangoya.cn/src/img/aimee/meishi.jpg" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>
                        微信小程序出来后，由于本人有做饭的兴趣爱好，就开发了一个菜谱小程序
                    </p>
                    <p>
                        针对于普通用户可以发布、点赞、收藏作品，还可以根据分类以及关键字搜索作品，对于管理员除了以上功能外，另外添加了审核用户发布的作品功能，可通过或者驳回，详细体验扫描下方二维码;后端使用PHP开发
                    </p>
                    <p>
                        <img src="img/xiaochengxu.jpg" alt="" class="smallImg" />
                        <img src="img/foodwx.png" alt="" class="bigImg" />
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>美食美荟网站 <a href="https://github.com/Aimee1608/foodWeb" target="_blank"><i
                                class="fa fa-fw fa-github"></i></a> <a href="http://food.mangoya.cn" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>
                        根据美食美荟小程序，使用vue.js改版了一个网页版的美食美荟网站，兼容pc和移动端，但只有数据展示，没有用户登录模块了，后端接口同美食美荟相同
                    </p>
                    <p>
                        使用 vue.js 框架和Element UI来实现前端组件及数据交互，使用node.js的 webpack 构建开发打包
                    </p>
                    <p>
                        <img src="img/foodweb.png" alt="" />

                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>个人博客 <a href="https://github.com/Aimee1608/myblogvue" target="_blank"><i
                                class="fa fa-fw fa-github"></i></a> <a href="http://www.mangoya.cn" target="_blank"><i
                                class="fa fa-fw fa-external-link"></i></a></h5>
                    <p>用于技术分享和总结的个人博客，兼容pc和移动端，用户可通过github oauth认证登录博客，实现收藏点赞以及留言评论回复等功能，还可以展示个人链接，未登录状态可以以游客方式留言等功能</p>
                    <p>后端使用koa2/mongodb开发</p>
                    <p>
                        挂在阿里云服务器上，通过远程桌面来维护修改
                    </p>
                    <p>
                        <img src="img/blog.png" alt="" />
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>Aimee resume <a href="https://github.com/Aimee1608/aimeeResume" target="_blank"><i
                                class="fa fa-fw fa-github"></i></a></h5>
                    <p>这个简历的源码,设计优雅、内容完善的静态简历页面,可下载world文档</p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>其他</h5>
                    <p>日常中研究封装的一些小插件，360度查看产品展示，jquery版手风琴，jqueryEmoji表情包，加载遮罩模板封装，折叠留言楼层-递归等，可查看 <a
                            href="https://github.com/Aimee1608" target="_blank"><i
                                class="fa fa-fw fa-github"></i></a>上的源代码</p>
                </li>
            </ul>

            <h3 class="rightTitle"><i class="fa fa-fw fa-rocket"></i>专业技能</h3>
            <h4>前端</h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>HTML/CSS</h5>
                    <p>
                        能够编写语义化的 HTML，模块化的 CSS，完成较复杂的布局,熟练使用css3属性完成炫酷的H5动画效果
                    </p>
                    <p>
                        熟悉Less等CSS预处理和模块化工具
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>JavaScript</h5>
                    <p>
                        ES6,ES7 等 Javascript 新特性
                    </p>
                    <p>
                        熟练使用js封装构造函数类插件及组件
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>框架插件类</h5>
                    <p>
                        前端框架 Vue、React 全家桶
                    </p>
                    <P>
                        前端主流 UI 框架 element-ui/vant/antd/echarts等
                    </P>
                    <p>
                        前端构建工具 Webpack 、rollup以及相关插件
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>小程序</h5>
                    <p>
                        小程序原生和uniapp、mpx、taro框架开发
                    </p>
                </li>
            </ul>
            <h4>后端</h4>
            <ul class="itemList">
                <li>
                    <div class="circle"></div>
                    <h5>Node.js</h5>
                    <p>
                        熟练使用node.js框架koa2/mongodb数据库架设服务
                    </p>
                </li>
                <li>
                    <div class="circle"></div>
                    <h5>nginx配置</h5>
                    <p>
                        熟悉nginx相关配置实现history路由模式和代理等
                    </p>
                </li>
            </ul>
        </div> -->
        <div class="rightBox" style="width:3000px;height:2000px">
           <h3 class="rightTitle"><box-icon type='solid' name='rocket'></box-icon>项目与工作经验</h3>
          <h4>个人项目</h4>
            <ul class="itemList">
                 <li>
                    <div class="circle"></div>
                    <h5>基于Python的房价预测系统<a href="http://www.100tal.com/"
                    target="_blank"><box-icon type='solid' name='navigation'></box-icon></a></h5>
                    <p>
                      题目来源于kaggle上的House Prices赛题，对1460个数据样本进行了大量的数据分析和预处理工作。然后<br>建立多种模型进行预测，最后按照模型融合的方式进行模型的优化，取得的最好成绩为129/4314的成绩，世界前3%。
                    </p>
                    <p>
                        技术栈：Python
                    </p>
                    <p>
                        1、数据探索：包括对数据的分布可视化展示，以及特征跟目标值的相关性分析。
                    </p>
                    <p>
                        2、特征工程：包括对目标值的正态变换、异常值的处理、缺失数据的填补、标签变量的转换、新特征的构建、倾斜特<br>征的转换。
                    </p>
                    <p>
                        3、模型的构建：根据翻阅文献和尝试，最后选取了Lasso回归、岭回归、弹性网络回归、梯度提升回归、<br>XGBoost、LightGBM模型进行<br>预测。并用网格搜索法对模型进行了调参。
                    </p>
                    <p>
                        4、模型的优化：选用了Stacking融合和Blending融合两种方式对模型进行了融合。
                    </p>
                </li>

                <li>
                    <div class="circle"></div>
                    <h5>分布式日志收集项目<a href="http://www.100tal.com/"
                    target="_blank"><box-icon type='solid' name='navigation'></box-icon></a></h5>
                    <p>
                     用于分布式系统的日志收集和可视化展示。
                    </p>
                    <p>
                        技术栈：Golang/Beego/Mysql/Kafka/Etcd/Elasticsearch/Kibana
                    </p>
                    <p>
                        1、基于Beego框架的Web管理平台：用于管理日志收集相关的配置信息。
                    </p>
                    <p>
                        2、Etcd：用于存储各个日志收集代理的要收集的日志文件信息的配置信息。
                    </p>
                    <p>
                        3、Kafka：用户缓存日志收集代理收集到的日志信息。
                    </p>
                    <p>
                        4、ElasticSearch：将Kafka中数据取出建立搜索引擎。
                    </p>
                     <p>
                        4、Kibana：用于可视化展示ElasticSearch中的数据。
                    </p>
                </li>

                <li>
                    <div class="circle"></div>
                    <h5>个人博客<a href="http://www.100tal.com/"
                    target="_blank"><box-icon type='solid' name='navigation'></box-icon></a></h5>
                    <p>
                       前后端分离技术开发了个人博客项目。主要功能有：文章的展示和分类，生活动态、个人项目、个人简历的展示，<br>留言板，博客的后台管理等博客的基本功能。
                    </p>
                    <p>
                        技术栈：Golang/Gin/Mysql/Vue/Docker/Gorm
                    </p>
                    <p>
                        1、前端采用了vue框架，前台页面集成了Vuesax组件库，后台页面集成了Element ui组件库。
                    </p>
                    <p>
                        2、后端采用了golang的gin框架，用Gorm对Mysql数据库进行操作，并使用了日志文件存储和分割的中间件来记录<br>日志。
                    </p>
                    <p>
                        3、根据jwt的方式进行后台登录身份的验证。
                    </p>
                    <p>
                      4.利用Docker容器技术将项目部署在腾讯云服务器上。
                    </p>
                </li>

                <li>
                    <div class="circle"></div>
                    <h5>雷霆战机<a href="http://www.100tal.com/"
                    target="_blank"><box-icon type='solid' name='navigation'></box-icon></a></h5>
                    <p>
                      玩家能够选择战机和地图，能够依照分数自动生成大中小型敌方战机，能够根据时间帧生成两种补给<br>（一种强化我机子弹，一种增加大招数量），实现了游戏中所应有的碰撞检测以及特效和声效的播放，游戏结束时<br>能够显示历史最高分和当前分数。
                    </p>
                    <p>
                        技术栈：Python/Pygame/Mysql
                    </p>
                </li>
                <li>
                </li>
            </ul>
          <br>
          <h4>奖项情况</h4>
            <ul>
             <li>
                    <h5>2020年第十七届五一数学建模竞赛一等奖(2020.06)</h5>
            </li>
            <li>
                    <h5>第十一届蓝桥杯全国软件和信息技术专业人才大赛山东赛区C/C++程序设计大学B组一等奖(2020.10)</h5>
            </li>
            <li>
                    <h5>全国大学生数学建模竞赛三等奖(2020.10)</h5>
            </li>
            <li>
                    <h5>中国高校计算机大赛-团体程序设计天梯赛铜牌(2020.11)</h5>
            </li>
            <li>
                    <h5>优秀学生奖学金(2020.12)</h5>
            </li>
            <li>
                    <h5>2021年全国大学生数据统计与分析竞赛本科生组三等奖(2021.05)</h5>
            </li>
            <li>
                    <h5>校级优秀毕业生(2022.05)</h5>
            </li>
            </ul>
          <br>
          <h4>专业技能</h4>
           <ul>
            <li>
                    <h5>1.熟悉golang语言。编写过beego框架和gin框架的web后端</h5>
            </li>
            <li>
                    <h5>2.有gorm框架基础，能够灵活地对mysql数据库进行操作</h5>
            </li>
            <li>
                    <h5>3.了解Docker容器技术，能够将项目部署在Docker容器上。</h5>
            </li>
            <li>
                    <h5>4.对Kafka、Etcd、ElasticSearch有一定的认知</h5>
            </li>
            <li>
                    <h5>5.对前端语言html、css、JavaScript以及前端框架vue有一定的了解</h5>
            </li>
             <li>
                    <h5>6.有Python、c、c++等多种语言的基础</h5>
            </li>
            <li>
                    <h5>7.掌握常用的数据结构和算法，不错的计算机专业知识基础，对个别的机器学习模型有一定的认识。</h5>
            </li>
           </ul>
        </div>
    </div>
</template>

<script>
export default {
  name: 'Resume',
  data: function () {
    return {
    }
  },
  components: {
  },
  mounted: function () {
    this.changeTitle('简历')
  }
}
</script>

<style scoped>
/*reset*/
html,
body,
div,
span,
applet,
object,
iframe,
h1,
h2,
h3,
h4,
h5,
h6,
p,
blockquote,
pre,
a,
abbr,
acronym,
address,
big,
cite,
code,
del,
dfn,
em,
img,
ins,
kbd,
q,
s,
samp,
small,
strike,
strong,
sub,
sup,
tt,
var,
b,
u,
i,
center,
dl,
dt,
dd,
ol,
ul,
li,
fieldset,
form,
label,
legend,
table,
caption,
tbody,
tfoot,
thead,
tr,
th,
td,
article,
aside,
canvas,
details,
embed,
figure,
figcaption,
footer,
header,
menu,
nav,
output,
ruby,
section,
summary,
time,
mark,
audio,
video {
  margin: 0;
  padding: 0;
  border: 0;
  font: inherit;
}
body {
  font-family: Arial, "STHeiti", Helvetica, sans-serif;
  background: #fff;
  font-size: 14px;
  height: 100%;
  -webkit-tap-highlight-color: transparent;
}
table {
  border-collapse: collapse;
  border-spacing: 0;
}
fieldset,
img {
  border: 0;
}
legend {
  display: none;
}
h5{
  color:black !important;
}
address,
caption,
cite,
code,
dfn,
em,
strong,
th,
var {
  font-style: normal;
  font-weight: normal;
}
ol,
ul {
  list-style: none;
}
caption,
th {
  text-align: left;
}
h1,
h2,
h3,
h4,
h5,
h6 {
  font-size: 100%;
  font-weight: normal;
}
q:before,
q:after {
  content: '';
}
abbr,
acronym {
  border: 0;
}
html {
  -webkit-text-size-adjust: none;
  width: 100%;
  height: 100%;
}
/*????iphone??safari????????????*/
input[type="text"],
input[type="button"],
input[type="submit"],
input[type="search"] {
  -webkit-appearance: none;
  border-radius: 0;
}
a {
  text-decoration: none;
  color: inherit;
}
html,
body {
  position: relative;
  color: #333;
  background: #EBEBEB;
}
a {
  color: #df2050;
}
a:hover {
  color: #df2050;
  text-decoration: underline;
}
/****************************����*****************************/
.container {
  width: 100%;
  height: auto;
  margin: 5px auto;
  position: relative;
  box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
  overflow: hidden;
  border-radius: 4px;
}
.container:after {
  content: "";
  display: block;
}
.leftBox {
  width: 370px;
  height: 100%;
  background: #394A5A;
  padding: 15px;
  position: absolute;
  top: 0;
  bottom: 0;
  box-sizing: border-box;
}
.leftBox .head {
  width: 100%;
  text-align: center;
  margin: 25px 0;
}
.leftBox .head img {
  border-radius: 50%;
  padding: 4px;
  background: #fff;
}
.leftBox .head h1 {
  color: #fff;
  font-size: 28px;
  font-weight: 400;
}
.leftBox div {
  position: relative;
}
.leftBox .leftTitle {
  height: 50px;
  line-height: 50px;
  width: 100%;
  position: relative;
  box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
  border-color: #6B3645;
  background-color: #B05051;
  margin: 5px 0 15px 30px;
  font-size: 22px;
  border-radius: 0 4px 0 0;
  display: inline-block;
  padding-left: 20px;
  box-sizing: border-box;
  color: #fff;
}
.leftBox .leftTitle::before {
  content: "";
  display: block;
  position: absolute;
  left: -19px;
  top: 0;
  width: 20px;
  height: 50px;
  background-color: #B05051;
  clip-path: polygon(80% 0%, 100% 0%, 100% 100%, 80% 100%, 0% 50%);
  -webkit-clip-path: polygon(80% 0%, 100% 0%, 100% 100%, 80% 100%, 0% 50%);
}
.leftBox .leftTitle:after {
  position: absolute;
  content: "";
  top: 100%;
  right: 0;
  border-top: 0 solid transparent;
  border-left-width: 15px;
  border-left-color: inherit;
  border-left-style: solid;
  border-bottom: 15px solid transparent;
  border-right: 0 solid transparent;
  width: 0;
  height: 0;
}
.leftBox ul {
  padding-left: 50px;
  margin-bottom: 30px;
}
.leftBox ul li {
  color: #ddd;
  margin: 5px 0;
}
.leftBox ul li span {
  font-size: 16px;
  color: #ABADB0;
  font-weight: bold;
}
.leftBox p {
  padding-left: 50px;
  color: #eee;
  font-size: 16px;
}
.leftBox .contact .leftTitle {
  border-color: #5A6238;
  background-color: #BDC293;
}
.leftBox .contact .leftTitle::before {
  background-color: #BDC293;
}
.leftBox .application .leftTitle {
  border-color: #887334;
  background-color: #DAB652;
}
.leftBox .application .leftTitle::before {
  background-color: #DFB651;
}
.leftBox .aboutme {
  padding: 40px 20px 0 20px;
  color: #eee;
  line-height: 1.5;
  text-indent: 2em;
}
.rightBox {
  width: 100%;
  height: auto;
  box-sizing: border-box;
  padding: 10px 40px 40px 400px;
  background: #fff;
}
.rightBox .rightTitle {
  font-size: 24px;
  color: #394A5A;
  line-height: 30px;
  border-bottom: 1px solid #394A5A;
  margin: 30px 0 15px -5px;
}
.rightBox h4 {
  font-size: 19px;
  color: #01579B;
  margin: 0 0 15px 5px;
  font-weight: bold;
}
.rightBox h4 img {
  height: 30px;
  margin-right: 10px;
  vertical-align: middle;
}
.rightBox h4 a {
  vertical-align: middle;
}
.rightBox .itemList {
  border-left: 2px solid #394A5A;
  margin-left: 15px;
}
.rightBox .itemList li {
  position: relative;
  top: -5px;
  margin: 0 0 20px 20px;
}
.rightBox .itemList li p {
  font-size: 15px;
  color: #616161;
  line-height: 23px;
  margin-bottom: 5px;
}
.rightBox .itemList li p img {
  max-width: 100%;
  max-height: 200px;
  object-fit: cover;
  vertical-align: middle;
  cursor: -webkit-zoom-in;
  cursor: zoom-in;
}
.rightBox .itemList li p .smallImg {
  width: 100px;
  height: 100px;
}
.rightBox .circle {
  height: 14px;
  width: 14px;
  -webkit-border-radius: 100%;
  border-radius: 100%;
  background: #394A5A;
  float: left;
  position: absolute;
  left: -28px;
  top: 5px;
}
.rightBox h5 {
  line-height: 30px;
  font-size: 18px;
  bottom: 3px;
  font-weight: bold;
}
.footer {
  text-align: center;
  padding: 20px 0;
}
.footer .fa-heart {
  color: #df2050;
}
.review {
  width: 100vw;
  height: 100vh;
  position: fixed;
  left: 0;
  top: 0;
  background: #fff;
  text-align: center;
  display: none;
  cursor: -webkit-zoom-out;
  cursor: zoom-out;
}
.review img {
  max-width: 70%;
  max-height: 100vh;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  -webkit-transform: translate(-50%, -50%);
}
@media screen and (min-width: 961px) {
  .container {
    width: 80%;
  }
}
@media screen and (max-width: 960px) and (min-width: 820px) {
  .container {
    width: 100%;
    padding: 0 10px;
    box-sizing: border-box;
    box-shadow: none;
  }
  .rightBox {
    margin: 0;
    border-radius: 0 4px 4px 0;
  }
  .leftBox {
    border-radius: 4px 0 0 4px ;
  }
}
@media screen and (max-width: 819px) {
  .container {
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
    box-shadow: none;
  }
  .leftBox {
    width: 100%;
    box-sizing: border-box;
    position: relative;
    border-radius: 4px 4px 0 0;
  }
  .rightBox {
    width: 100%;
    padding-left: 40px;
    box-sizing: border-box;
    border-radius: 0 0 4px 4px ;
  }
}
@media screen and (max-width: 376px) {
  .container {
    margin: 0;
    width: 100%;
    padding: 0;
    box-sizing: border-box;
    box-shadow: none;
    border-radius: 0 ;
  }
  .leftBox {
    width: 100%;
    box-sizing: border-box;
    position: relative;
    border-radius: 0;
  }
  .rightBox {
    width: 100%;
    padding: 15px;
    box-sizing: border-box;
    border-radius: 0 ;
  }
}

</style>
