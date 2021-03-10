
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
  }
}

const mutations = {
  SET_EsConnect: (state, EsConnect) => {
    state.EsConnect = EsConnect
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
