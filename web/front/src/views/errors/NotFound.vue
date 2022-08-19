<template>
  <div class="container text-center">
    <h1>404</h1>
    <hr />
    <br />
    您要查找的页面不存在
    <div class="center">
      <vs-button size="large" @click="$router.push('/')">返回主页</vs-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
  mounted: function () {
    this.changeTitle('404')
  }
}
</script>
