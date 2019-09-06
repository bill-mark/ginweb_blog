<template>
  <div class="home">
     <div>主页</div>
     <div  style="margin-top: 20px;cursor:pointer" @click="goschool">学校列表</div>
  </div>
</template>

<script>

export default {
  name: 'home',
  methods:{
  	goschool(){
  		this.$router.push({path:'/school'})
  	}
  }
}
</script>
