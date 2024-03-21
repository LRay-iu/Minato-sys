//写一个路由器并暴露出去
import { createRouter, createWebHistory } from "vue-router";
import HomePage from "@/pages/HomePage.vue";

//创建路由器
const router = createRouter({
  //路由器工作模式,很重要👇👇👇
  history: createWebHistory(),
  routes: [
    // {
    //     path:'路径',
    //     component:组件
    // }
    {
      path: "/home",
      component: HomePage,
    },
    {
        path: "/",
        redirect: "/home",
      },
  ],
});

export default router;