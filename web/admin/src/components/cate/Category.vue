<template>
  <div>
    <div class="btn">
      <el-button type="primary" @click="addCategory">新增分类</el-button>
    </div>
    <el-table :data="CategoryList" border style="width: 100%">
      <el-table-column label="ID" prop="id" align="center"  width="100"></el-table-column>
      <el-table-column prop="name" label="分类名" align="center"  width="300">
        <template slot-scope="scope">
            <span v-show="scope.row.id!==editId">{{scope.row.name}}</span>
            <input type="text" v-show="scope.row.id===editId" v-model="scope.row.name" @blur="FinEdit(scope.row)">
        </template>
      </el-table-column>
      <el-table-column fixed="right" align="center"  label="操作" >
        <template slot-scope="scope">
          <el-button
            @click="Edit(scope.row)"
            type="primary"
            plain
            size="small"
            >编辑</el-button
          >
          <el-button
            @click="deleteClick(scope.row)"
            type="danger"
            plain
            size="small"
            >删除</el-button
          >
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
      :total="CommentList.length">
    </el-pagination>
    </div> -->
  </div>
</template>

<script>

export default {
  name: "Category",
  data() {
    return {
      CategoryList: [],
      editId:-1,
      InfoName:''
      //  currentPage: 1, // 当前页码
      //  total: 20, // 总条数
      //  pageSize: 2 // 每页的数据条数
    };
  },
  methods: {
    async getCategoryList() {
      const { data: res } = await this.$http.get("category");
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          console.log(res);
          window.sessionStorage.clear();
          this.$router.push("/login");
        }
        this.$message.error(res.$message);
      }
      this.CategoryList = res.data;
    },
    deleteClick(row) {
      this.$confirm("此操作将永久删除该分类, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(async() => {
            const { data: res } = await this.$http.delete(`category/delete/${row.id}`);
            if (res.status != 200) {
              console.log(res);
              return this.$message.error(res.message);
              }
            this.$message({
              type: "success",
              message: "删除成功!",
            });
            this.getCategoryList();
        }).catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    Edit(row){
      this.editId=row.id
      this.InfoName=row.name
    },
    async FinEdit(row){
      this.editId=-1
      if(row.name==''){
      row.name=this.InfoName
      return this.$message.error('分类名不能为空')
      }
      const {data:res}=await this.$http.put(`category/${row.id}`,{
        name:row.name,
      })
      if(res.status!=200) {
        row.name=this.InfoName
        return this.$message.error(res.message)
      }
      this.$message({
        type:"success",
        message:"更新分类成功"
      })
    },
    addCategory(){
     this.$prompt('请输入新增分类名', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
        }).then(async ({ value }) => {
          if(value=='') return this.$message.error('分类名不为空')
          const { data: res } = await this.$http.post('category/add', {
          name: value,
          })
          if (res.status != 200) return this.$message.error(res.message)
          this.$message({
            type: 'success',
            message: '新增分类成功'
          });
          this.getCategoryList()
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消输入'
          });       
        });
    }
    },
  created() {
    this.getCategoryList();
  },
};
</script>

<style scoped>
.el-image {
  width: 200px;
  height: 100px;
}
.btn{
  margin-bottom: 10px;
}
</style>
