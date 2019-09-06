const initheaders = function (config){
  config.headers.token = localStorage.token
  
  return config;
}

export default {
  initheaders:initheaders
}