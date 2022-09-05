import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/queryWyuUser',
    method: 'post',
    data
  })
}


export function register(data) {
  return request({
    url: '/createWyuUser',
    method: 'post',
    data
  })
}