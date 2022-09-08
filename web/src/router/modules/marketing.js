/** When your routing table is too long, you can split it into small modules **/

import Layout from "@/layout";

const marketingRouter = {
  path: "/marketing",
  component: Layout,
  redirect: "/marketing/campaign",
  meta: { title: "广告投放", icon: "el-icon-s-promotion" },
  children: [
    {
      path: "campaign",
      name: "CampaignList",
      component: () => import("@v/marketing/campaign"),
      meta: { title: "计划列表" },
    },
    {
      path: "adgroup",
      name: "AdgroupList",
      component: () => import("@v/marketing/adgroup"),
      meta: { title: "任务列表" },
    },
    {
      path: "creative",
      name: "CreativeList",
      component: () => import("@v/marketing/creative"),
      meta: { title: "创意列表" },
    },
    {
      hidden: true,
      path: "adgroup-create",
      name: "AdgroupCreate",
      component: () => import("@v/marketing/components/adgroup-create"),
      meta: { title: "任务创建" },
    },
  ],
};

export default marketingRouter;
