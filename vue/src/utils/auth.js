import Cookies from 'js-cookie'

const TokenKey = 'grantUser_id_token'
const TokenKeyPwd = 'grantUser_pwd_token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function getTokenPwd() {
  return Cookies.get(TokenKeyPwd)
}
export function setTokenPwd(token)  {
  return Cookies.set(TokenKeyPwd, token)

}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
