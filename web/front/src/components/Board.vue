<template>
  <div>
      <vs-input v-model="comment.content" placeholder="说点什么吧" />
      <vs-button
        :active="active == false"
        @click="active=true"
      >
        提交
      </vs-button>


      <div class="center-tags" v-if="commentList.length">
        <!-- <div class="grid-3_xs-1_sm-2_md-2"> -->
        <div v-for="(c, index) in commentList" :key="index" class="col">
        <div class="list">
            <h3>评论内容：{{c.content}}</h3>
            <p>评论时间：{{c.CreatedAt|formatDate}} 评论人：{{ c.name }}</p>
       </div>
        </div>
        <!-- </div> -->
      </div>

      <vs-dialog width="50px" not-center v-model="active">
        <template #header>
          <h4 class="not-margin">
            请输入您的<b>姓名</b>
          </h4>
        </template>
        <div class="con-content">
          <vs-input v-model="comment.name" placeholder="请输入姓名"></vs-input>
        </div>
        <template #footer>
          <div class="con-footer">
            <vs-button @click="send_msg" transparent>
              完成
            </vs-button>
            <vs-button @click="active=false" dark transparent>
              取消
            </vs-button>
          </div>
        </template>
      </vs-dialog>
        
    </div>
</template>

<script>
export default {
    name:"Board",
      data:function(){
            return {
                active:false,
                comment:{
                    content:'',
                    name:'',
                    aid:parseInt(this.id),
                },
                commentList:[]
                }
        },
         filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		},
        props:['flag','id'],
        methods: {
            async send_msg() {
               if(this.comment.name===''||this.comment.content===''){
                  this.active=false
                  return alert('输入信息不能为空')
               }
                this.comment.aid=parseInt(this.id)
                const { data: res } =await this.$http.post(this.flag==1?'comment/add':'lifecomment/add',this.comment)
                console.log(res);
                if(res.status!=200) return alert('添加评论失败！')
                alert('添加评论成功！')
                // console.log(this.flag==1?'comment/add':'lifecomment/add',res,"comment",this.comment);
                this.active=false
                this.comment.content=''
                this.comment.name=''
                this.getList()
            },
            async getList(){
                const { data: res } =await this.$http.get(this.flag==1?`comment/${this.id}`:`lifecomment/${this.id}`)
                this.commentList=res.data
            }
        },
        created(){
            this.getList()
        },
        watch:{
            id:function(newVal,oldVal){
                this.getList()
                this.comment.aid=parseInt(newVal)
            }
        }
        
}
</script>

<style>
    .vs-input{
        width: 90% !important;
        height: 100px !important;
    }
    .vs-button{
        margin-top: 10px !important;
    }
    .con-footer .vs-button{
        display: inline;
        width: 100px;
    }
    .list{
        width:90%;
        height:150px;
        background-color: #fff;
        padding: 5px 5px;
    }
</style>