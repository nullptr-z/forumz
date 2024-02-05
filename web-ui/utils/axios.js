import axios from "axios";

export function initAxios() {
  axios.defaults.baseURL = 'http://127.0.0.1:8899';
  // axios.defaults.headers = { 'Content-Type': 'application/json' }

  // 请求拦截器
  axios.interceptors.request.use(
    config => {
      // 检查本地存储是否有token
      const token = localStorage.getItem('token');
      if (token) {
        // 如果有token，则将其添加到所有请求头中
        config.headers['Authorization'] = 'Bearer ' + token;
      }
      return config;
    },
    error => {
      // 请求错误处理
      Promise.reject(error);
    }
  );


  axios.interceptors.response.use(
    rsp => {
      if (rsp.status == 200) {
        return rsp.data
      }
      return rsp
    },
    error => {
      // 处理错误响应
      return Promise.reject(error);
    }

  )
}

initAxios(axios)

export default axios
