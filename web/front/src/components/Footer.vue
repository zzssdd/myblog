<template>
  <div class="footer">
    <div class="container text-center">
      <h3>Yogen的博客</h3>
      <br />
      <p v-if="curYear === 2022">
        &copy; {{ curYear }} .&nbsp;
        前台页面参考于 <a href="https://samzhangjy.github.io/#/" target="_blank">VueBlogger</a>.
      </p>
      <p v-else>
        &copy; 2022 - {{ curYear }} {{ config.name }}.&nbsp;
        Powered by <a href="https://samzhangjy.github.io/#/" target="_blank">VueBlogger</a>.
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data: function () {
    return {
      curYear: new Date().getFullYear()
    }
  },
  methods: {
    getYear: function () {
      return new Date().getFullYear()
    }
  }
}
</script>
