import request from '../interceptors/marketing'

export function creativeCategory() {
  return request({
    url: "/marketing/resource/creative-category",
    method: "get"
  });
}

export function pricing() {
  return request({
    url: "/marketing/resource/pricing",
    method: "get"
  });
}
