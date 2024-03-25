//写一个路由器并暴露出去
import { createRouter, createWebHistory } from "vue-router"
import HomePage from "@/pages/HomePage.vue"
import Login from "@/pages/Login.vue"
import Register from "@/pages/Register.vue"
import userClaim from "@/pages/user_Claim.vue"
import userClaimQuery from "@/pages/user_ClaimQuery.vue"
import userInsuranceBought from "@/pages/user_InsuranceBought.vue"
import userInsuranceQuery from "@/pages/user_InsuranceQuery.vue"
Request
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
            name: "home",
            path: "/",
            component: HomePage,
        },
        {
            name: "login",
            path: "/login",
            component: Login,
        },
        {
            name: "register",
            path: "/register",
            component: Register,
        },
        {
            name: "userInsuranceBought ",
            path: "/2-1",
            component: userInsuranceBought,
        },
        {
            name: "userInsuranceQuery",
            path: "/2-2",
            component: userInsuranceQuery,
        },
        {
            name: "userClaim",
            path: "/3-1",
            component: userClaim,
        },
        {
            name: "claimQuery",
            path: "/3-2",
            component: userClaimQuery,
        },
    ],
})

export default router
