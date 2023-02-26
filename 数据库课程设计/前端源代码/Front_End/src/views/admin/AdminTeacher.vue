<template>
  <div class="teacher-wrap">
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-fa fa-user-plus"></i> 教师管理
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
            <div margin="5px" height="80px">
              <el-upload action="http://139.9.143.161:8080/api/v1/upload/teacherTable"
                :on-preview="handlePreview" :on-remove="handleRemove" :before-remove="beforeRemove"
                :with-credentials="true" multiple :limit="3" :on-exceed="handleExceed"
                :file-list="fileList">
                <el-button type="primary" class="add">点击上传</el-button>
                <div slot="tip" class="el-upload__tip">
                  只能上传xlsx文件，且不超过500kb
                </div>
              </el-upload>
            </div>
          </el-col>
          <el-col :offset="6" :span="3">
            <el-input @keyup.enter.native="query" placeholder="教师" v-model="queryForm.name" />
          </el-col>
          <el-col :span="3">
            <el-input @keyup.enter.native="query" placeholder="系名"
              v-model="queryForm.departmentName" />
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
          <el-table-column label="教师Id" prop="id" />
          <el-table-column label="教师姓名" prop="name" />
          <el-table-column label="教师电话" prop="number" />
          <el-table-column label="所属系" prop="departmentName" />
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
          <el-form-item label="教师姓名">
            <el-input v-model="entityForm.name"></el-input>
          </el-form-item>
          <el-form-item label="教师电话">
            <el-input v-model="entityForm.number"></el-input>
          </el-form-item>
          <el-form-item label="所属系">
            <el-select placeholder="请选择系" v-model="entityForm.departmentName">
              <el-option :key="index" :label="item.name" :value="item.name"
                v-for="(item, index) in departments" />
            </el-select>
          </el-form-item>
          <el-form-item label="登录密码">
            <el-input type="password" v-model="entityForm.password"></el-input>
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
import * as api from "../../api/admin/teacher";
import * as departmentApi from "../../api/admin/department";

export default {
  name: "AdminTeacher",
  data() {
    return {
      queryForm: {
        departmentName: "",
        name: ""
      },
      entityForm: {},
      tableData: [],
      pageSize: api.pageSize,
      pageCount: 1,
      pageIndex: 1,
      editing: false,
      departments: []
    };
  },
  methods: {
    query() {
      api
        .getPageCount(this.queryForm.departmentName, this.queryForm.name)
        .then(res => {
          this.pageCount = res;
          this.pageIndex = 1;
          this.getPage(1);
        });
    },
    getPage(pageIndex) {
      api
        .getPage(pageIndex, this.queryForm.departmentName, this.queryForm.name)
        .then(res => {
          this.tableData = res;
        });
    },
    create() {
      this.entityForm = {
        id: -1,
        number: "",
        name: "",
        departmentName: "",
        password: ""
      };
      this.editing = true;
    },
    edit(id) {
      api.get(id).then(res => {
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
    getDepartments() {
      departmentApi.listName().then(res => {
        this.departments = res;
      });
    }
  },
  created() {
    this.query();
    this.getDepartments();
  }
};
</script>

<style scoped>

</style>
