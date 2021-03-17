import request from '@/utils/request'

const api = '/api/es_index/'

export function CreateAction(data) {
  return request({
    url: api + 'CreateAction',
    method: 'post',
    data
  })
}

export function GetSettingsAction(data) {
  return request({
    url: api + 'GetSettingsAction',
    method: 'post',
    data
  })
}
