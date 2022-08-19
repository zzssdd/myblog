<template>
  <div class="container">
    <h1 class="text-center">分类</h1>
    <hr />

    <div class="grid-3_xs-1_sm-2_md-2" v-if="category.length">
      <div v-for="(tag, index) in category" :key="index" class="col">
        <router-link :to="`/tags/${tag.name}`">
          <vs-card type="3" class="center">
            <template #text>
              <h2>{{tag.name}}</h2>
              <p>文章数：{{tag.num}}</p>
            </template>
         </vs-card>
        </router-link>
      </div>
    </div>
    <Nothing v-else />
  </div>
</template>

<script>
import Nothing from '@/components/Nothing.vue'

export default {
  name: 'ViewTags',
  created(){
    this.getTags()
  },
  mounted: function () {
    this.changeTitle('分类')
  },
  data(){
    return{
      category:[]
    }
  },
  methods:{
    async getTags(){
       const { data: res } = await this.$http.get('category' )
        this.category=res.data
    }
  },
  components: {
    Nothing
  }
}
</script>
