<template>
  
  <div>

    <input type="text" v-model="username" placeholder="Username">
    <input type="password" v-model="password" placeholder="Password" >
    <br>
    <button type="submit" @click="submit">登录</button>
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
        this.$http.post('/login',{      
               username:this.username,
               password:this.password,         
        })
        .then((res)=>{        
            if(res.data.code == 0){
              alert(res.data.message)
              localStorage.token = res.data.token
              this.$router.push({path:'/'})
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