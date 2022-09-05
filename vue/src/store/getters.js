const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.grantUser.token,
  wyuUserId: state => state.grantUser.wyuUserId,
  wyuUserName: state => state.grantUser.wyuUserName,
  balance: state => state.grantUser.balance,
  roles: state => state.grantUser.roles,
  permission_routes: state => state.permission.routes
}
export default getters
