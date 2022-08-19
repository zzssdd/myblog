<template>
    <el-row class="tac">
    <el-menu
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router
      >
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-s-fold" ref="Icon"></i>
          <span>文章管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="1-1" route="addart">
         <i class="el-icon-edit"></i>
          <span>写文章</span>
            </el-menu-item>
          <el-menu-item index="1-2" route="artlist">
            <i class="el-icon-notebook-1"></i>
            <span>文章列表</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>

       <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-s-fold" ref="Icon"></i>
          <span>动态管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="2-1" route="addlife">
            <i class="el-icon-edit-outline"></i>
            <span>写动态</span>
          </el-menu-item>
            <el-menu-item index="2-2" route="lifelist">
            <i class="el-icon-notebook-2"></i>
            <span>动态列表</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>

       <el-submenu index="3">
        <template slot="title">
          <i class="el-icon-s-fold" ref="Icon"></i>
          <span>项目管理</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="3-1" route="addpro">
         <i class="el-icon-edit"></i>
          <span>新增项目</span>
            </el-menu-item>
          <el-menu-item index="3-2" route="prolist">
            <i class="el-icon-notebook-1"></i>
            <span>项目列表</span>
          </el-menu-item>
        </el-menu-item-group>
      </el-submenu>

      <el-menu-item index="3" route="comment">
        <i class="el-icon-chat-dot-square"></i>
        <span slot="title">文章评论管理</span>
      </el-menu-item>
        <el-menu-item index="4" route="lifecomment">
        <i class="el-icon-chat-line-round"></i>
        <span slot="title">生活评论管理</span>
      </el-menu-item>
           <el-menu-item index="5" route="board">
        <i class="el-icon-chat-round"></i>
        <span slot="title">留言板管理</span>
      </el-menu-item>
      <el-menu-item index="6" route="category">
        <i class="el-icon-tickets"></i>
        <span slot="title">分类列表</span>
      </el-menu-item>


    </el-menu>
</el-row>
</template>

<script>
export default {
    name:'Nav',
    methods:{
    handleOpen(){
        this.$refs.Icon.className='el-icon-s-unfold'
     },
     handleClose(){
         this.$refs.Icon.className='el-icon-s-fold'
     }
    }
}
</script>

<style>

</style>