<template>
  <div class="login-container">
    <el-form ref="loginForm" class="login-form" auto-complete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">基于区块链的高校助学金系统</h3>
      </div>
      <el-form-item label="学号" prop="wyuUserId">
        <el-input v-model="registerdata.wyuUserId" placeholder="请输入学号"></el-input>
      </el-form-item>
      <el-form-item label="姓名" prop="wyuUserName">
        <el-input v-model="registerdata.wyuUserName" placeholder="请输入姓名"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="wyuPasswd">
        <el-input placeholder="请输入密码" v-model="registerdata.wyuPasswd" show-password></el-input>
      </el-form-item>
      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleRegister">注册</el-button>
      <el-link type="primary" style="float:right" @click.native.prevent="returnLogin">返回登录界面</el-link>
    </el-form>
  </div>
</template>

<script>
  import { register } from '@/api/grantUser'

  export default {
    name: 'Login',
    data() {
      return {
        loading: false,
        redirect: undefined,
        registerdata: {
          wyuUserId: "",
          wyuUserName: "",
          wyuPasswd: ""
        },
        value: ''
      }
    },
    methods: {
      handleRegister() {

        if (this.registerdata.wyuUserId) {
          this.loading = true
          register(this.registerdata).then(() => {
            // this.$router.push({ path: '/' })
            this.$router.push({ path: "/login" })
            this.loading = false
            this.$message('注册成功')
          }).catch(() => {
            this.loading = false
          })
        } else {
          this.$message('请输入学号')
        }
      },
      returnLogin() {
        this.loading = true
        this.$router.push({ path: "/login" })
        this.loading = false
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
