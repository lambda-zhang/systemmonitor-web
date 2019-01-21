import axios from 'axios'
import {getAPIUrl} from './window_location.js'

var APIhost = getAPIUrl()

// http request 拦截器
axios.interceptors.request.use(
  config => {
    return config
  },
  err => {
    return Promise.reject(err)
  }
)

// http response 拦截器
axios.interceptors.response.use(
  response => {
    if (!response || !response.hasOwnProperty('data') ||
      !response.data.hasOwnProperty('data') ||
      !response.data.data.hasOwnProperty('UpdatedAt')) {
      return {'data': null}
    }
    return response
  },
  error => {
    if (error.response) {
      return error.response
    } else {
      return {'data': null}
    }
  }
)

export function post (path, data = {}) {
  var url = APIhost + path
  return new Promise((resolve, reject) => {
    axios.post(url, data).then(response => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}

export function get (path, data = {}) {
  var url = APIhost + path
  return new Promise((resolve, reject) => {
    axios.get(url, data).then(response => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}
