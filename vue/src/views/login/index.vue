<template>
  <div class="login-container" id="background">
    <el-form ref="loginForm" class="login-form" auto-complete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">基于区块链的高校助学金系统</h3>
      </div>
      <el-form-item label="账号" prop="wyuUserId">
        <el-input v-model="logindata.wyuUserId" placeholder="请输入学号"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="wyuPasswd">
        <el-input placeholder="请输入密码" v-model="logindata.wyuPasswd" show-password></el-input>
      </el-form-item>
      <el-button :loading="loading" type="primary" style="width:45%;margin-bottom:30px;" @click.native.prevent="handleLogin">登录</el-button>
      <el-button :loading="loading" type="primary" style="width:45%;float:right;margin-bottom:30px;" @click.native.prevent="handleRegister">注册</el-button>
      <div class="tips">
        <span style="margin-right:20px;">若还没有注册，请先注册账号</span>
      </div>
      <el-link type="primary" style="float:right" @click="relation">联系开发者</el-link>

    </el-form>
    <el-form ref="myLoginForm">
    </el-form>
  </div>
</template>

<script>
  export default {
    name: 'Login',
    data() {
      return {
        loading: false,
        logindata: {
          wyuUserId: "",
          wyuPasswd: ""
        },
        registerdata: {
          wyuUserId: "",
          wyuUserName: "",
          wyuPasswd: ""
        },
        value: ''
      }
    },
    methods: {
      handleLogin() {
        if (this.logindata.wyuUserId) {
          this.loading = true
          this.$store.dispatch('grantUser/login', this.logindata).then(() => {
            this.$router.push({ path: '/' })
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          this.$message('请输入学号')
        }
      },
      handleRegister() {
        this.loading = true
        this.$router.push({ path: "/register" })
        this.loading = false
      },
      relation() {
        this.$alert('email:zhipingZeng_43@163.com', '开发者联系方式', {
          confirmButtonText: '确定',
        });
      }
    }
  }
</script>

<style lang="scss" scoped>
  $bg:#2d3a4b;
  $dark_gray:#889aa4;
  $light_gray:#eee;

  .login-container {
    min-height: 100%;
    width: 100%;
    background-color: $bg;
    overflow: hidden;

    .login-form {
      position: relative;
      width: 520px;
      max-width: 100%;
      padding: 160px 35px 0;
      margin: 0 auto;
      overflow: hidden;
    }

    .login-select {
      padding: 20px 0px 30px 0px;
      min-height: 100%;
      width: 100%;
      background-color: $bg;
      overflow: hidden;
      text-align: center;
    }

    .tips {
      font-size: 14px;
      color: #fff;
      margin-bottom: 10px;

      span {
        &:first-of-type {
          margin-right: 16px;
        }
      }
    }

    .svg-container {
      padding: 6px 5px 6px 15px;
      color: $dark_gray;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }

    .title-container {
      position: relative;

      .title {
        font-size: 26px;
        color: $light_gray;
        margin: 0px auto 40px auto;
        text-align: center;
        font-weight: bold;
      }
    }

    .show-pwd {
      position: absolute;
      right: 10px;
      top: 7px;
      font-size: 16px;
      color: $dark_gray;
      cursor: pointer;
      user-select: none;
    }
  }
</style>
