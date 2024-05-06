<template>
    <div>
        <PageTop />
        <el-container class="Maincontent">
            <InsuranceSide />
            <el-container direction="vertical" class="content">
                <el-table :data="filterTableData" style="width: 100%; font-size: 15px" size="large">
                    <el-table-column type="expand">
                        <template #default="props">
                            <div style="display: flex; justify-content: center">
                                <el-table :data="[props.row]" height:15px size="small" style="width: 65%"
                                    :border="true">
                                    <el-table-column label="用户名" prop="username" />
                                    <el-table-column label="联系方式" prop="callnumber" />
                                    <el-table-column label="保险单号" prop="insuranceid" />
                                    <el-table-column label="车牌号" prop="carid" />
                                    <el-table-column label="赔偿金额" prop="compensation" />
                                </el-table>
                            </div>
                            <div style="
                                    margin: 30px 0 0 200px;
                                    font-size: smaller;
                                ">
                                交易地址：{{ props.row.address }}
                            </div>
                            <!-- 步骤条 -->
                            <div style="
                                    display: flex;
                                    justify-content: center;
                                    width: 100%;
                                    margin: 20px;
                                ">
                                <el-steps style="width: 600px" :space="200" :active="props.row.status"
                                    finish-status="success" process-status="wait" align-center="true">
                                    <el-step title="发起申请&nbsp;" />
                                    <el-step title="审批&nbsp;" />
                                    <el-step title="创建订单&nbsp;" />
                                    <el-step title="钱款到账&nbsp;" />
                                </el-steps>
                            </div>
                            <!--  -->
                        </template>
                    </el-table-column>
                    <el-table-column label="日期" prop="date" />
                    <el-table-column label="理赔单号" prop="claimid" />
                    <el-table-column label="地区" prop="region" />
                    <el-table-column align="right">
                        <template #header>
                            <el-input v-model="search" class="search" size="default" placeholder="输入理赔单号"
                                :prefix-icon="Search" />
                        </template>
                    </el-table-column>
                </el-table>
            </el-container>
        </el-container>
        <PageBotton />
    </div>
</template>

<script lang="ts" setup name="adminSearch">
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import InsuranceSide from "@/components/InsuranceSide.vue"
import { Search } from "@element-plus/icons-vue"
import { computed } from "vue"

const search = ref("")

//-------------------------------------向后端访问数据--------------------------------
import { ref, onBeforeMount } from "vue"
import axios from "axios"
import { ElMessage } from "element-plus"
import { type ClaimResult } from "@/interface/claimForm"

const tableData = ref<ClaimResult[]>([])
const filterTableData = computed(() =>
    tableData.value.filter(
        (data) =>
            !search.value ||
            data.claimid.toLowerCase().includes(search.value.toLowerCase()),
    ),
)
onBeforeMount(async () => {
    fetchData()
    setTimeout(() => { }, 500)
})

async function fetchData() {
    try {
        const response = await axios.get(
            "http://localhost:8080/showallrecords",
        )
        switch (response.data.code) {
            case 200:
                tableData.value = response.data.data
                setTimeout(() => {
                    console.log("tabledata怎么你了", tableData.value)
                }, 5000)
                break
            case 3003:
                ElMessage({
                    message: "申报查询失败！",
                    type: "error",
                })
                break
            default:
                break
        }
    } catch (error) {
        console.error("query failed:", error)
        // 查询失败的处理
    }
}
</script>

<style scoped>
.Maincontent {
    height: 100vh;
    width: 100%;
}

.content {
    padding: 20px 50px 70px 50px;
    min-height: 100vh;
    /* border: 2px blue solid; */
}
</style>
