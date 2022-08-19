<template>
  <el-card class="box-card">
    <el-form
      ref="artForm"
      :model="artInfo"
      :inline="true"
      label-position="top"
      :rules="artInfoRules"
    >
      <el-row>
        <el-col :span="14">
          <el-form-item label="文章标题：">
            <el-input
              size="big"
              v-model="artInfo.title"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="14">
          <el-form-item label="文章描述：">
            <el-input
              type="textarea"
              :rows="8"
              style="width: 400px"
              v-model="artInfo.desc"
            >
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="文章封面" v-show="!artInfo.img" :on-success="handleAvatarSuccess">
            <el-upload
              class="upload-demo"
              drag
              action="#"
              multiple
            >
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">
                将文件拖到此处，或<em>点击上传</em>
              </div>
            </el-upload>
          </el-form-item>
          <el-image :src="artInfo.img" v-show="artInfo.img"></el-image>
        </el-col>
      </el-row>
      <el-form-item label="文章内容：" style="width:100%" prop="content">
          <Edit v-model="artInfo.content"></Edit>
      </el-form-item>
    </el-form>
     <div class="btn">
        <el-button type="primary" @click="submitForm"
          >完成</el-button
        >
        <el-button type="danger" @click="back"
          >返回</el-button
        >
        </div>
  </el-card>
</template>

<script>
import Edit from '@/components/editor/Edit.vue'
export default {
  name: "AddLife",
  components:{ Edit },
  props:['id'],
  data() {
    return {
      catelist:[],
      artInfo: {
        id: 0,
        title: '',
        desc: '',
        content: '',
        img: '',
      },
       artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'change' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'change' },
        ],
        content: [{ required: true, message: '请输入文章内容', trigger: 'change' }],
      },
    };
  },
  methods: {
    async getArtInfo(id){
      const {data:res}=await this.$http.get(`life/${id}`)
      console.log(res);
      if(res.status!=200){
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$message.error(res.message)
          this.$router.push('/login')
          return
        }
        this.$message.error(res.message)
        return
      }
      this.artInfo=res.data;
      this.artInfo.img=''
    },
    handleAvatarSuccess(res,file){
      this.artInfo.img = URL.createObjectURL(file.raw);
    },
    async submitForm(){
        console.log(this.artInfo.content);
          this.$refs.artForm.validate(async (valid) => {
          if(valid){
          if(this.id){
          const { data: res } = await this.$http.put(`life/edit/${this.id}`, this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$router.push('/lifelist')
          this.$message({
            type:'success',
            message:'更新文章成功！'
          })
        }else{
          const { data: res } = await this.$http.post('life/add', this.artInfo)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$router.push('/lifelist')
          this.$message({
            type:'success',
            message:'添加文章成功！'
          })
        };
        }else{
          return this.$message.error('输入数据非法，请重新输入')
        }})
    },
    back(){
      this.$router.push('/lifelist')
    }
  },
  created(){
    if(this.id){
      this.getArtInfo(this.id)
    }
  }
};
</script>

<style scoped>
</style>
