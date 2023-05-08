import axios from "axios";
import { getToken, removeToken } from "@/utils/cookie";
import userToken from "@/store/modules/user";
import { ElMessage } from "element-plus";
import Cookies from "js-cookie";

const http = axios.create({
  baseURL: import.meta.env.VITE_SERVER_HOST,
  // baseURL: "",
});

http.interceptors.request.use(
  (config) => {
    if (userToken.getters.token) {
      config.headers["Authorization"] = getToken();
      config.headers["Accept-Language"] =
        localStorage.getItem("lang") == "en" ? "en" : "ch" || "ch";
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);
http.interceptors.response.use(
  (res) => {
    // if (res.data.err_code === 403) {
    //   console.log("here error code 403");
    // }
    //if the custom code is not 20000, it is judged as an error.
    if (res.data.err_code !== 0 && res.data.err_code !== 200) {
      ElMessage({
        message: res.data.err_msg || "Error",
        type: "error",
      });

      // if (res.err_code === "11005") {
      //   return res;
      // }

      //   if (res.err_code === '15012') {
      //     Notification({
      //       title: '错误',
      //       message: '账户不存在',
      //       type: 'error',
      //     });
      //   }
      //   // if (res.err_code == '100') {
      //   //   Notification({
      //   //     title: '错误',
      //   //     message: '11111',
      //   //     type: 'error',
      //   //   });
      //   // }

      return res;

      // return Promise.reject(new Error(res || "Error"));
    } else {
      return res;
    }
  },
  (error) => {
    if (error.response.data.err_msg == "Group URL Duplicate ") {
      // ElMessage.error("域名已存在!");
    } else if (error.response.data.err_msg == "Group Name Duplicate ") {
      // ElMessage.error("分组名称已存在!");
    } else {
      ElMessage.error(error.response.data.err_msg);
    }
    if (error.response.status === 401) {
      ElMessage.error(error.response.data.err_msg);
      removeToken();
      Cookies.remove("timer_access");
      // location.reload();
    }
    return Promise.reject(error);
  }
);

export default http;
