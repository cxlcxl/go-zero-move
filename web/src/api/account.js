import request from "@/interceptors/account";

export function accountUpdate(data) {
  return request({
    url: "/account/update",
    method: "post",
    data,
  });
}

export function accountCreate(data) {
  return request({
    url: "/account/create",
    method: "post",
    data,
  });
}

export function accountList(params) {
  return request({
    url: "/account/list",
    method: "get",
    params,
  });
}

export function accountInfo(account_id) {
  return request({
    url: "/account/info",
    method: "get",
    params: { id: account_id },
  });
}

export function accountAuth(account_id) {
  return request({
    url: "/account/auth",
    method: "get",
    params: { id: account_id },
  });
}

export function refreshAuth(id) {
  return request({
    url: "/account/refresh",
    method: "post",
    data: { id },
  });
}

export function searchAccounts(accountName) {
  return request({
    url: "/account/search",
    method: "get",
    params: { account_name: accountName },
  });
}

export function defaultAccounts() {
  return request({
    url: "/account/default",
    method: "get",
  });
}

export function parentAccounts(params) {
  return request({
    url: "/account/parents",
    method: "get",
    params
  });
}

export function getAccessToken(data) {
  return request({url: '/account/token', method: 'post', data})
}
