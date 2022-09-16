import request from '../interceptors/marketing'

export function campaignList(params) {
  return request({
    url: "/marketing/campaign/list",
    method: "get",
    params,
  });
}

export function campaignCreate(data) {
  return request({
    url: "/marketing/campaign/create",
    method: "post",
    data
  });
}

export function campaignResources() {
  return request({
    url: "/marketing/campaign/resources",
    method: "get",
  });
}

export function targetingCreate(data) {
  return request({
    url: "/marketing/targeting/create",
    method: "post",
    data
  });
}

export function targetingList(params) {
  return request({
    url: "/marketing/targeting/list",
    method: "get",
    params
  });
}

export function targetingLocation() {
  return request({
    url: "/marketing/targeting/location",
    method: "get"
  });
}

export function dictionaryQuery(params) {
  return request({
    url: "/marketing/dictionary/query",
    method: "get",
    params
  });
}
