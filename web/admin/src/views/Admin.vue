<template>
<el-container class="page">
  <el-header height="50px">
      <span>Yogen的博客后台管理</span>
      <el-button id="exit" type="danger" round @click='exist' size="small">退出</el-button>
  </el-header>
  <el-container>
    <el-aside width="200px">
    <Nav />
    </el-aside>
    <el-main>
        <router-view replace :key="$router.path"></router-view>
    </el-main>
  </el-container>
</el-container>
</template>

<script>
import Nav from '@/components/Nav.vue';
export default {
  components: { Nav },
 name:'Admin',
 methods:{
     exist(){
        window.sessionStorage.clear();
        this.$router.push("/login");
     }
 }
}
</script>

<style scoped>
    .page{
        position: absolute;
        width: 100%;
        height: 100%;
    }
    .el-header,.el-aside{
        background-color: rgba(3, 12, 54,1);
    }
    .el-aside{
        height: 567.24px;
    }
    .el-header{
        text-align: center;
        line-height: 50px;
        color: #fff;
    }
    #exit{
        margin-top: 10px;
        float: right;
    } 
</style>