import request from "@/interceptors/marketing";

export function accountUpdate(data) {
  return request({
    url: "/marketing/account/update",
    method: "post",
    data,
  });
}

export function accountCreate(data) {
  return request({
    url: "/marketing/account/create",
    method: "post",
    data,
  });
}

export function accountList(params) {
  return request({
    url: "/marketing/account/list",
    method: "get",
    params,
  });
}

export function accountInfo(account_id) {
  return request({
    url: "/marketing/account/info",
    method: "get",
    params: { id: account_id },
  });
}

export function accountAuth(account_id) {
  return request({
    url: "/marketing/account/auth",
    method: "get",
    params: { id: account_id },
  });
}

export function refreshAuth(client_id) {
  return request({
    url: "/marketing/refresh",
    method: "post",
    data: { client_id },
  });
}

export function searchAccounts(accountName) {
  return request({
    url: "/marketing/account/search",
    method: "get",
    params: { account_name: accountName },
  });
}

export function defaultAccounts() {
  return request({
    url: "/marketing/account/default",
    method: "get",
  });
}
