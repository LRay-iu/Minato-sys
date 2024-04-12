<template>
    <div>
        <div>
            <PageTop />
            <el-container>
                <ClaimSide />
                <el-container direction="vertical" class="content">
                    <!-- 理赔审批页面 -->
                    <el-descriptions v-if="tableData" class="margin-top" :title="'理赔单号：&nbsp;' + tableData.claimid"
                        :column="3" size="large" border>
                        <template #extra>
                            <!-- <el-button type="primary">Operation</el-button> -->
                        </template>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <user />
                                    </el-icon>
                                    理赔用户
                                </div>
                            </template>
                            {{ tableData.username }}
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <iphone />
                                    </el-icon>
                                    联系方式
                                </div>
                            </template>
                            {{ tableData.callnumber }}
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <location />
                                    </el-icon>
                                    申报地区
                                </div>
                            </template>
                            {{ tableData.region }}
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <tickets />
                                    </el-icon>
                                    状态
                                </div>
                            </template>
                            <el-tag effect="light" size="large">待办</el-tag>
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <office-building />
                                    </el-icon>
                                    申诉车辆
                                </div>
                            </template>
                            {{ tableData.carid }}
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <office-building />
                                    </el-icon>
                                    保险单号
                                </div>
                            </template>
                            {{ tableData.insuranceid }}
                        </el-descriptions-item>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <office-building />
                                    </el-icon>
                                    日期
                                </div>
                            </template>
                            {{ tableData.date }}
                        </el-descriptions-item>
                    </el-descriptions>
                    <el-descriptions v-if="tableData" class="margin-top" direction="vertical" :column="1" size="large"
                        border>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <office-building />
                                    </el-icon>
                                    申诉文件
                                </div>
                            </template>
                            <el-container>
                                <div v-for="(item, index) in tableFile" :key="index" style="margin: 10px">
                                    <el-image style="
                                            width: 150px;
                                            height: 150px;
                                            border-radius: 10px;
                                        " :src="'data:image/jpeg;base64,' +
                                            item.file
                                            " :zoom-rate="1.2" :max-scale="7" :min-scale="0.2" :preview-src-list="[
                                                'data:image/jpeg;base64,' +
                                                item.file,
                                            ]" :initial-index="4" fit="cover" />
                                </div>
                            </el-container>
                        </el-descriptions-item>
                    </el-descriptions>
                    <el-descriptions v-if="tableData" class="margin-top" direction="vertical" :column="1" size="large"
                        border>
                        <el-descriptions-item>
                            <template #label>
                                <div class="cell-item">
                                    <el-icon>
                                        <office-building />
                                    </el-icon>
                                    审批意见
                                </div>
                            </template>
                            <el-input v-model="note" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                                placeholder="批注" />
                            <el-container> </el-container>
                        </el-descriptions-item>
                    </el-descriptions>
                    <div  v-if="tableData" style="
                             display:flex ;
                            justify-content: space-between;
                            align-items: center;
                            height: 150px;
                            padding:0 400px
                        ">
                        <el-button type="primary" size="large" @click="open">通过</el-button>
                        <el-button type="danger" size="large">退回</el-button>
                    </div>
                </el-container>
            </el-container>
            <PageBotton />
        </div>
    </div>
</template>

<script lang="ts" setup name="adminClaimDetail">
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"
import { useRoute } from "vue-router" //1.先在需要跳转的页面引入useRouter
const { params } = useRoute()
import { ref, onBeforeMount } from "vue"
// import { type ClaimResult } from "@/interface/claimForm";
//-----------------装饰----------------------
import {
    Iphone,
    Location,
    OfficeBuilding,
    Tickets,
    User,
} from "@element-plus/icons-vue"
//------------获取数据----------------------
let note = ref("")
const tableData = ref(null)
const tableFile = ref(null)
// 向后端请求数据
import axios from "axios"
onBeforeMount(() => {
    fetchData()
    fetchFile()
    setTimeout(() => { }, 500)
})
async function fetchData() {
    await axios
        .get("http://localhost:8080/showclaimdetail/" + params.claimid)
        .then((response) => {
            console.log(response.data.data)
            switch (response.data.code) {
                case 200:
                    tableData.value = response.data.data
                    setTimeout(() => {
                        // console.log("tabledata怎么你了", tableData.value)
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
        })
}
async function fetchFile() {
    await axios
        .get("http://localhost:8080/showclaimfile/" + params.claimid)
        .then((response) => {
            console.log(response.data.data)
            switch (response.data.code) {
                case 200:
                    tableFile.value = response.data.data
                    setTimeout(() => {
                        // console.log("tableFile怎么你了", tableFile.value)
                    }, 500)
                    break
                case 3003:
                    ElMessage({
                        message: "申报文件查询失败！",
                        type: "error",
                    })
                    break
                default:
                    break
            }
        })
}
//------------------------------后续，提交交易表单---------------------------------
import { ElMessage, ElMessageBox } from 'element-plus'
import {compensation} from '@/scripts/Minatosys'
import {updateCompensation} from '@/scripts/claim_helper'
import { useRouter } from "vue-router"
const router = useRouter()
const open = () => {
    ElMessageBox.prompt('请输入赔偿金额（元）', '确认', {
        confirmButtonText: '交易执行',
        cancelButtonText: '取消',
        inputErrorMessage: 'Invalid Email',
    })
        .then(({ value }) => {
            const updatePromise = updateCompensation(parseFloat(value), tableData.value.claimid);
            const compenstationPromise = compensation(tableData.value.publicKey, BigInt(value), tableData.value.claimid);
            ElMessage({
                type: 'success',
                message: `正在建立交易，请稍后`,
            });
            return Promise.all([updatePromise, compenstationPromise]);
        })
        .then(() => {
            setTimeout(() => {
                router.push("/3-4");
            }, 3000);
        })
        .catch(() => {
            ElMessage({
                type: 'error',
                message: '交易取消',
            })
        })
}
</script>

<style scoped>
.content {
    padding: 20px 50px 70px 50px;
    min-height: 100vh;
    /* border: 2px blue solid; */
}

.cell-item {
    display: flex;
    align-items: center;
}
</style>
