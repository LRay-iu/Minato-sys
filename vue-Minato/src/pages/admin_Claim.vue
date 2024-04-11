<template>
    <div>
        <div>
            <PageTop />
            <el-container>
                <ClaimSide />
                <el-container direction="vertical" class="content">
                    <!-- 理赔审批页面 -->
                    <el-table :data="filterTableData" style="width: 100%;font-size: 15px;" size="large">
                        <el-table-column label="日期" prop="date" />
                        <el-table-column label="理赔单号" prop="claimid" />
                        <el-table-column label="申请人" prop="username" />
                        <el-table-column label="地区" prop="region" />
                        <el-table-column align="right">
                            <template #header>
                                <el-input v-model="search" class="search" size="default" placeholder="输入理赔单号"
                                    :prefix-icon="Search" />
                            </template>
                            <template #default="scope">
                                <el-button type="warning" plain @click="goToDetail(scope.row.claimid)">待审</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-container>
            </el-container>
            <PageBotton />
        </div>
    </div>
</template>

<script lang='ts' setup name='adminClaim'>
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"
import { Search } from '@element-plus/icons-vue'
import { ref, onBeforeMount, computed } from "vue";
import { type ClaimResult } from "@/interface/claimForm";
import { ElMessage } from "element-plus";
//---------------------------------做一个路由跳转---------------------------------------
import { useRouter } from 'vue-router';

const router = useRouter();

function goToDetail(claimId:string) {
    router.push({name:"ClaimDetail", params: { claimid: claimId }});
}

const search = ref('')
const tableData = ref<ClaimResult[]>([]);
const filterTableData = computed(() =>
    tableData.value.filter(
        (data) =>
            !search.value ||
            data.claimid.toLowerCase().includes(search.value.toLowerCase())
    )
)
// 向后端请求数据
import axios from "axios";
onBeforeMount(() => {
    fetchData();
    setTimeout(() => { }, 500)
})
async function fetchData() {
    try {
        const response = await axios.get("http://localhost:8080/showclaims/")
        switch (response.data.code) {
            case 200:
                tableData.value = response.data.data;
                setTimeout(() => { console.log("tabledata怎么你了", tableData.value) }, 5000)
                break;
            case 3003:
                ElMessage({
                    message: "申报查询失败！",
                    type: "error",
                });
                break;
            default:
                break;
        }
    } catch (error) {
        console.error("query failed:", error);
        // 查询失败的处理
    }
} 
</script>

<style scoped>
.content {
    padding: 20px 50px 70px 50px;
    min-height: 100vh;
    /* border: 2px blue solid; */
}

</style>