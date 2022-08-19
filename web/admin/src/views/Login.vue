<template>
  <div class="container">
    <el-form
      :model="Form"
      status-icon
      :rules="rules"
      ref="ruleForm"
      class="demo-ruleForm"
    >
      <el-form-item label="用户名" prop="username">
        <el-input v-model="Form.username" suffix-icon="el-icon-user"></el-input>
      </el-form-item>

      <el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input
          type="password"
          v-model="Form.password"
          autocomplete="off"
          @keyup.enter.native="submitForm"
          suffix-icon="el-icon-info"
          >
        ></el-input>
      </el-form-item>

        <div class="btn">
        <el-button type="primary" @click="submitForm"
          >登录</el-button
        >
        <el-button @click="resetForm()">重置</el-button>
        </div>

      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  username: "Login",
  data(){
      return{
          Form:{
              username:'',
              password:''
          },
          rules:{
              username:[
                  {required:true,message:'用户名不能为空',trigger:'blur'},
                  {min:3,max:20,message:'用户名长度在3到20之间',trigger:'blur'}
              ],
              password:[
                  {required:true,message:'密码不能为空',trigger:'blur'},
                  {min:6,message:'密码长度必须大于6',trigger:'blur'}
              ]
          }
      }
  },
  methods:{
      submitForm(){
          this.$refs.ruleForm.validate(async(valid)=>{
              if(valid){
                  const {data:res}=await this.$http.post('login',this.Form)
                  console.log(res);
                  if(res.status!==200){
                      return this.$message.error('用户不存在，请重新输入');
                  }else{
                      window.sessionStorage.setItem('Username',this.Form.username)
                      window.sessionStorage.setItem('token',res.token)
                      this.$router.push('/admin')
                      return this.$message({
                        message:'登录成功！',
                        type:'success'
                      })
                  }
              }else{
                  return this.$message.error('输入数据非法，请重新输入')
              }
          })
      },
      resetForm(){
          this.$refs.ruleForm.resetFields();
      }
  },
};
</script>

<style scoped>
.container {
  position: absolute;
  height: 100%;
  width: 100%;
  z-index: -1;
  background: url("@/assets/login_bg.jpg") no-repeat;
  background-size: cover;
}
.el-form{
  width: 400px;
  height: 250px;
  background-color: rgba(1, 7, 28,0.6);
  position: absolute;
  top: 60%;
  left: 80%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.el-form-item{
    margin-left: 20px;
    margin-top: 25px;
    width: 380px;
}
.el-input{
    width: 300px;
}
.btn{
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
