<template>
  <div>
    <vs-row justify="space-between">
      <vs-col w="2" v-if="prev.ID">
        <vs-button
          size="large"
          block
          border
          :to="flag===1?`/article/${prev.ID}`:`/life/${prev.ID}`"
        >
          <i class="bx bx-left-arrow-alt" style="margin-right: 10px"></i>
          {{ prev.title }}
        </vs-button>
      </vs-col>
      <vs-col w="2"></vs-col>
      <vs-col w="2" v-if="next.ID" class="align-right">
        <vs-button
          size="large"
          block
          border
          :to="flag===1?`/article/${next.ID}`:`/life/${next.ID}`"
        >
          {{ next.title }}
          <i class="bx bx-right-arrow-alt" style="margin-left: 10px"></i>
        </vs-button>
      </vs-col>
    </vs-row>
  </div>
</template>

<script>
export default {
  name: 'PostNavBtn',
  props: [
    'next',
    'prev',
    'flag'
  ]
}
</script>

<style scoped>
.align-right {
  text-align: right !important;
}
</style>
