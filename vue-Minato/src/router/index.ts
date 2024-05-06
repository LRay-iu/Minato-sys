//写一个路由器并暴露出去
import { createRouter, createWebHistory } from "vue-router"
import HomePage from "@/pages/HomePage.vue"
import Login from "@/pages/Login.vue"
import Register from "@/pages/Register.vue"
import userClaim from "@/pages/user_Claim.vue"
import userClaimQuery from "@/pages/user_ClaimQuery.vue"
import userInsuranceBought from "@/pages/user_InsuranceBought.vue"
import userInsuranceQuery from "@/pages/user_InsuranceQuery.vue"
import adminClaim from "@/pages/admin_Claim.vue"
import adminClaimDetail from "@/pages/admin_Claim_Detail.vue"
import { ElMessage } from "element-plus"
import adminSearch from "@/pages/admin_Search.vue"
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
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "user") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "您不能这么做！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "userInsuranceQuery",
            path: "/2-2",
            component: userInsuranceQuery,
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "user") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "您不能这么做！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "userInsuranceQuery",
            path: "/2-5",
            component: adminSearch,
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "admin") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "权限不足！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "userClaim",
            path: "/3-1",
            component: userClaim,
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "user") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "您不能这么做！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "claimQuery",
            path: "/3-2",
            component: userClaimQuery,
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "user") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "您不能这么做！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "adminClaim",
            path: "/3-4",
            component: adminClaim,
            //设置路由守卫
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "admin") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "权限不够！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
        {
            name: "ClaimDetail",
            path: "/3-4/detail/:claimid",
            component: adminClaimDetail,
            beforeEnter: (to, from, next) => {
                if (localStorage.getItem("role") != "admin") {
                    if (to.name == "login") {
                        next()
                    } else {
                        ElMessage({
                            message: "权限不够！",
                            type: "error",
                        })
                        setTimeout(() => {
                            router.push("login")
                        }, 1000)
                    }
                } else {
                    next()
                }
            },
        },
    ],
})
router.beforeEach((to, from, next) => {
    // 判断有没有登录
    if (!localStorage.getItem("konohaToken")) {
        if (to.name == "login" || to.name == "register") {
            next()
        } else {
            router.push("login")
        }
    } else {
        next()
    }
})

export default router
