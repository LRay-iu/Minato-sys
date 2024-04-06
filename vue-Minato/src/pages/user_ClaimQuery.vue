<template>
    <div>
        <PageTop />
        <el-container class="Maincontent">
            <ClaimSide />
            <el-container direction="vertical" class="content">
                <el-table :data="filterTableData" style="width: 100%;font-size: 15px;" size="large">
                    <el-table-column type="expand">
                        <template #default="props">
                            <div style="display: flex;justify-content: center;">
                                <el-table :data="[props.row]" height:15px size="small" style="width:65%" :border="true">
                                    <el-table-column label="用户名" prop="username" />
                                    <el-table-column label="身份证号" prop="userid" />
                                    <el-table-column label="联系方式" prop="callnumber" />
                                    <el-table-column label="保险单号" prop="insuranceid" />
                                    <el-table-column label="车牌号" prop="carid" />
                                    <el-table-column label="赔偿金额" prop="compensation" />
                                </el-table>
                            </div>
                            <!-- 步骤条 -->
                            <div style="display: flex;justify-content: center;width:100%;margin: 20px;">
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
                            <el-input v-model="search" class="search" size="default" placeholder="Type to search"
                                :prefix-icon="Search" />
                        </template>
                        <template #default="scope">
                            <el-button @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
                            <el-button type="danger" @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
                        </template>
                    </el-table-column>
                </el-table>
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
import { computed, ref } from 'vue'
import { type ClaimResult } from "@/interface/claimForm"
// interface User {
//     date: string
//     name: string
//     address: string
// }

const search = ref('')
const filterTableData = computed(() =>
    tableData.filter(
        (data) =>
            !search.value ||
            data.claimid.toLowerCase().includes(search.value.toLowerCase())
    )
)
const handleEdit = (index: number, row: ClaimResult) => {
    console.log(index, row)
}
const handleDelete = (index: number, row: ClaimResult) => {
    console.log(index, row)
}

const tableData: ClaimResult[] = [
    {
        claimid: 'CLM001',
        username: 'Tom',
        userid: '001',
        callnumber: '123456789',
        insuranceid: 1,
        carid: 'CAR001',
        region: 'New York',
        date: '2016-05-03',
        status: 4,
        compensation: 5000,
    },
    {
        claimid: 'CLM002',
        username: 'John',
        userid: '002',
        callnumber: '987654321',
        insuranceid: 2,
        carid: 'CAR002',
        region: 'Chicago',
        date: '2016-05-02',
        status: 3,
        compensation: 770,
    },
    {
        claimid: 'CLM003',
        username: 'Morgan',
        userid: '003',
        callnumber: '456789123',
        insuranceid: 3,
        carid: 'CAR003',
        region: 'San Francisco',
        date: '2016-05-04',
        status: 2,
        compensation: 660,
    },
    {
        claimid: 'CLM004',
        username: 'Jessy',
        userid: '004',
        callnumber: '789123456',
        insuranceid: 4,
        carid: 'CAR004',
        region: 'Los Angeles',
        date: '2016-05-01',
        status: 1,
        compensation: 7500,
    },
];



</script>

<style scoped>
.Maincontent {
    height: 100vh;
}

.content {
    padding: 20px 50px 70px 50px;
    /* border: 2px blue solid; */
}


.search {
    width: 300px;
    margin-right: 5px;
    /* --el-input-focus-border: #ff9900 !important; */
}
</style>