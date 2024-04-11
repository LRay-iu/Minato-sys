<template>
    <div style="width:240px;box-shadow: 0px 2px 10px 1px rgba(0, 0, 0, 0.1);;">
       <template v-if="loginStore.role=='user'">
        <h2 class="mb-2" style="display: flex;justify-content: center;margin: 20px;">
            服务中心</h2>
        <el-menu mode="vertical" class="el-menu-vertical-demo" router="true" :default-active="$route.path"
            style="border: 0;padding-left:15px">
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
        </el-menu>
       </template>
       <template v-if="loginStore.role=='admin'">
        <h2 class="mb-2" style="display: flex;justify-content: center;margin: 20px;">
            数据中心</h2>
        <el-menu mode="vertical" class="el-menu-vertical-demo" router="true" :default-active="$route.path"
            style="border: 0;padding-left:15px">
            <el-menu-item index="/2-4">
                <el-icon>
                    <ShoppingCart />
                </el-icon>
                <span>&nbsp; 车险流水</span>
            </el-menu-item>
            <el-menu-item index="/2-5">
                <el-icon>
                    <Search />
                </el-icon>
                <span>&nbsp; 流程查询</span>
            </el-menu-item>
        </el-menu>
       </template>
    </div>
</template>

<script lang='ts' setup name=''>
import { useLoginStore } from '@/store/loginStore'
const loginStore = useLoginStore()
</script>

<style scoped></style>