<template>
    <!------------------------------------- 头部 ------------------------------------->
    <div class="header-box">
        <el-space direction="horizontal" alignment="start" style="margin-left: 15px;">
            <div>
                您好，欢迎来到水门车险！
            </div>
            <el-space direction="horizontal" alignment="start" :size="15">
                <el-link :underline="false" @click="Login" style="margin-left: 25px;">登录</el-link>
                <div style="margin-left: 5px;">
                    |
                </div>
                <el-link :underline="false" @click="Register" style="margin-left: 15px;">注册</el-link>
            </el-space>
        </el-space>
        <el-space direction="horizontal" alignment="start" :size="20">
            <div>如有意外,联系作者</div>
            <div>vx : LRay-iu</div>
        </el-space>
    </div>
    <!------------------------------------- 导航栏 ------------------------------------->
    <el-menu v-bind="menuProps" class="menu-wrapper">
        <div class="logo">
            <img src="../assets/logo/minato_logo2.png" alt="" style="width:160px;height:60px">
        </div>
        <div class="menu-items">
            <el-menu-item index="1">首页</el-menu-item>
            <el-sub-menu index="2">
                <template #title>服务中心</template>
                <el-menu-item index="/search">保单查询 </el-menu-item>
                <el-menu-item index="2-2">订单支付</el-menu-item>
                <el-menu-item index="2-3">一键投诉</el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="3">
                <template #title>理赔中心</template>
                <el-menu-item index="3-1">车险在线理赔</el-menu-item>
                <el-menu-item index="3-2">理赔查询</el-menu-item>
                <el-menu-item index="3-3">理赔资料下载</el-menu-item>
            </el-sub-menu>
            <el-menu-item index="4">事故上报</el-menu-item>
            <el-menu-item index="5">车友圈</el-menu-item>
            <div style="padding-top: 10px;margin-left: 80px;margin-right: 25px;">
                <img src="../assets/logo/SUES.png" alt="" style="width:220px;height:85px">
            </div>
        </div>
    </el-menu>
    <!------------------------------------- 正文 ------------------------------------->
    <el-container class="Maincontent" direction="vertical">
        <el-container direction="horizontal">
            <RouterView></RouterView>
        </el-container>
    </el-container>
</template>

<script lang="ts" setup name="Base">
import { ref } from 'vue'
import { RouterView } from 'vue-router';
const activeIndex = ref('1');
const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
// el-menu的属性定义
const menuProps = {
    defaultActive: activeIndex.value,
    class: 'el-menu-demo custom-menu',
    mode: 'horizontal',
    uniqueOpened: true,
    router: true,
    backgroundColor: '#FF9900',
    textColor: 'rgba(255,255,255,0.75)',
    activeTextColor: '#fff',
    onSelect: handleSelect
}

function Login() {
    console.log("login")
}
function Register() {
    console.log("Register")
}
</script>

<style>
@import '../assets/main.css';

.flex-grow {
    flex-grow: 1;
}

.header-box {
    /* height: 30px; */
    /* border: 1px red solid; */
    margin: 0 200px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: rgba(0, 0, 0, 0.55);
}

/* .header-box .top-bar {
    color: rgba(0, 0, 0, 0.55);
} */

.el-menu-demo {
    padding-left: 100px;
    margin-bottom: 10px;
}

.custom-menu,
.el-menu-item,
.el-sub-menu__title {
    font-size: 18px;
    /* 设置自定义菜单的字体大小 */
}

.menu-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
    /* 居中 */
}

.menu-wrapper .logo {
    padding-top: 8px;
    margin-left: 100px;
    margin-right: 30px;
}

.menu-wrapper .menu-items {
    display: flex;
    align-items: center;
    margin-right: 180px;
    /* 居中 */
}

/* 正文 */
.Maincontent {
    /* border: 1px solid red; */
    justify-content: center;
    align-items: center;
    margin-top: 10px;
}
</style>