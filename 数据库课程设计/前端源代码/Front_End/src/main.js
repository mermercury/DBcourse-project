import Vue from "vue";
import ElementUI from "element-ui";
import VueClipboard from "vue-clipboard2";
import ECharts from "vue-echarts";
import "element-ui/lib/theme-chalk/index.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Config from "@/common/config";
import { vueInstance } from "./common/ajax";

import axios from "axios";
import "./common/element-variables.scss";


import 'echarts'

/* 引入 axios 并挂载到 Vue 实例上 */
Vue.prototype.$axios = axios;

/* 指定 axios 发送请求的目标后端地址的根路径，一般为后端服务器IP+端口，若有部署域名则可以是域名地址 */
axios.defaults.baseURL = Config.backEndUrl;

Vue.config.productionTip = false;

Vue.use(ElementUI);
Vue.use(VueClipboard);

Vue.component("v-chart", ECharts);

let instance = new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");

vueInstance.instance = instance;
vueInstance.store = instance.$store;

//引入表格
// import echarts from 'echarts'
// Vue.prototype.$echarts = echarts
