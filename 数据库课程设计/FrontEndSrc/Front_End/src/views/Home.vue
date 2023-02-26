<template>
  <div class="home-wrap">
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-fa fa-id-badge"></i> 首页
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="main-wrap">
      <el-container class="home-container">
        <el-header>
          <span>
            <div style="height:100px;width:100%;">
              <h1 align="center">欢迎使用大乌苏课程管理系统</h1>
            </div>
          </span>

        </el-header>
        <el-main>
          <div style="width:100%;">

            <v-chart :option="option_column" style="height: 400px; width: 800px;"></v-chart>

          </div>
        </el-main>





        <!-- <el-aside class="home-aside" width="65%">
          <el-table :data="tableData" stripe>
            <el-table-column
              align="center"
              label="日期"
              prop="date"
              width="120px"
            ></el-table-column>
            <el-table-column
              align="center"
              label="标题"
              prop="title"
            ></el-table-column>
            <el-table-column align="center" label="操作" width="120px">
              <template slot-scope="scope">
                <el-button
                  @click="openNews(scope.row.url)"
                  size="mini"
                  type="primary"
                  >查看
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-aside>
        <el-aside class="home-aside" width="35%">
          <div class="aside-container">
            <el-card :body-style="{ padding: '0px' }">
              <img class="aside-img" src="../assets/home-aside-img.png" />
              <div style="padding: 20px">
                <a
                  href="http://www.sdnu.edu.cn/bwcxljsm.htm"
                  style="text-decoration:none"
                  target="_blank"
                >
                  <div class="aside-title">
                    <b>“不忘初心、牢记使命”主题教育专题</b>
                  </div>
                  <div class="aside-content">
                    开展“不忘初心、牢记使命”主题教育，是以习近平同志为核心的党中央统揽伟大斗争、伟大工程、伟大事业、伟大梦想作出的重大部署
                    。 按照党中央决策部署和省委工作要求， 学校为第二批开展
                    “不忘初心、牢记使命” 主题...
                  </div>
                </a>
              </div>
            </el-card>
          </div>
        </el-aside> -->
      </el-container>
    </div>
  </div>
</template>

<script>
import * as api from "../api/news";
import * as userApi from "../api/user";
// import * as echarts from 'echarts'
// Vue.prototype.$echarts = echarts

export default {
  name: "Home",
  data() {
    return {
      tableData: [],
      option_column: {
        title: { text: "Column Chart" },
        tooltip: {},

        xAxis: {
          data: [""],
          show: false
        },
        yAxis: {},
        series: [
          {
            name: "销量",
            type: "pie",
            data: [{ value: 335, name: '直接访问' }, { value: 335, name: '直接访问' }],
          },
        ],
      },
    };
  },
  methods: {
    query() {
      api.get().then((res) => {
        this.tableData = res;
      });
    },
    openNews(url) {
      window.open(url, "_blank");
    },
    drawLineChart() {
      userApi.getLoginStatus().then((res) => {
        if (res.usertype == 1) {
          api.getStudent().then((res) => {
            this.option_column.title = { text: "学生成绩表" },
              this.option_column.xAxis = {
                data: ["0~60分", "61~80分", "81~100分"],
                show: true
              }
            this.option_column.series = [
              {
                name: "成绩",
                type: "bar",
                data: res,
              },
            ]
          })
        } else if (res.usertype == 2) {
          api.getTeacher().then((res) => {
            this.option_column.title = { text: "学生评教评分图" },
              this.option_column.series = [
                {
                  name: "评分",
                  type: "pie",
                  data: [{ value: res[0], name: '0分' },
                  { value: res[1], name: '1分' },
                  { value: res[2], name: '2分' },
                  { value: res[3], name: '3分' },
                  { value: res[4], name: '4分' },
                  { value: res[5], name: '5分' },
                  ],
                },
              ]
          })
        } else if (res.usertype == 3) {
          api.getAdmin().then((res) => {
            this.option_column.title = { text: "学生评教评分图" },
              this.option_column.series = [
                {
                  name: "评分",
                  type: "pie",
                  data: [{ value: res[0], name: '0分' },
                  { value: res[1], name: '1分' },
                  { value: res[2], name: '2分' },
                  { value: res[3], name: '3分' },
                  { value: res[4], name: '4分' },
                  { value: res[5], name: '5分' },
                  ],
                },
              ]
          })
        } else {

        }
      })
      // this.$echarts.init(document.getElementById(id)).dispose(); //初始化echarts之前先手动销毁之前的echarts图，防止显示错乱
      // let lineChart = this.$echarts.init(document.getElementById(id)); //初始化echarts
      // // 添加配置项
      // lineChart.setOption(
      //   {
      //     xAxis: {
      //       type: 'category',
      //       data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
      //     },
      //     yAxis: {
      //       type: 'value'
      //     },
      //     series: [
      //       {
      //         data: [150, 230, 224, 218, 135, 147, 260],
      //         type: 'line'
      //       }
      //     ]
      //   }
      // );
    },

  },
  created() {
    // this.query();
  },
  mounted() {
    this.drawLineChart()
  }

};
</script>

<style scoped>
.home-wrap {
  height: 100%;
  width: 100%;
}

.main-wrap {
  margin-top: 10px;
  height: 100%;
  width: 100%;
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.home-container {
  height: 100%;
  width: 90%;
  margin: auto;
}

.home-aside {
  height: 100%;
}

.aside-container {
  width: 90%;
  margin-left: auto;
  margin-right: auto;
  height: 200px;
  margin-top: 20px;
}

.aside-img {
  width: 100%;
}

.aside-title {
  color: #333;
  font-size: 18px;
}

.aside-content {
  font-size: 12px;
  color: #999;
}
</style>
