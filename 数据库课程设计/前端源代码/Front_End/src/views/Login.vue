<template>
  <div class="login-wrap">
    <div class="login-form">
      <div class="form-title">大乌苏课程管理系统</div>
      <el-form
        :model="formData"
        :rules="rules"
        class="form-content"
        label-width="0px"
        ref="form"
      >
        <el-form-item prop="userid">
          <el-input placeholder="学号/工号/用户名" v-model="formData.userid">
            <span slot="prepend"><i class="el-icon-user"></i></span>
          </el-input>
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            @keyup.enter.native="submit()"
            placeholder="密码"
            type="password"
            v-model="formData.password"
          >
            <span slot="prepend"><i class="el-icon-edit"></i></span>
          </el-input>
        </el-form-item>

        <el-form-item prop="usertype">
          <el-radio-group v-model="formData.usertype">
            <el-radio label="1">学生</el-radio>
            <el-radio label="2">教师</el-radio>
            <el-radio label="3">教务管理员</el-radio>
          </el-radio-group>
        </el-form-item>

        <div class="login-btn" v-loading="this.$store.state.loading">
          <el-button @click="submit()" type="primary">登录</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
// import { login } from "../api/user";
import axios from "axios";
import Config from "@/common/config";

export default {
  data: function() {
    return {
      formData: {
        userid: "",
        password: "",
        usertype: "1",
      },
      rules: {
        userid: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
        usertype: [
          { required: true, message: "请选择用户类型", trigger: "blur" },
        ],
      },
    };
  },
  methods: {
    submit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          axios
            .post(Config.backEndUrl + "/auth/login", {
              userid: Number(this.formData.userid),
              password: this.formData.password,
              usertype: Number(this.formData.usertype),
            })
            .then((res) => {
              // console.log(res);
              this.$message.success("登录成功: " + this.formData.userid);
              this.$store.commit("login", res);
              this.$router.push({ name: "container" });
            })
            .catch((error) => {
              console.log(error);
              this.$message.warning("登录失败");
            });
          /////////////////////
          // login(
          //   this.formData.userid,
          //   this.formData.password,
          //   this.formData.usertype
          // ).then((res) => {
          //   console.log(res.data);
          //   this.$message.success("登录成功: " + res.data.userid);
          //   this.$store.commit("login", res);
          //   this.$router.push({ name: "container" });
          // });
        }
      });
    },
  },
};
</script>

<style scoped>
.login-wrap {
  position: relative;
  width: 100%;
  height: 100%;
  background-image: url("../assets/login-background.jpg");
  background-size: 100% 100%;
}

.form-title {
  width: 100%;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  color: #fff;
  border-bottom: 1px solid #ddd;
}

.login-form {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 350px;
  margin: -190px 0 0 -175px;
  border-radius: 5px;
  background: rgba(0, 0, 0, 0.6);
  overflow: hidden;
}

.form-content {
  padding: 30px 30px;
}

.login-btn {
  text-align: center;
}

.login-btn button {
  width: 100%;
  height: 36px;
}

.el-radio {
  color: #fff;
}
</style>
