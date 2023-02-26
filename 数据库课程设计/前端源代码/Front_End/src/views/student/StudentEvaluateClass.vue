<template>
  <div class="student-evaluate-class">
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-fa fa-edit"></i> 学生评课
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">
      <div class="table">
        <el-table :data="tableData" stripe>
          <el-table-column label="选课Id" prop="courseId" />
          <el-table-column label="课程名" prop="courseName" />
          <el-table-column label="教师" prop="teacherName" />
          <el-table-column label="学分" prop="credit" />
          <el-table-column align="center" label="操作" width="200px">
            <template slot-scope="scope">
              <el-button @click="evaluateClass(scope.row.courseId)" size="mini"
                type="primary">评价
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <el-dialog
      title="学生评课"
      :visible.sync="showWin"
      width="50%"
      :before-close="handleClose"
    >
      <div class="block">
        <span style="align-items: center;">
          请选择您对该课程的评分
        </span>
        <div style="margin: 15px 0;"></div>
        <el-rate v-model="evaluation.star" :colors="colors"> </el-rate>
        <div style="margin: 20px 0;"></div>
        <el-input
          type="textarea"
          placeholder="请问您对本课程有何意见或建议"
          v-model="evaluation.textarea"
          maxlength="200"
          show-word-limit
          :autosize="{ minRows: 5, maxRows: 6 }"
        >
        </el-input>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="closeWin()">取 消</el-button>
        <el-button type="primary" @click="uploadEvaluation()">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import * as api from "../../api/student/evaluateCourse";

export default {
  name: "StudentEvaluateClass",

  data() {
    return {
      tableData: [],
      showWin: false,
      selectedClass: 0,
      colors: ["#99A9BF", "#F7BA2A", "#FF9900"],

      evaluation: {
        textarea: "",
        star: 0,
      }
    };
  },
  methods: {
    getList() {
      api.list().then(res => {
        this.tableData = res;
      })
    },
    evaluateClass(studentCourseId) {
      this.showWin = true;
      this.selectedClass = studentCourseId;
      this.evaluation = {
        textarea: "",
        star: 0,
        courseId: studentCourseId,
      };
    },
    uploadEvaluation() {
      this.closeWin();
      api.create(this.evaluation).then(() => {
        this.$message.success("评价成功");
        this.getList();
      })
    },
    closeWin() {
      this.showWin = false;
    },
  },
  created() {
    this.getList();
  }
};
</script>

<style scoped>

</style>
