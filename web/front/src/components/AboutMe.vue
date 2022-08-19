<template>
  <div class="container text-center">
    <h1>
      关于我
    </h1>
    <hr />
    <h3>欢迎来到Yogen的博客！我是Yogen，一个来自于山东的22岁的男性青年。</h3>
    <h3>目前在是还是一名软件工程研究生，未来想去做一名golang开发工程师。</h3>
    <h3>热爱音乐，喜欢发呆。讨厌内卷，但不会摆烂。对现状基本知足，对未来充满期待。</h3>
    <h3>如果有啥实习或工作机会，恳请各位大佬联系我。也欢迎各位朋友添加好友来一起学习，共同进步</h3>
  </div>
</template>

<script>
export default {
  name: 'AboutMe',
  components: {
  },
  mounted: function () {
    this.changeTitle('About')
  }
}
</script>
