/** When your routing table is too long, you can split it into small modules **/

import Layout from "@/layout";

const userRouter = {
  path: "/account",
  component: Layout,
  redirect: "/account/list",
  meta: { title: "账户&应用", icon: "el-icon-s-grid" },
  children: [
    {
      path: "list",
      name: "AccountList",
      component: () => import("@v/account/list"),
      meta: { title: "账户列表" },
    },
    {
      path: "app",
      name: "AppList",
      component: () => import("@v/application/list"),
      meta: { title: "应用列表" },
    },
  ],
};

export default userRouter;
