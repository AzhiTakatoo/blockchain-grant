import request from '@/utils/request'

export function createQueryStipendRanking(data) {
  return request({
    url: '/createQueryStipendRanking',
    method: 'post',
    data
  })
}