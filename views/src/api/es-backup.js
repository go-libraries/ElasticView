import request from '@/utils/request'

const api = '/api/backUp/'

export function SnapshotListAction(data) {
  return request({
    url: api + 'SnapshotListAction',
    method: 'post',
    data
  })
}
