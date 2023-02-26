import * as ajax from "../../common/ajax";

export const select = (id) => ajax.post("/student/course/select" , {courseId:id});
export const getPageCount = (courseName, teacherName) =>
  ajax.get("/student/course/page/count", {
    courseName: courseName,
    teacherName: teacherName,
    pageSize: 6
  });

export const getPage = (index, courseName, teacherName) =>
  ajax.get("/student/course/page/" + index, {
    courseName: courseName,
    teacherName: teacherName,
    pageSize: 6
  });

export const pageSize = 6; // 一页最多放6门课
