<template>
  <div class="app">
    <div class="top">
       <router-link to="/login">登录</router-link> |
       <router-link to="/register">注册</router-link>   |
        <router-link to="/update_password">修改密码</router-link>  |
        <router-link to="/">主页</router-link>
    </div>
    <div class="down">
        <router-view/>
    </div>
    
  </div>
</template>

<style lang="scss" scoped>
.app{
  
  .top{
    
  }
  .down{
      margin-top: 50px;
      //width: 500px;
      text-align: center;
  }
}

</style>
