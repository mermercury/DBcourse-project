<template>
  <div class="course-wrap">
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-fa fa-book"></i> 课程管理
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="container">
      <div class="query-form">
        <el-row :gutter="20">
          <el-col :span="2">
            <el-button @click="create" icon="el-icon-plus">创建</el-button>
          </el-col>
          <el-col :offset="10" :span="3">
            <el-input
              @keyup.enter.native="query"
              placeholder="课程名"
              v-model="queryForm.name"
            />
          </el-col>
          <el-col :span="3">
            <el-input
              @keyup.enter.native="query"
              placeholder="教师名"
              v-model="queryForm.teacherName"
            />
          </el-col>
          <el-col :span="3">
            <el-input
              @keyup.enter.native="query"
              placeholder="系名"
              v-model="queryForm.departmentName"
            />
          </el-col>
          <el-col :span="3">
            <el-button
              @click="query"
              icon="el-icon-search"
              type="primary"
              class="add"
              >搜索
            </el-button>
          </el-col>
        </el-row>
      </div>

      <el-row justify="center" type="flex">
        <el-pagination
          :current-page.sync="pageIndex"
          :page-size="pageSize"
          :total="pageSize * pageCount"
          @current-change="getPage"
          background
          layout="prev, pager, next"
        >
        </el-pagination>
      </el-row>

      <div class="table">
        <el-table :data="tableData" stripe>
          <el-table-column label="课程Id" prop="id" />
          <el-table-column label="课程名" prop="courseName" />
          <el-table-column label="教师ID" prop="teacherID" />
          <el-table-column label="所属系" prop="departmentName" />
          <el-table-column label="年级" prop="grade" />
          <el-table-column label="学分" prop="credit" />
          <el-table-column
            align="center"
            label="上课时间"
            prop="courseTime"
            width="130px"
          />
          <el-table-column label="已选人数" prop="selected" />
          <el-table-column label="最大容量" prop="size" />
          <el-table-column align="center" label="操作" width="200px">
            <template slot-scope="scope">
              <el-button @click="edit(scope.row.id)" size="mini" type="success"
                >编辑
              </el-button>
              <el-button
                @click="deleteItem(scope.row.id)"
                size="mini"
                type="danger"
                >删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <el-dialog :visible.sync="editing" title="编辑" width="30%">
        <el-form :model="entityForm" label-width="70px" ref="form">
          <el-form-item label="课程名">
            <el-input v-model="entityForm.courseName"></el-input>
          </el-form-item>
          <el-form-item label="所属系">
            <el-select
              placeholder="请选择系"
              v-model="entityForm.departmentName"
            >
              <el-option
                :key="index"
                :label="item.name"
                :value="item.name"
                v-for="(item, index) in departments"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="授课教师">
            <el-select placeholder="请选择教师" v-model="entityForm.teacherID">
              <el-option
                :key="index"
                :label="item.name"
                :value="item.id"
                v-for="(item, index) in teachers"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="年级">
            <el-input type="number" v-model="entityForm.grade"></el-input>
          </el-form-item>
          <el-form-item label="上课时间">
            <el-select v-model="courseDay">
              <el-option
                :key="index"
                :label="item"
                :value="index"
                v-for="(item, index) in days"
              >
              </el-option>
            </el-select>
            <el-select v-model="courseT">
              <el-option
                :key="index"
                :label="item"
                :value="index"
                v-for="(item, index) in times"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="时长(节)">
            <el-input type="number" v-model="courseLength"></el-input>
          </el-form-item>
          <el-form-item label="上课地点">
            <el-input v-model="entityForm.location"></el-input>
          </el-form-item>
          <el-form-item label="学分">
            <el-input type="number" v-model="entityForm.credit"></el-input>
          </el-form-item>
          <el-form-item label="最大容量">
            <el-input type="number" v-model="entityForm.size"></el-input>
          </el-form-item>
          <el-form-item label="考试时间">
            <el-date-picker
              format="yyyy-MM-dd HH:mm"
              type="datetime"
              v-model="entityForm.examDate"
              value-format="yyyy-MM-dd HH:mm"
            >
            </el-date-picker>
          </el-form-item>
          <el-form-item label="考试地点">
            <el-input v-model="entityForm.examLocation"></el-input>
          </el-form-item>
        </el-form>
        <span class="dialog-footer" slot="footer">
          <el-button @click="save" type="primary">确 定</el-button>
          <el-button @click="editing = false">取 消</el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import * as api from "../../api/admin/course";
