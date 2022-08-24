import request from "@/interceptors/rbac";

export function roleList(data) {
  return request({
    url: "/role/list",
    method: "post",
    data,
  });
}

export function rolePermissions(id) {
  return request({
    url: "/role/permissions",
    method: "get",
    params: { id },
  });
}

export function roleCreate(data) {
  return request({
    url: "/role/create",
    method: "post",
    data,
  });
}

export function roleUpdate(data) {
  return request({
    url: "/role/update",
    method: "post",
    data,
  });
}

export function roleDestroy(id) {
  return request({
    url: "/role/destroy",
    method: "post",
    data: {
      id,
    },
  });
}

// 权限部分
export function permissionList() {
  return request({
    url: "/permission/list",
    method: "get",
  });
}

export function permissionCreate(data) {
  return request({
    url: "/permission/create",
    method: "post",
    data,
  });
}

export function permissionDestroy(id) {
  return request({
    url: "/permission/destroy",
    method: "post",
    data: { id },
  });
}

export function permissionUpdate(data) {
  return request({
    url: "/permission/update",
    method: "post",
    data,
  });
}
