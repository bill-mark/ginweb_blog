<template>
	
	<div>
		<input type="text" v-model="username" placeholder="Username">
		<input type="password" v-model="password" placeholder="Password" >
		<br>
		<button type="submit" @click="submit">注册</button>
	</div>
</template>

<script>
export default{
	data(){
		return {
          username:'',
          password:'',
          repassword:'',
		}
	},
	methods:{
      submit(){
      	this.$http.post('/register',{	     
      	       username:this.username,
      	       password:this.password,  	     
      	})
      	.then((res)=>{        
      	    if(res.data.code == 0){
      	    	alert(res.data.message)
      	    }
      	})
      	.catch( (error)=>{
      	       
      	})
      }
	}
}
</script>

<style lang="scss">
	
</style>