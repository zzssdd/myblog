<template>
  <div class="container">
    <hr>
    <br />
    <vs-row id="row">
      <vs-col
        :key="index"
        v-for="(post, index) in ArtList"
        lg="4"
        sm="12"
        md="6"
        style="margin-bottom: 30px"
        id="col"
      >
        <PostCard :post="post" :flag="1"/>
      </vs-col>
    </vs-row>
  </div>
</template>

<script>
import PostCard from '@/components/PostCard.vue'

export default {
  name: 'ViewTag',
  components: {
    PostCard
  },
  data: function () {
    return {
      name: this.$route.params.tag,
      ArtList:[]
    }
  },
  methods:{
    async getList(){
      const { data: res } = await this.$http.get(`articlecate/${this.name}` )
      this.ArtList=res.data
    }
  },
  computed: {
    tag: function () {
      for (var i = 0; i < this.tags.length; i++) { // match current tag
        if (this.tags[i].name === this.tagName) {
          return this.tags[i]
        }
      }
      return null
    }
  },
  created(){
    this.getList()
  }
}
</script>
