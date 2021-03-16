import request from '@/utils/request'

const api = '/api/es_map/'

export function ListAction(data) {
  return request({
    url: api + 'ListAction',
    method: 'post',
    data
  })
}
