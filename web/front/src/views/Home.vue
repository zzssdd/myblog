<template>
  <div class="main_page">
    <div class="container">
      <div class="text-center">
        <div class="center">
          <Avatar size="300" home="true" />
        </div>
        <h1>Yogen</h1>
        <h4 style="color: #666666"></h4>
        <h3>我虽无意逐鹿，却知苍生苦楚</h3>
        <h3>只愿荡涤四方，护得浮世一隅</h3>
      </div>
      <h2>往日文章</h2>
      <hr/>
      <div class="grid-3_xs-1_sm-2_md-2" v-if="posts.length">
        <div
          :key="index"
          v-for="(post, index) in posts.slice(Math.max(posts.length - 6, 0))"
          class="col"
        >
          <PostCard :post="post" :flag=1 />
        </div>
      </div>
      <Nothing v-else />
    </div>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar.vue'
import PostCard from '@/components/PostCard.vue'
import Nothing from '@/components/Nothing.vue'

export default {
  name: 'Home',
  components: {
    Avatar,
    PostCard,
    Nothing
  },
  methods:{
    async getPost(){
      const {data:res}=await this.$http.get(`article/list`)
      this.posts=res.data
    },
  },
  data: function () {
    return {
      posts:[],
    }
  },
  created(){
    this.getPost()
  },
  mounted: function () {
    this.changeTitle('主页')
  }
}
</script>
