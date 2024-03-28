<template>

    <!------------------------------------- 头部 ------------------------------------->
    <el-container style="justify-content: center;">
        <div class="header-box">
            <el-space direction="horizontal" alignment="start">
                <div>
                    您好，欢迎来到水门车险！
                </div>
                <el-space direction="horizontal" alignment="start" :size="15">
                    <el-link :underline="false" @click="Login" style="margin-left: 25px;font-size: 15px;">登录</el-link>
                    <div style="margin-left: 5px;">
                        |
                    </div>
                    <el-link :underline="false" @click="Register"
                        style="margin-left: 15px;font-size: 15px;">注册</el-link>
                </el-space>
            </el-space>
            <el-space direction="horizontal" alignment="start" :size="20">
                <div>如有意外,联系作者</div>
                <div>vx : LRay-iu</div>
            </el-space>
        </div>
    </el-container>
    <!------------------------------------- 导航栏 ------------------------------------->
    <el-menu v-bind="menuProps" class="menu-wrapper " :popper-offset="5" mode="horizontal"
        :default-active="$route.path">
        <div style="margin-right: 350px;">
            <img src="../assets/logo/minato_logo2.png" alt="" style="width:160px;height:60px">
        </div>
        <el-menu-item index="/">首页</el-menu-item>
        <el-sub-menu index="2">
            <template #title>
                <el-container style="height: 78px;align-items: center;">服务中心</el-container>
            </template>
            <el-menu-item index="/2-1">
                <el-icon>
                    <ShoppingCart />
                </el-icon>
                <span>&nbsp; 车险购买</span>
            </el-menu-item>
            <el-menu-item index="/2-2">
                <el-icon>
                    <Search />
                </el-icon>
                <span>&nbsp; 保单查询</span>
            </el-menu-item>
            <el-menu-item index="/2-3">
                <el-icon>
                    <Message />
                </el-icon>
                <span>&nbsp; 一键投诉</span>
            </el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="3">
            <template #title>
                <el-container style="height: 78px;align-items: center;">理赔中心</el-container>
            </template>
            <el-menu-item index="/3-1">
                <el-icon>
                    <Coin />
                </el-icon>
                <span>&nbsp; 车险在线理赔</span>
            </el-menu-item>
            <el-menu-item index="/3-2">
                <el-icon>
                    <Search />
                </el-icon>
                <span>&nbsp; 理赔查询</span>
            </el-menu-item>
            <el-menu-item index="/3-3">
                <el-icon>
                    <Download />
                </el-icon>
                <span>&nbsp; 理赔资料下载</span>
            </el-menu-item>
        </el-sub-menu>

        <el-menu-item index="/4">车友圈</el-menu-item>
        <div style="margin-left: 80px;margin-right: 25px;">
            <img src="../assets/logo/SUES.png" alt="" style="width:220px;height:78px" @click="SUES" />
        </div>
    </el-menu>

</template>

<script lang='ts' setup>
import { useRouter } from 'vue-router';
const router = useRouter()

// el-menu的属性定义
const menuProps = {
    mode: 'horizontal',
    uniqueOpened: true,
    router: true,
}

function Login() {
    router.push('/login')
}
function Register() {
    router.push('/register')
}

function SUES() {
    // window.open("https://www.sues.edu.cn/");//新窗口打开
    window.location.href = "https://www.sues.edu.cn/";//当前窗口打开
}
</script>

<style scoped>
@import '../assets/main.css';

.flex-grow {
    flex-grow: 1;
}

.header-box {
    height: 30px;
    width: 1100px;
    /* border: 1px red solid; */
    /* margin-left: 300px; */
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: rgba(0, 0, 0, 0.55);
    font-size: 15px;
}


.menu-items,
.el-sub-menu__title {
    height: 100px;
}


.el-sub-menu__title :hover {
    color: #ff9900 !important;
}

.menu-wrapper {
    display: flex;
    /* border-top: 1px solid rgba(0, 0, 0, 0.2); */
    border-bottom: 0;
    /* border: 1px red solid; */
    height: 80px;
    align-items: center;
    justify-content: center;
    margin-bottom: 5px;
    box-shadow: 0px 2px 10px 1px rgba(0, 0, 0, 0.1);
    /* padding-left: 100px; */
}

/* .menu-wrapper .logo {
    padding-top: 8px;
    /* margin-left: 200px; 
    margin-right: 350px;
} */


.menu-wrapper .menu-items {
    display: flex;
    align-items: center;
    /* margin-right: 180px; */
    /* 居中 */
}
</style>