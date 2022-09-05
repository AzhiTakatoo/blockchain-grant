import {
  login
} from '@/api/grantUser'
import {
  getToken,
  getTokenPwd,
  setToken,
  setTokenPwd,
  removeToken
} from '@/utils/auth'
import {
  resetRouter
} from '@/router'

const getDefaultState = () => {
  return {
    token: getToken(),
    tokenPwd: getTokenPwd(),
    wyuUserId: '',
    wyuUserName: '',
    balance: 0,
    roles: []
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token.wyuUserId
    state.tokenPwd = token.wyuPasswd
  },
  SET_wyuUserId: (state, wyuUserId) => {
    state.wyuUserId = wyuUserId
  },
  SET_USERNAME: (state, wyuUserName) => {
    state.wyuUserName = wyuUserName
  },
  SET_BALANCE: (state, balance) => {
    state.balance = balance
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  }
}

const actions = {
  login({
    commit
  }, {wyuUserId, wyuPasswd}) {
    return new Promise((resolve, reject) => {
      login({
        wyuUserId,
        wyuPasswd
      }).then(response => {
        // console.log(response);
        commit('SET_TOKEN', {wyuUserId, wyuPasswd})
        // commit('SET')
        setToken(wyuUserId)
        setTokenPwd(wyuPasswd)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  // get user info
  getInfo({
    commit,
    state
  }) {
    return new Promise((resolve, reject) => {
      login({
        wyuUserId: state.token,
        wyuPasswd: state.tokenPwd
      }).then(response => {
        var roles
        if (response[0].wyuUserName === '管理员') {
          roles = ['admin']
        } else {
          roles = ['editor']
        }
        commit('SET_ROLES', roles)

        commit('SET_wyuUserId', response[0].wyuUserId)
        commit('SET_USERNAME', response[0].wyuUserName)
        // commit('SET_BALANCE', response.balance)
        resolve(roles)
      }).catch(error => {
        reject(error)
      })
    })
  },
  logout({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      resetRouter()
      commit('RESET_STATE')
      resolve()
    })
  },

  resetToken({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
