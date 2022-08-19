<template>
  <div>
    <el-table :data="CategoryList" border style="width: 100%">
      <el-table-column label="ID" prop="ID" align="center"  width="100"></el-table-column>
      <el-table-column prop="name" label="项目名" align="center"  width="300">
        <template slot-scope="scope">
            <span >{{scope.row.name}}</span>
        </template>
      </el-table-column>
            <el-table-column prop="name" label="图片" align="center"  width="300">
        <template slot-scope="scope">
             <el-image :src="scope.row.img" lazy></el-image>  
        </template>
      </el-table-column>
            <el-table-column prop="url" label="地址" align="center"  width="300">
        <template slot-scope="scope">
            <span>{{scope.row.url}}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" align="center"  label="操作" >
        <template slot-scope="scope">
          <el-button
            @click="Edit(scope.row.ID)"
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
import router from '@/router'

export default {
  name: "Project",
  data() {
    return {
      CategoryList: [],
      InfoName:''
      //  currentPage: 1, // 当前页码
      //  total: 20, // 总条数
      //  pageSize: 2 // 每页的数据条数
    };
  },
  methods: {
    async getCategoryList() {
      const { data: res } = await this.$http.get("project");
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
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
            const { data: res } = await this.$http.delete(`project/delete/${row.ID}`);
            if (res.status != 200) return this.$message.error(res.message);
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
    Edit(id){
        router.push(`/addpro/${id}`)
    }
    },
  created() {
    this.getCategoryList();
  }
}
</script>

<style scoped>
.el-image {
  width: 200px;
  height: 100px;
}

</style>