import * as teacherApi from "../../api/admin/teacher";

export default {
  name: "AdminCourse",
  data() {
    return {
      queryForm: {
        departmentName: "",
        teacherName: "",
        name: "",
      },
      entityForm: {},
      departments: [],
      tableData: [],
      pageSize: api.pageSize,
      pageCount: 1,
      pageIndex: 1,
      editing: false,
      teachers: [],
      courseDay: "",
      courseT: "",
      courseLength: 0,
      days: [
        "星期一",
        "星期二",
        "星期三",
        "星期四",
        "星期五",
        "星期六",
        "星期日",
      ],
      times: [
        "第一节",
        "第二节",
        "第三节",
        "第四节",
        "第五节",
        "第六节",
        "第七节",
        "第八节",
        "第九节",
      ],
    };
  },
  methods: {
    query() {
      // api
      //   .getPageCount(
      //     this.queryForm.departmentName,
      //     this.queryForm.teacherName,
      //     this.queryForm.name
      //   )
      //   .then(res => {
      //     this.pageCount = res;
      //     this.pageIndex = 1;
      this.getPage(1);
      //   });
    },
    getPage(pageIndex) {
      api
        .getPage(
          pageIndex,
          this.queryForm.departmentName,
          this.queryForm.teacherName,
          this.queryForm.name
        )
        .then((res) => {
          this.tableData = res;
        });
    },
    create() {
      this.entityForm = {
        id: -1,
        teacherID: null,
        name: "",
        grade: 2019,
        time: "",
        location: "",
        credit: 2,
        size: 50,
        examDate: null,
        examLocation: null,
        departmentName: null,
      };
      this.courseDay = 1;
      this.courseTime = 1;
      this.courseLength = 2;
      this.editing = true;
    },
    edit(id) {
      api.get(id).then((res) => {
        // let res = this.tableData[0];
        let split = res.courseTime.split("-");
        this.courseDay = parseInt(split[0]) - 1;
        this.courseT = parseInt(split[1]) - 1;
        this.courseLength = parseInt(split[2]);
        this.entityForm = res;
        this.editing = true;
      });
    },
    save() {
      let day = this.courseDay + 1;
      let time = this.courseT + 1;
      this.entityForm.time = day + "-" + time + "-" + this.courseLength;
      this.entityForm.credit = Number(this.entityForm.credit);
      this.entityForm.size = Number(this.entityForm.size);
      this.entityForm.grade = Number(this.entityForm.grade);
      if (this.entityForm.id === -1) {
        api.create(this.entityForm).then(() => {
          this.finishSave();
        });
      } else {
        api.update(this.entityForm).then(() => {
          this.finishSave();
        });
      }
    },
    finishSave() {
      this.$message.success("成功");
      this.getPage(this.pageIndex);
      this.editing = false;
    },
    deleteItem(id) {
      api.deleteItem(id).then(() => {
        this.$message.success("删除成功");
        this.getPage(this.pageIndex);
      });
    },
    getDepartments() {
      teacherApi.listName().then((res) => {
        this.departments = res;
      });
    },
    getTeachers() {
      api.listName().then((res) => {
        this.teachers = res;
      });
    },
  },

  created() {
    this.query();
    this.getTeachers();
    this.getDepartments();
  },
};
</script>

<style scoped></style>
