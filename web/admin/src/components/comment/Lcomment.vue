<template>
  <div>
    <el-table :data="CommentList" border style="width: 100%">
      <el-table-column label="ID" align="center"  prop="ID" width="100"></el-table-column>
      <el-table-column label="评论时间" align="center"  width="200">
        <template slot-scope="scope" align="center" >
          <span>{{ scope.row.CreatedAt | DateForm }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="Life.title" align="center" label="评论动态" width="250">
      </el-table-column>
      <el-table-column prop="name" align="center"  label="评论者" width="200">
      </el-table-column>
      <el-table-column prop="content" align="center"  label="评论内容" width="500">
      </el-table-column>
      <el-table-column fixed="right" align="center"  label="操作" width="200">
        <template slot-scope="scope">
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
  </div>
</template>

<script>
import day from "dayjs";

export default {
  name: "Lcomment",
  data() {
    return {
      CommentList: [],
    };
  },
  filters: {
    DateForm: function (date) {
      return date ? day(date).format("YYYY年MM月DD日 HH:mm") : "暂无";
    },
  },
  methods: {
    async getCommentList() {
      const { data: res } = await this.$http.get("lifecomment/list");
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          console.log(res);
          window.sessionStorage.clear();
          this.$router.push("/login");
        }
        this.$message.error(res.$message);
      }
      this.CommentList = res.data;
    },
    deleteClick(row) {
      this.$confirm("此操作将永久删除该记录, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async() => {
            const { data: res } = await this.$http.delete(
              `lifecomment/delete/${row.ID}`
            );
            if (res.status != 200) return this.$message.error(res.message);
            this.$message({
              type: "success",
              message: "删除成功!",
            });
            this.getCommentList();
          })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    }
  },
  created() {
    this.getCommentList();
  },
};
</script>

<style scoped>
.el-image {
  width: 200px;
  height: 100px;
}
</style>
