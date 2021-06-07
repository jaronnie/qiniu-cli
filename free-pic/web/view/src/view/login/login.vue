<template>
  <div class="login-layout">
    <el-form status-icon  :rules="rules"  ref="ruleForm" :model="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item
        label="用户名" prop="username" style="width: 400px;">
        <el-input autocomplete="off" v-model="ruleForm.username"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" style="width: 400px;">
        <el-input type="password" autocomplete="off" v-model="ruleForm.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
        <el-button type="primary" @click="trueLogin">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'login',
  data() {
    return {
      ruleForm:{
        username: '',
        password: '',
      },
      rules:{
        username:[
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {required: true, pattern: '[a-zA-Z-]', message: "只支持英文和-", trigger: 'blur'}
        ],
        password:[
          {required: true, message: '请输入密码', trigger: 'blur'},
          {min: 6, max: 16, message: '长度在6到16位',trigger: 'blur'}
        ]
      },
    }
  },
  methods:{
    register() {
      let param = new URLSearchParams()
      param.append('username', this.ruleForm.username)
      param.append('password', this.ruleForm.password)
      this.$axios ({
        method:'post',
        url:"http://localhost:8081/api/v3/login/add",
        data:param
      }).then(function(response){
        if (response.data.message==="success add user"){
          alert("注册成功")
        } else {
          alert("该用户已经存在")
        }
      })
    },
    trueLogin() {
      let me = this
      let param = new URLSearchParams()
      param.append('username', this.ruleForm.username)
      param.append('password', this.ruleForm.password)
      this.$axios ({
        method:'post',
        url:"http://localhost:8081/api/v3/trueLogin",
        data:param
      }).then(function(response){
        console.log(response.data)
        if (response.data.message==="login successfully"){
          me.$router.push('/Dashboard')
          let user_token = response.data.data;
          // 存token到localStorage中
          localStorage.setItem("currentUser_token", user_token);
        } else {
          alert("用户名/密码错误")
        }
      })
    }
  }
}
</script>

<style scoped>
@import '../../style/login.css';
</style>
