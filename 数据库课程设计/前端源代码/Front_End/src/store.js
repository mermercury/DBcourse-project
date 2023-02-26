import Vue from "vue";
import Vuex from "vuex";
import SideBarItem from "./common/sideBarItem";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    status: {
      loggedIn: false,
      userid: -1,
      username: "",
      usertype: 0,
      permission: 0,
    },
    sideBarItems: [],
    loading: false,
  },
  mutations: {
    login(state, res) {
      state.status.loggedIn = res.loggedIn;
      state.status.userid = res.userid;
      state.status.username = res.username;
      state.status.usertype = res.usertype;
      state.status.permission = res.permission;
      console.log(res)

      let items = [];
      for (let i = 0; i < SideBarItem.items.length; i++) {
        let item = SideBarItem.items[i];
        
        if (
          item.userType === res.usertype &&
          (item.permission & res.permission) === item.permission
        ) {
          items.push(item);
          // console.log("sideBar add one ")
        }
      }
      state.sideBarItems = items;
    },
    logout(state) {
      state.status = {
        loggedIn: false,
        userid: -1,
        username: "",
        usertype: 0,
        permission: 0,
      };
    },
  },
  actions: {},
  modules: {},
});
