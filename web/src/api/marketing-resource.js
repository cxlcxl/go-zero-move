import request from '../interceptors/marketing'

export function pricing() {
  return request({
    url: "/marketing/resource/pricing",
    method: "get"
  });
}

export function trackingList(params) {
  return request({
    url: "/marketing/tracking/list",
    method: "get",
    params
  });
}

export function trackingRefresh(params) {
  return request({
    url: "/marketing/tracking/refresh",
    method: "get",
    params
  });
}

// export const assetUpload = 'https://ads-dra.cloud.huawei.com/ads/v1/tools/creative_asset/create'
export const assetUpload = process.env.VUE_APP_MARKETING_API + "/marketing/asset/upload"

export function syncAsset(params) {
  return request({
    url: "/marketing/asset/sync",
    method: "get",
    params
  });
}

export function assetList(params) {
  return request({
    url: "/marketing/asset/list",
    method: "get",
    params
  });
}

export function assetDelete(data) {
  return request({
    url: "/marketing/asset/delete",
    method: "post",
    data
  });
}

export function assetDimension() {
  return request({
    url: "/marketing/asset/dimension",
    method: "get"
  });
}

export function assetToken(params) {
  return request({
    url: "/marketing/asset/token",
    method: "get",
    params
  });
}

export function bindAsset(data) {
  return request({
    url: "/marketing/asset/bind",
    method: "post",
    data
  });
}
