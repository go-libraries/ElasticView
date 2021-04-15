import request from '@/utils/request'

const api = '/api/es_doc/'

export function DeleteRowByIDAction(data) {
  return request({
    url: api + 'DeleteRowByIDAction',
    method: 'post',
    data
  })
}
