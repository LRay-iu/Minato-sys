<template>
    <div style="width:240px;box-shadow: 0px 2px 10px 1px rgba(0, 0, 0, 0.1);;">
        <h2 class="mb-2" style="display: flex;justify-content: center;margin: 20px;">
            理赔中心</h2>
        <el-menu mode="vertical" class="el-menu-vertical-demo" router="true" :default-active="$route.path"
            style="border: 0;padding-left: 15px;">
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
        </el-menu>
    </div>
</template>

<script lang='ts' setup name='ClaimSide'>

</script>

<style scoped></style>