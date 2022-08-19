<template>
  <div class="container">
    <h1 class="text-center">项目</h1>
    <hr />
    <div class="grid-3_xs-1_sm-2_md-2" v-if="projects.length">
      <div
        v-for="(project, index) in projects"
        :key="index"
        class="col"
      >
        <a :href="project.url" target="_blank">
          <vs-card
            class="post-card"
          >
            <template #title>
              <h1>{{ project.name }}</h1>
            </template>
            <template #text>
              <p>{{ project.desc }}</p>
            </template>
            <template #img>
              <!-- <img :src="project.img" /> -->
              <img :src=project.img>
            </template>
            <template #interactions>
              <vs-tooltip right>
                <vs-button size="large" dark icon :href="project.url" blank>
                  <i class="bx bxl-github"></i>
                </vs-button>
                <template #tooltip>
                  GitHub
                </template>
              </vs-tooltip>
            </template>
          </vs-card>
        </a>
      </div>
    </div>
    <Nothing v-else />
  </div>
</template>

<script>
import Nothing from '@/components/Nothing.vue'

export default {
  name: 'ViewProjects',
  components: {
    Nothing
  },
  methods:{
    async getList(){
        const { data: res } = await this.$http.get('project' )
        this.projects=res.data
    }
  },
  data: function () {
    return {
      projects: []
    }
  },
  mounted: function () {
    this.changeTitle('项目')
  },
  created(){
    this.getList()
  }
}
</script>

<style scoped>
  img{
    width: 100%;  
  }
  .shadow{
    border: 1px solid;
  }
</style>