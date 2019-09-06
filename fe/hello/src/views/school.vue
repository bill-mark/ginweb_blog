<template>
<div>
	<div class="addschool">
		<input type="newname" v-model="newname" placeholder="学校名">
		<input type="newarea" v-model="newarea" placeholder="地区" >
	
		<button type="submit" @click="submit">增加</button>
	</div>

	<div class="top">
		<div>学校名</div>
		<div>地区</div>
	</div>
	<div class="school" v-for="item in list">
		<div>{{item.name}}</div>
		<div>{{item.area}}</div>
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
			total:4,
                  list_total:0,//总共多少条
		}
	},
	mounted(){
       //this.get_alllist()
       this.get_list()
	},
	methods:{
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
	width: 500px;
	justify-content: space-around;
	margin:0 auto;
}
.addschool{
	margin-top: 20px;
}
</style>






