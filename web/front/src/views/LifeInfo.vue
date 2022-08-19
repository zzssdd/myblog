<template>
  <div>
    <img :src="post.img" class="cover" />
    <div class="container">
      <div class="text-center">
        <h1 class="headline">{{ post.title }}</h1>
        <p>
          发表日期: {{ post.CreatedAt|formatDate }}
        </p>
      </div>
      <hr>
      <div class="frontmatter-markdown post" v-html="post.content">
      </div>
      <Board :flag=2 />
      <PostNavBtn :prev="prevPost" :next="nextPost" :flag=2 />
    </div>
  </div>
</template>

<script>
import PostNavBtn from '@/components/PostNavBtn.vue'
import Board from '@/components/Board.vue'

export default {
  name: 'LifeInfo',
  components: {
    PostNavBtn,
    Board
  },
  data: function () {
    return {
      postId: this.$route.path.split('posts/')[1],
      posts: [],
      post: {},
      prevPost: {
        title: null,
        id: null
      },
      nextPost: {
        title: null,
        id: null
      }
    }
  },
  methods: {
    async getPost() {
      this.postId = this.$route.path.split('life/')[1]
      const { data: res } = await this.$http.get(`life/${this.postId}` )
      this.post=res.data
      const curPostIdx = parseInt(this.postId)
        this.changeTitle(this.post.title)
        try {
          const { data: pre } = await this.$http.get(`life/${curPostIdx+1}` )
          this.prevPost = pre.data
          const { data: next } = await this.$http.get(`life/${curPostIdx-1}` )
          this.nextPost=next.data
        } catch (err) {
          console.log(err) // Handle error
        }
    }
  },
  filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		},
  mounted: function () {
    this.getPost()
  },
  watch: {
    $route (to, from) {
      this.getPost()
    }
  }
}
</script>

<style scoped>
.cover {
  width: 100vw;
}

.margin {
  margin-left: 20px;
}

.post img {  /* limit image max width to 100vw in a post */
  max-width: 100%;
}
</style>
