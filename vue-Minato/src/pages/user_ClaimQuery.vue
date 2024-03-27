<template>
    <div>
        <PageTop />
        <el-container class="Maincontent">
            <ClaimSide />
            <el-container direction="vertical" class="content">
                <div style="display:flex;justify-content:right;height: 35px;">
                    <div>
                        <el-input v-model="input" class="search" placeholder="Type something" :prefix-icon="Search" />
                        <el-button type="primary" color="#ff9900"><span
                                style="font-weight: 600;font-size:16px;color: white;">搜索</span></el-button>
                    </div>
                </div>
                <div v-for="index in 10" :key="index">
                </div>
            </el-container>

        </el-container>
        <PageBotton />
    </div>
</template>

<script lang='ts' setup name='userClaimQuery'>
//引入组件
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"
import { Search } from '@element-plus/icons-vue'
import { ref } from 'vue'
let input = ref('')
</script>

<style scoped>
.Maincontent {
    height: 100vh;
}

.content {
    padding: 20px 50px 70px 50px;
    border: 2px blue solid;
}

.div {
    border: solid 1px red;
}

.search {
    width: 260px;
    margin-right: 5px;
    --el-input-focus-border: #ff9900 !important;
}
</style>