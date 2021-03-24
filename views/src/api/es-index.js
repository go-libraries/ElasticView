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

export function IndexNamesAction(data) {
  return request({
    url: api + 'IndexNamesAction',
    method: 'post',
    data
  })
}

export function ReindexAction(data) {
  return request({
    url: api + 'ReindexAction',
    method: 'post',
    data
  })
}

export function GetAliasAction(data) {
  return request({
    url: api + 'GetAliasAction',
    method: 'post',
    data
  })
}

