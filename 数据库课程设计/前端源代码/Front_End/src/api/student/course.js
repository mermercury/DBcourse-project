import * as ajax from "../../common/ajax";

export const list = () => ajax.pureGet("/student/course");

export const deleteItem = (studentCourseId) =>
  ajax.pureDelete("/student/course/" + studentCourseId);
