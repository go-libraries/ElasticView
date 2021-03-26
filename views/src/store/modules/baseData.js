
const state = {
  EsConnect: {
    ip: '',
    user: '',
    pwd: '',
    remark: '',
    version: 0,
    created: '',
    id: 0,
    updated: ''
  },
  EsConnectID: 0
}

const mutations = {
  SET_EsConnect: (state, EsConnect) => {
    state.EsConnectID = EsConnect
  }
}

const actions = {
  SetEsConnect({ commit }, p) {
    commit('SET_EsConnect', p)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
