<template>
    <div>
        <PageTop />
        <el-container>
            <ClaimSide />
            <el-container style="justify-content: center;padding-top: 7vh;">
                <!-- 理赔页面 -->
                <el-form ref="claimFormRef" style="max-width: 1500px;min-width: 50vw" :model="claimForm" :rules="rules"
                    label-width="auto" class="demo-ruleForm  Maincontent" :size="formSize" status-icon="false"
                    label-position="top">
                    <!-- 索赔人姓名 -->
                    <el-form-item label="用户姓名" prop="username">
                        <el-input v-model="claimForm.username" placeholder="请输入真实姓名" />
                    </el-form-item>
                    <!-- 索赔人身份证号 -->
                    <el-form-item label="身份证号" prop="userid">
                        <el-input v-model="claimForm.userid" placeholder="请输入身份证号" />
                    </el-form-item>
                    <!-- 索赔人电话 -->
                    <el-form-item label="联系电话" prop="callnumber">
                        <el-input v-model="claimForm.callnumber" />
                    </el-form-item>
                    <!-- 索赔人车牌号 -->
                    <el-form-item label="车牌号" prop="carid">
                        <el-input v-model="claimForm.carid" />
                    </el-form-item>
                    <!-- 索赔人购置的保险单号 -->
                    <el-form-item label="保险单号" prop="insuranceid">
                        <el-input v-model="claimForm.insuranceid" />
                    </el-form-item>
                    <!--申请地区 -->
                    <el-form-item label="申请地区" prop="region">
                        <el-select-v2 v-model="claimForm.region" placeholder="请选择申请地区" :options="chineseCities" />
                    </el-form-item>
                    <el-form-item label="申报日期" required>
                        <!-- 选择日期-->
                        <el-col :span="11">
                            <el-form-item prop="date">
                                <el-date-picker v-model="claimForm.date" type="date" label="请输入日期" placeholder="请输入日期"
                                    style="width: 100%" />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                    <!-- 上传文件 -->
                    <el-form-item label="上传事故责任确认书、保险单复印件、驾驶证复印件以及维修费清单" required prop="claimfile">
                        <el-upload class="upload-demo" drag
                            action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15" multiple
                            style="width:50vw">
                            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                            <div class="el-upload__text">
                                将文件拖至此处或 <em>点击上传</em>
                            </div>
                        </el-upload>
                    </el-form-item>
                    <!-- 提交按钮 -->
                    <div style="display: flex;margin-top: 30px;justify-content: center;align-items: center;">
                        <el-form-item>
                            <el-button type="primary" @click="submitForm(claimFormRef)" size="large">
                                提交
                            </el-button>
                            <el-button @click="resetForm(claimFormRef)" size="large"
                                style="margin-left: 100px;">重置</el-button>
                        </el-form-item>
                    </div>
                </el-form>
            </el-container>
        </el-container>
        <PageBotton />
    </div>
</template>

<script lang='ts' setup name='userClaim'>
//页面组件
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"
//中国城市数组
import { chineseCities } from "@/scripts/chineseCities"
//身份证校验
import { checkID } from "@/scripts/Codecheck"
// 理赔test
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
//引入上传的图标
import { UploadFilled } from '@element-plus/icons-vue'
//引入ClaimForm的钩子
import { type ClaimForm } from '@/interface/claimForm'

const formSize = ref('default')
const claimFormRef = ref<FormInstance>()

const claimForm = reactive<ClaimForm>({
    username: '',
    userid: '',
    callnumber: '',
    insuranceid: '',
    carid: '',
    region: '',
    date: '',
    claimfile: ''
})


const rules = reactive<FormRules<ClaimForm>>({
    username: [
        { required: true, message: '请输入用户名！', trigger: 'blur' },
        { min: 2, max: 10, message: '请输入真实用户名！', trigger: 'blur' },
    ],
    userid: [
        { required: true, message: '请输入身份证号！', trigger: 'blur' },
        { validator: checkID, trigger: 'blur' },
    ],
    callnumber: [
        { required: true, message: '请输入电话号码！', trigger: 'blur' },
        { max: 11, message: '请输入有效电话号码', trigger: 'blur' },
    ],
    carid: [
        { required: true, message: '请输入有效车牌号！', trigger: 'blur' },
        { len: 7, message: '请输入有效车牌号', trigger: 'blur' },
    ],
    insuranceid: [
        { required: true, message: '请输入保险单号！', trigger: 'blur' },
        { len: 5, message: '请输入有效保险单号！', trigger: 'blur' },
    ],
    region: [
        {
            required: true,
            message: '请选择申报地区！',
            trigger: 'change',
        },
    ],
    date: [
        {
            type: 'date',
            required: true,
            message: '请填写申报时间',
            trigger: 'change',
        },
    ],
    claimfile: [
        {
            required: true,
            message: '请上传必要的文件！',
            trigger: 'change',
        },
    ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!', fields)
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>

<style scoped>
.Maincontent {
    margin-bottom: 100px;
}
</style>