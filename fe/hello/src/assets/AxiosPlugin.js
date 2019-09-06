import axios from 'axios'
import Vue from 'vue'
import headers from './headers'

const Axios = axios.create({
  timeout: 8000,
})

Axios.interceptors.request.use(
    config => {
        return headers.initheaders(config)
    },
    error => {
        return Promise.reject(error);
    }
);


export default {
  install(Vue) {
    Object.defineProperty(Vue.prototype, '$http', { value: Axios })
  }
}
