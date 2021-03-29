import request from '@/utils/request'

const api = '/api/es_link/'

export function DeleteAction(data) {
  return request({
    url: api + 'DeleteAction',
    method: 'post',
    data
  })
}

export function ListAction() {
  return request({
    url: api + 'ListAction',
    method: 'get'
  })
}

export function UpdateAction(data) {
  return request({
    url: api + 'UpdateAction',
    method: 'post',
    data
  })
}

export function InsertAction(data) {
  return request({
    url: api + `InsertAction`,
    method: 'post',
    data
  })
}
