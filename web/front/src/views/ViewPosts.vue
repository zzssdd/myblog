<template>
  <div class="container">
    <h1 class="text-center">文章</h1>
    <hr />
    <div v-if="getPosts().length">
      <br />
      <div class="grid-3_xs-1_sm-2_md-2">
        <div
          :key="index"
          v-for="(post, index) in getPosts()"
          class="col"
        >
        <PostCard :post="post" :flag=1 />
        </div>
      </div>
      <div class="center con-pagination">
        <vs-pagination v-model="curPage" :length="allPosts.length" not-margin progress />
      </div>
    </div>
    <Nothing v-else />
  </div>
</template>

<script>
import PostCard from '@/components/PostCard.vue'
import Nothing from '@/components/Nothing.vue'

export default {
  name: 'ViewPosts',
  components: {
    PostCard,
    Nothing
  },
  data: function () {
    return {
      allPosts:[],
      curPage: 1,
    }
  },
  mounted: function () {
    this.changeTitle('文章')
 },
  methods: {
    async getList(){
        const { data: res } = await this.$http.get('article/list' )
        this.allPosts=res.data
        const posts = this.allPosts
        this.allPosts = []
        for (let i = 0; i < posts.length; i += 6) this.allPosts.push(posts.slice(i, i + 6)) 
    },
     getPosts: function () {
      try {
        return this.allPosts[this.curPage - 1].slice()
      } catch (e) {
        return []
      }
    }
    // getPosts: function () {
    //   try {
    //     return this.allPosts[this.curPage - 1].slice().reverse()
    //   } catch (e) {
    //     return []
    //   }
    // }
  },
  created(){
    this.getList()
  }
}
</script>
