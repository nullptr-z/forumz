import axios from "axios";

export function initAxios() {
  axios.defaults.baseURL = 'http://127.0.0.1:8899';
  // axios.defaults.headers = { 'Content-Type': 'application/json' }

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
