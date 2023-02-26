import { post, pureGet } from "../common/ajax";

export const login = (userid, password, usertype) =>
  post("/login", {
    userid: userid,
    password: password,
    usertype: usertype,
  });

export const getLoginStatus = () => pureGet("/auth/status");

export const logout = () => post("/auth/logout");
