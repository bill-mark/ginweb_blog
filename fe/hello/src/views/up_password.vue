<template>
	
	<div>
		<input type="text" v-model="username" placeholder="Username">
		<input type="password" v-model="password" placeholder="Password" >
    <input type="password" v-model="newpassword" placeholder="NewPassword" >
		<br>
		<button type="submit" @click="submit">确定修改</button>
	</div>
</template>

<script>
export default{
	data(){
		return {
          username:'',
          password:'',
          newpassword:'',
		}
	},
	methods:{
      submit(){
      	this.$http.post('/update_password',{	     
      	       username:this.username,
      	       password:this.password, 
               newpassword:this.newpassword 	     
      	})
      	.then((res)=>{        
      	    if(res.data.code == 0){
      	    	alert(res.data.message)
      	    }else{
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