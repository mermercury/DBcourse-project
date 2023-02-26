import * as ajax from "../../common/ajax";

export const create = entity => ajax.put("/student/evaluateCourse", entity);
// don't need to pass id here
export const list = () => ajax.pureGet("/student/evaluateCourse");
