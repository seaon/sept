<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="header">
          <span class="title">front</span>
        </div>
      </div>
      <div class="main">
        <el-form :model="loginForm" :rules="rules" ref="loginForm">
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input placeholder="请输入密码" v-model="loginForm.password" type="password">
              <i :class="'el-input__icon el-icon-lock'" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%">
              登 录
            </el-button>
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="copyright">
          Copyright &copy; {{ curYear }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
    return {
      curYear: 0,
      loginForm: {
        username: "admin",
        password: "123456",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }],
      },
    };
  },
  created() {
    this.curYear = new Date().getFullYear();
  },
  methods: {
    submitForm() {
      return false;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.login-register-box {
  height: 100vh;
  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top: 5vh;
    .logo {
      height: 35vh;
      width: 35vh;
    }
  }
}

.vPic {
  width: 33%;
  height: 38px;
  float: right !important;
  background: #ccc;
  img {
    cursor: pointer;
    vertical-align: middle;
  }
}

.logo_login {
  width: 100px;
}

#userLayout.user-layout-wrapper {
  height: 100%;
  display: flex;
  width: 100%;
  background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
  justify-content: center; /* 水平居中 */
  align-items: center;     /* 垂直居中 */
  &.mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }

  .container {

    a {
      text-decoration: none;
    }

    .top {
      text-align: center;
      .header {
        height: 44px;
        line-height: 44px;
        margin-bottom: 30px;
        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }

        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }

        .title {
          font-size: 33px;
          color: rgba(0, 0, 0, 0.85);
          font-family: Avenir, "Helvetica Neue", Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
      }
    }

    .main {
      min-width: 260px;
      width: 368px;
      margin: 0 auto;
    }
  }
}
</style>
