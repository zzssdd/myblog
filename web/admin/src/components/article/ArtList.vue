<template>
  <div>
    <el-row>
      <el-col :span="10">
      <el-input style="width:200px" size="samll" v-model="key"></el-input>
      <el-button type="primary" @click="search">搜索</el-button>
      </el-col>
      <el-col :span="5">
        <el-button type="primary" @click="$router.push('/addart')">新增</el-button>
      </el-col>
      <el-col :span="8">
              <el-select v-model="Category" placeholder="请选择分类" @change="cateShow">
              <el-option
                v-for="item in catelist"
                :key="item.id"
                :label="item.name"
                :value="item.name"
              >
              </el-option>
            </el-select>
          <el-button type="primary" @click="showAll">显示全部</el-button>
      </el-col>
    </el-row>
  
    <el-table
    :data="artList"
    border
    style="width: 100%">
    <el-table-column label="ID" prop="ID" align="center"  width="100"></el-table-column>
    <el-table-column
      label="更新日期"
      align="center" 
      width="220">
      <template slot-scope="scope">
        <span>{{scope.row.UpdatedAt|DateForm}}</span>
      </template>
    </el-table-column>
    <el-table-column
      prop="Category.name"
      label="分类"
      align="center" 
      width="180">
    </el-table-column>
    <el-table-column
      prop="title"
      label="文章标题"
      align="center" 
      width="200">
    </el-table-column>
    <el-table-column
      prop="desc"
      label="文章描述"
      align="center" 
      width="300">
    </el-table-column>
    <el-table-column
      prop="img"
      label="封面"
      align="center" 
      width="200">
      <template slot-scope="scope">
      <el-image :src="scope.row.img" lazy></el-image>  
      </template>
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      align="center" 
      width="200">
      <template slot-scope="scope">
        <el-button @click="$router.push(`/addart/${scope.row.ID}`)" type="primary" plain size="small">编辑</el-button>
        <el-button @click="deleteClick(scope.row)" type="danger"  plain size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  </div>
  
</template>

<script>
import day from 'dayjs'
import router from '@/router'

export default {
    name:'ArtList',
    data(){
      return {
        Category:undefined,
        catelist:[],
         artList:[],
         key:'',
      }
    },
    filters:{
      DateForm:function(date){
        return date?day(date).format('YYYY年MM月DD日 HH:mm'):'暂无'
      }
    },
    methods:{
      async getArtList(){
        const {data:res}=await this.$http.get('article/list')
        if(res.status!==200){
          console.log(res);
          if(res.status===1004||1005||1006||1007){
            window.sessionStorage.clear()
            this.$router.push('/login')
          }
          this.$message.error(res.$message)
        }
        this.artList=res.data
      },
      async getCategoryList() {
      const { data: res } = await this.$http.get("category");
      if (res.status !== 200) {
        console.log(res);
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear();
          this.$router.push("/login");
          console.log(res)
        }
        this.$message.error(res.$message);
      }
      this.catelist = res.data;
    },
      handleClick(row){
        router.push(`/addart/${row.ID}`)
      },
      deleteClick(row){
       this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`article/delete/${row.ID}`)
          console.log(res);
          if (res.status != 200) return this.$message.error(res.message)
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          this.getArtList()
          }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });          
        });
      },
      async cateShow(){
        const {data:res}=await this.$http.get(`articlecate/${this.Category}`)
        if (res.status!==200) {
          console.log(res);
          if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
        }
        this.artList = res.data
      },
      async search(){
          const {data:res}=await this.$http.get(`article/search`,{
          params:{
            title:this.key
          }
        })
        if (res.status!==200) {
          console.log(res);
          if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
        }
        this.artList = res.data
      },
      showAll(){
        this.getArtList()
      }
    },
    created(){
      this.getArtList()
      this.getCategoryList()
    },
}
</script>

<style scoped>
  .el-image{
    width: 200px;
    height: 100px;
  }
  .el-table{
    margin-top: 20px;
  }
</style>