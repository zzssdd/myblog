<template>
  <div :class="`${home ? 'avatar-home' : ''}`">
    <vs-avatar :size="size" circle>
      <img
        src="@/assets/avater.jpg"
        @click="playmusic()"
      >
    </vs-avatar>
  </div>
</template>

<script>
export default {
  name: 'Avatar',
  data: function () {
    return {
      bgaudiocnt:0
    }
  },
  props: [
    'size',
    'home'
  ],
  methods:{
    playmusic(){
      this.bgaudiocnt^=1
      if (this.bgaudiocnt==1) {
        document.getElementById('bgaudio').play();
      }else{
        document.getElementById('bgaudio').pause();
      }
    }
  }
}
</script>

<style scoped>
.avatar-home .vs-avatar, .avatar-home .vs-avatar-content {
  min-width: 200px !important;
  min-height: 200px !important;
}
.vs-avatar img:hover{
  cursor:pointer;
}
</style>
