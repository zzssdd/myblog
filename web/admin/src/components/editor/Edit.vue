<template>
  <div>
    <vue-tinymce
      :setting="setting"  v-model="content" />
  </div>
</template>

<script>
export default {
  name: 'Edit',
   props: {
    value: {
      type: String,
      default: '',
    },
  },
  data(){
    return {
      setting: {
      language: 'zh_CN',
        height: '600px',
        plugins: 'preview paste wordcount code imagetools image media codesample lists table',
        branding: false,
        paste_data_images: true,
        toolbar: [
          'undo redo | styleselect |fontsizeselect| bold italic underline strikethrough |alignleft aligncenter alignright alignjustify |blockquote removeformat |numlist bullist table',
          'preview paste code codesample |image media',
        ],
        //上传图片
        images_upload_handler: async (blobInfo, succFun, failFun) => {
          let formdata = new FormData()
          formdata.append('file', blobInfo.blob(), blobInfo.name())
          const { data: res } = await this.$http.post('upload', formdata)
          succFun(res.url)
          failFun(this.$message.error('上传图片失败'))
        },
        imagetools_cors_hosts: ['*'],
        imagetools_proxy: '',
      },
      content: this.value,
    }
  },
  watch: {
    value(newV) {
      this.content = newV
    },
    content(newV) {
      this.$emit('input', newV)
    },
  },
}
</script>

<style scoped>
@import url('tinymce/skins/ui/oxide/skin.min.css');
@import url('tinymce/skins/content/default/content.min.css');
</style>
