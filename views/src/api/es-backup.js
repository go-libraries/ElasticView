import request from '@/utils/request'

const api = '/api/backUp/'

export function SnapshotListAction(data) {
  return request({
    url: api + 'SnapshotListAction',
    method: 'post',
    data
  })
}

export function SnapshotCreateRepositoryAction(data) {
  return request({
    url: api + 'SnapshotCreateRepositoryAction',
    method: 'post',
    data
  })
}
export function SnapshotDeleteRepositoryAction(data) {
  return request({
    url: api + 'SnapshotDeleteRepositoryAction',
    method: 'post',
    data
  })
}
