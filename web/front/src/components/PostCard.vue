<template>
  <router-link :to="flag===1?`/article/${post.ID}`:`/life/${post.ID}`">
    <vs-card class="post-card">
      <template #title>
        <h1>{{ post.title }}</h1>
        <small>{{post.CreatedAt|formatDate}}</small>
        <br/>
        <small>{{ post.desc }}</small>
      </template>
      <template #img>
        <img :src="post.img" alt />
      </template>
      <template #text>
        <p class="post-txt">{{ post.des }}</p>
        <small class="post-card-tag">
          <b  style="margin-right: 5px">
            阅读数：{{ post.readCount }}
          </b>
        </small>
      </template>
      <template #interactions>
        <vs-tooltip right shadow interactivity>
          <Avatar />
        </vs-tooltip>
      </template>
    </vs-card>
  </router-link>
</template>

<script>
import Avatar from '@/components/Avatar.vue'

export default {
  name: 'PostCard',
  props: [
    'post',
    'flag'
  ],
  filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		},
  components: {
    Avatar
  }
}
</script>

<style>
.post-txt {
  overflow: hidden;
  text-overflow: ellipsis;
  display: auto;
}

.post-card .vs-card {
  height: 380px !important;
}

.post-card-tag {
  position: absolute;
  bottom: 20px;
}
</style>
