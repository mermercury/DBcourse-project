import * as ajax from "../../common/ajax";

export const get = id => ajax.pureGet("/teacher/grade/" + id);

export const update = entity => ajax.put("/teacher/grade", entity);

export const getPageCount = (courseName, studentName, pageSize) =>
  ajax.get("/teacher/grade/page/count", {
    courseName: courseName,
    studentName: studentName,
    pageSize: pageSize
  });

export const getPage = (index, courseName, studentName, pageSize) =>
  ajax.get("/teacher/grade/page/" + index, {
    courseName: courseName,
    studentName: studentName,
    pageSize: pageSize
  });

export const pageSize = 20;
