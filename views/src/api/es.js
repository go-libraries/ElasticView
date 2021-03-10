import request from '@/utils/request'

const api = '/api/es/'

export function PingAction(data) {
  return request({
    url: api + `PingAction`,
    method: 'post',
    data
  })
}
