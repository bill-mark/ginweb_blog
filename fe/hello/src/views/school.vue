<template>
<div>
	<div class="addschool">
		<input  v-model="newname" placeholder="学校名">
		<input  v-model="newarea" placeholder="地区" >
		<button type="submit" @click="submit">增加</button>
	</div>

      <div>
            <input type="txt" v-model="likekey" >
            <button @click="like_search">搜索</button>
      </div>
      <div>
            <input type="txt" v-model="area" >
            <button @click="area_search">地区筛选</button>
      </div>

	<div class="top">
		<div style="width: 100px;">学校名</div>
		<div style="width: 50px;text-align: left;">地区</div>
            <div>操作</div>
	</div>
	<div class="school" v-for="item in list">
		<div style="width: 100px;">{{item.name}}</div>
		<div style="width: 50px;">{{item.area}}</div>
            <div @click="delet(item.id)" style="cursor: pointer;">删除</div>
	</div>

	 <div class="pages"  style="margin-top:30px;margin-left:50px;padding-bottom:40px;display:flex;justify-content:center;user-select:none;">
            <el-pagination 
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="page"
              :page-sizes="[5, 10]"
              :page-size=total
              layout="total, sizes, prev, pager, next,slot,jumper"
              :total="list_total">
            </el-pagination>
       </div>
</div>
</template>

<script >
export default{
	data(){
		return{
			list:[],
			newname:'',
			newarea:'',
			page:1,
			total:5,
                  list_total:0,//总共多少条
                  likekey:'',
                  area:'',
		}
	},
	mounted(){
       //this.get_alllist()
       this.get_list()
	},
	methods:{
       like_search(){
           this.$http.get('/school/likesearch',{
               params:{
                   key:this.likekey
               }                                     
           })
           .then((res)=>{        
               if(res.data.code == 0){
                  this.list = res.data.data
               }
               if(res.data.code == -100){
                  this.$router.push({path:'/login'})
               }
           })
           .catch( (error)=>{
                  
           })
       },
       area_search(){
           this.$http.get('/school/areasearch',{
               params:{
                   key:this.area
               }                                     
           })
           .then((res)=>{        
               if(res.data.code == 0){
                  this.list = res.data.data
               }
               if(res.data.code == -100){
                  this.$router.push({path:'/login'})
               }
           })
           .catch( (error)=>{
                  
           })
       },
       delet(id){
           this.$http.post('/school/delet',{
                     id:id                                   
                  })
                  .then((res)=>{        
                      if(res.data.code == 0){
                        this.get_list()
                      }
                      if(res.data.code == -100){
                        this.$router.push({path:'/login'})
                      }
                  })
                  .catch( (error)=>{
                         
                  })
       },
       handleSizeChange(val){  //修改每页展示条数
           this.total = val;
           this.get_list();
       },
       handleCurrentChange(val){ //分页回调
           this.page = val;
           this.get_list();
       },
       get_alllist(){
       	  	this.$http.get('/school/list',{
       	  	    params:{

       	  	    }	         	  	      	     
       	  	})
       	  	.then((res)=>{        
       	  	    if(res.data.code == 0){
       	  	    	this.list = res.data.data
       	  	    }
       	  	    if(res.data.code == -100){
       	  	    	this.$router.push({path:'/login'})
       	  	    }
       	  	})
       	  	.catch( (error)=>{
       	  	       
       	  	})
       },
       submit(){
            this.$http.post('/school/add',{	     
      	       name:this.newname,
      	       area:this.newarea,  	     
      	})
      	.then((res)=>{        
      	    if(res.data.code == 0){
      	    	this.newname = ''
      	    	this.newarea = ''
      	    	this.get_list()
      	    }
      	})
      	.catch( (error)=>{
      	       
      	})
       },
       get_list(){
       	  	this.$http.get('/school/get',{
       	  	    params:{
                          page:this.page,
                          total:this.total
       	  	    }	         	  	      	     
       	  	})
       	  	.then((res)=>{        
       	  	    if(res.data.code == 0){
       	  	    	this.list_total = res.data.allnumber
                        this.list = res.data.data
       	  	    }
       	  	    if(res.data.code == -100){
       	  	    	this.$router.push({path:'/login'})
       	  	    }
       	  	})
       	  	.catch( (error)=>{
       	  	       
       	  	})
       },
      }
}
</script>
<style lang="scss" scoped>
.top{
	margin-top: 30px;
	line-height: 30px;
	display: flex;
	width: 500px;
	height: 30px;
	justify-content: space-around;
	margin:0 auto;
}
.school{
	margin-top: 5px;
	display: flex;
      height: 25px;
      line-height: 25px;
	width: 500px;
	justify-content: space-around;
	margin:0 auto;
}
.addschool{
	margin-top: 20px;
}
</style>






