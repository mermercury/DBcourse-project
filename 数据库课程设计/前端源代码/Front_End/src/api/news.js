import { pureGet } from "../common/ajax";

// export const get = () => pureGet("/sdnu/news");
export const getAdmin = () => pureGet("/admin/course/static/evaluation");

export const getStudent = () => pureGet("/student/static/score");

export const getTeacher = () => pureGet("/teacher/static/evaluation");