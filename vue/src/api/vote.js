import request from '@/utils/request'

export function createVote(data) {
  return request({
    url: '/createVote',
    method: 'post',
    data
  })
}

export function queryVote(data) {
  return request({
    url: '/queryVote',
    method: 'post',
    data
  })
}

export function queryVoteOnly(data) {
  return request({
    url: '/queryVoteOnly',
    method: 'post',
    data
  })
}

export function updatePower(data) {
  return request({
    url: '/updatePower',
    method: 'post',
    data
  })
}

export function setPower(data) {
  return request({
    url: '/setPower',
    method: 'post',
    data
  })
}