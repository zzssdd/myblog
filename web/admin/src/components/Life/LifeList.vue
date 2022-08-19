<template>
  <div>
    <el-row>
      <el-col :span="5">
        <el-button type="primary" @click="$router.push('/addlife')">新增</el-button>
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
      prop="title"
      label="文章标题"
      align="center" 
      width="200">
    </el-table-column>
    <el-table-column
      prop="desc"
      label="文章描述"
      align="center" 
      width="400">
    </el-table-column>
    <el-table-column
      prop="img"
      label="封面"
      align="center" 
      width="300">
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
        <el-button @click="$router.push(`/addlife/${scope.row.ID}`)" type="primary" plain size="small">编辑</el-button>
        <el-button @click="deleteClick(scope.row)" type="danger"  plain size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
    <!-- <div class="block">
      <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[1, 5, 10, 20]"
      :page-size=pageSize
      layout="total, sizes, prev, pager, next, jumper"
      :total="artList.length">
    </el-pagination>
    </div> -->
  </div>
  
</template>

<script>
import day from 'dayjs'
import router from '@/router'

export default {
    name:'LifeList',
    data(){
      return {
        Category:undefined,
        catelist:[],
         artList:[],
         key:'',
        //  currentPage: 1, // 当前页码
        //  total: 20, // 总条数
        //  pageSize: 2 // 每页的数据条数
      }
    },
    filters:{
      DateForm:function(date){
        return date?day(date).format('YYYY年MM月DD日 HH:mm'):'暂无'
      }
    },
    methods:{
      async getArtList(){
        const {data:res}=await this.$http.get('life')
        if(res.status!==200){
          if(res.status===1004||1005||1006||1007){
            window.sessionStorage.clear()
            this.$router.push('/login')
          }
          this.$message.error(res.$message)
        }
        this.artList=res.data
      },
      handleClick(row){
        router.push(`/addlife/${row.ID}`)
      },
      deleteClick(row){
       this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`life/delete/${row.ID}`)
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
    },
    created(){
      this.getArtList()
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