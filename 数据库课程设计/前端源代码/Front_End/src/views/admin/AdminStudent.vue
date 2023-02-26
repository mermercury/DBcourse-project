<template>
  <div class="student-wrap">
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-fa fa-user"></i> 学生管理
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="container">
      <div class="query-form">
        <el-row :gutter="20">
          <el-col :span="2">
            <el-button @click="create" icon="el-icon-plus">创建</el-button>


          </el-col>
          <el-col :span="5" :offset="2">
            <div>
              <div margin="5px" height="80px">
                <el-upload action="http://139.9.143.161:8080/api/v1/upload/studentTable"
                  :on-preview="handlePreview" :on-remove="handleRemove"
                  :before-remove="beforeRemove" :with-credentials="true" multiple :limit="3"
                  :on-exceed="handleExceed" :file-list="fileList">
                  <el-button type="primary" class="add">点击上传</el-button>
                </el-upload>
              </div>

              <div slot="tip" class="el-upload__tip">
                只能上传xlsx文件，且不超过500kb
              </div>
            </div>
          </el-col>
          <el-col :offset="3" :span="3">
            <el-input @keyup.enter.native="query" placeholder="学生姓名" v-model="queryForm.name" />
          </el-col>
          <el-col :span="3">
            <el-input @keyup.enter.native="query" placeholder="专业名" v-model="queryForm.majorName" />
          </el-col>
          <el-col :span="3">
            <el-input @keyup.enter.native="query" placeholder="班级名" v-model="queryForm.className" />
          </el-col>
          <el-col :span="3">
            <el-button @click="query" icon="el-icon-search" type="primary" class="add">搜索
            </el-button>
          </el-col>
        </el-row>
      </div>

      <el-row justify="center" type="flex">
        <el-pagination :current-page.sync="pageIndex" :page-size="pageSize"
          :total="pageSize * pageCount" @current-change="getPage" background
          layout="prev, pager, next">
        </el-pagination>
      </el-row>

      <div class="table">
        <el-table :data="tableData" stripe>
          <el-table-column label="学号" prop="id" width="80px" />
          <el-table-column label="姓名" prop="name" />
          <el-table-column label="班级" prop="className" />
          <el-table-column label="专业" min-width="150px" prop="majorName" />
          <el-table-column label="性别" prop="sex" width="80px" />
          <el-table-column align="center" label="操作" width="200px">
            <template slot-scope="scope">
              <el-button @click="edit(scope.row.id)" size="mini" type="success">编辑
              </el-button>
              <el-button @click="deleteItem(scope.row.id)" size="mini" type="danger">删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <el-dialog :visible.sync="editing" title="编辑" width="30%">
        <el-form :model="entityForm" label-width="70px" ref="form">
          <el-form-item label="姓名">
            <el-input v-model="entityForm.name"></el-input>
          </el-form-item>
          <el-form-item label="学号">
            <el-input type="id" v-model.number="entityForm.id"></el-input>
          </el-form-item>
          <!-- <el-form-item label="所属专业">
            <el-select placeholder="请选择专业" v-model="entityForm.majorName">
              <el-option
                :key="index"
                :label="item.name"
                :value="item.name"
                v-for="(item, index) in majors"
              />
            </el-select>
          </el-form-item> -->
          <el-form-item label="所属班级">
            <el-select placeholder="请选择班级" v-model="entityForm.className">
              <el-option :key="index" :label="item.name" :value="item.name"
                v-for="(item, index) in classes" />
            </el-select>
          </el-form-item>
          <el-form-item label="密码">
            <el-input type="password" v-model="entityForm.password"></el-input>
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="entityForm.email"></el-input>
          </el-form-item>
          <el-form-item label="生日">
            <el-date-picker format="yyyy-MM-dd" placeholder="选择生日" type="date"
              v-model="entityForm.birthday">
            </el-date-picker>
          </el-form-item>
          <el-form-item label="性别">
            <el-radio-group v-model="entityForm.sex">
              <el-radio :label="1">男</el-radio>
              <el-radio :label="0">女</el-radio>
            </el-radio-group>
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
import * as api from "../../api/admin/student";
import * as classApi from "../../api/admin/class";
import * as majorApi from "../../api/admin/major";

export default {
  name: "AdminStudent",
  data() {
    return {
      queryForm: {
        majorName: "",
        className: "",
        name: "",
      },
      entityForm: {},
      tableData: [],
      pageSize: api.pageSize,
      pageCount: 1,
      pageIndex: 1,
      editing: false,
      classes: [],
      fileList: [],
      majors: [],
    };
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(file);
    },
    handleExceed(files, fileList) {
      this.$message.warning(
        `当前限制选择 3 个文件，本次选择了 ${files.length
        } 个文件，共选择了 ${files.length + fileList.length} 个文件`
      );
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },

    query() {
      api
        .getPageCount(
          this.queryForm.majorName,
          this.queryForm.className,
          this.queryForm.name
        )
        .then((res) => {
          this.pageCount = res;
          this.pageIndex = 1;
          this.getPage(1);
        });
    },
    getPage(pageIndex) {
      api
        .getPage(
          pageIndex,
          this.queryForm.majorName,
          this.queryForm.className,
          this.queryForm.name
        )
        .then((res) => {
          for (let i = 0; i < res.length; i++) {
            res[i].sex = res[i].sex === 1 ? "男" : "女";
          }
          this.tableData = res;
        });
    },
    create() {
      this.entityForm = {
        id: -1,
        name: "",
        className: "",
        majorName: "",
        password: "",
        email: null,
        birthday: null,
        sex: 0,
      };
      this.editing = true;
    },
    edit(id) {
      api.get(id).then((res) => {
        this.entityForm = res;
        this.editing = true;
      });
    },
    save() {
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
    getClasses() {
      classApi.listName().then((res) => {
        this.classes = res;
      });
    },
    getMajors() {
      majorApi.listName().then((res) => {
        this.majors = res;
      });
    },
  },
  created() {
    this.query();
    this.getClasses();
    this.getMajors();
  },
};
</script>

<style scoped>

</style>
