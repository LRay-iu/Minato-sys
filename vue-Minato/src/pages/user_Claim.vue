<template>
    <div>
        <PageTop />
        <el-container>
            <ClaimSide />
            <el-container style="justify-content: center; padding-top: 7vh">
                <!-- 理赔页面 -->
                <el-form ref="claimFormRef" style="max-width: 1500px; min-width: 50vw" :model="claimForm" :rules="rules"
                    label-width="auto" class="demo-ruleForm Maincontent" :size="formSize" status-icon="false"
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
                                    format="YYYY/MM/DD" value-format="YYYY-MM-DD" style="width: 100%" />
                            </el-form-item>
                        </el-col>
                    </el-form-item>
                    <!-- 上传文件 -->
                    <el-form-item label="上传事故责任确认书、保险单复印件、驾驶证复印件以及维修费清单" required prop="claimfile">
                        <el-upload ref="uploadRef" class="upload-demo" drag action="http://localhost:8080/update"
                            multiple style="width: 50vw" :auto-upload="false" list-type="picture"
                            :before-upload="beforeClaimfileUpload" @change="handleChange">
                            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                            <div class="el-upload__text">
                                将文件拖至此处或 <em>点击上传</em>
                            </div>
                        </el-upload>
                    </el-form-item>
                    <!-- 提交按钮 -->
                    <div style="
                            display: flex;
                            margin-top: 30px;
                            justify-content: center;
                            align-items: center;
                        ">
                        <el-form-item>
                            <el-button type="primary" @click="combinedClick" size="large">
                                提交
                            </el-button>
                            <el-button @click="resetForm(claimFormRef)" size="large"
                                style="margin-left: 100px">重置</el-button>
                        </el-form-item>
                    </div>
                </el-form>
            </el-container>
        </el-container>
        <PageBotton />
    </div>
</template>

<script lang="ts" setup name="userClaim">
//页面组件
import PageTop from "@/components/PageTop.vue"
import PageBotton from "@/components/PageBotton.vue"
import ClaimSide from "@/components/ClaimSide.vue"
//中国城市数组
import { chineseCities } from "@/scripts/chineseCities"
//身份证校验
import { checkID } from "@/scripts/Codecheck"
// 理赔test
import { reactive, ref } from "vue"
import type { FormInstance, FormRules } from "element-plus"
//引入上传的图标
import { UploadFilled } from "@element-plus/icons-vue"
//引入ClaimForm的钩子
import { type ClaimForm } from "@/interface/claimForm"
import { nanoid } from "nanoid";
const formSize = ref("default")
const claimFormRef = ref<FormInstance>()
// 初始化各项数据，其中insuranceid不论怎么写最终读取都是字符串格式
// 因此不必过多纠结，打包的时候将其转换成number类型即可
const claimForm = reactive<ClaimForm>({
    claimid: nanoid(),
    username: "",
    userid: "",
    callnumber: "",
    insuranceid: '',
    carid: "",
    region: "",
    date: "",
    claimfile: [""],
})
//---------------------------------------------------表单验证--------------------------------------------------
import { ElMessage } from "element-plus"
import type { UploadProps } from "element-plus"
const rules = reactive<FormRules<ClaimForm>>({
    username: [
        { required: true, message: "请输入用户名！", trigger: "blur" },
        { min: 2, max: 10, message: "请输入真实用户名！", trigger: "blur" },
    ],
    userid: [
        { required: true, message: "请输入身份证号！", trigger: "blur" },
        { validator: checkID, trigger: "blur" },
    ],
    callnumber: [
        { required: true, message: "请输入电话号码！", trigger: "blur" },
        { max: 11, message: "请输入有效电话号码", trigger: "blur" },
    ],
    carid: [
        { required: true, message: "请输入有效车牌号！", trigger: "blur" },
        { len: 7, message: "请输入有效车牌号", trigger: "blur" },
    ],
    insuranceid: [
        { required: true, message: "请输入保险单号！", trigger: "blur" },
    ],
    region: [
        {
            required: true,
            message: "请选择申报地区！",
            trigger: "change",
        },
    ],
    date: [
        {
            type: "date",
            required: true,
            message: "请填写申报时间",
            trigger: "change",
        },
    ],
    claimfile: [
        {
            required: true,
            message: "请上传必要的文件！",
            trigger: "change",
        },
    ],
})

const formcheck = async (
    formEl: FormInstance | undefined,
): Promise<boolean> => {
    if (!formEl) return false
    try {
        await formEl.validate()
        console.log("submit!")
        return Promise.resolve(true)
    } catch (error) {
        console.log("error submit!", error)
        return Promise.reject(false)
    }
}
const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
//上传前会对文件进行检查，只允许jpg格式的文件上传至后端
const beforeClaimfileUpload: UploadProps["beforeUpload"] = (rawFile) => {
    if (rawFile.type !== "image/jpeg") {
        ElMessage.error("图片仅支持jpeg格式")
        return false
    } else if (rawFile.size / 1024 / 1024 > 5) {
        ElMessage.error("图片不能超过5MB!")
        return false
    }
    return true
}
// 监听上传组件的 change 事件，只要一有文件被加进来就会触发
//目的是为了获取上传的文件名并就将其保存在claimForm.claimfile[]这个数组中
const handleChange = (file:any, fileList:any) => {
    // 获取第一个文件的文件名
    if (fileList.length > 0) {
        for (let i = 0; i < fileList.length; i++) {
            claimForm.claimfile[i] = fileList[i].name
            // console.log("上传文件为", claimForm.claimfile[i])
        }
    } else {
        claimForm.claimfile[0] = ''
    }
}
//--------------------------------------------------表单提交---------------------------------------------------
//上传文件
import type { UploadFile } from 'element-plus/lib/components/upload';
import type { UploadInstance } from "element-plus"
import axios from "axios";
import { useRouter } from 'vue-router';
const router = useRouter()
const uploadRef = ref<UploadInstance>()
//功能是对提交表单的补充，目的是上传文件并将上传的文件名，绑定的申报单号发送至后端
const submitUpload = () => {
    uploadRef.value!.submit()
    let fileMessage = {
        "filename": claimForm.claimfile,
        "file_according": claimForm.claimid,
        "createtime": claimForm.date
    }
    // console.log("修改表单内容为", fileMessage)
    setTimeout(async () => {
        await axios.post(
            "http://localhost:8080/updatesupport",
            fileMessage,
        )
    }, 200)

    // console.log("send successful:", response.data.data)


}
//通过验证后提交表单
async function submitClaimForm() {
    //打包表单内容
    let claimMessage = {
        claim_id: claimForm.claimid,
        claim_user: claimForm.userid,
        claim_insurance: parseInt(claimForm.insuranceid),
        callnumber: claimForm.callnumber,
        car_id: claimForm.carid,
        region: claimForm.region,
        createtime: claimForm.date,
        // claim_file: claimForm.claimfile,
    }
    // console.log("提交表单内容为", claimMessage)
    try {
        const response = await axios.post(
            "http://localhost:8080/addclaim",
            claimMessage,
        )
        // console.log("send successful:", response.data.data)
        switch (response.data.code) {
            case 200:
                ElMessage({
                    message: "申报成功！",
                    type: "success",
                })
                //3秒后跳转至登陆界面
                setTimeout(() => {
                    router.push("/3-2")
                }, 300)
                // 可以添加其他情况的处理
                break
            case 2001:
                ElMessage({
                    message: "申报失败，请联系管理员LRay-iu！",
                    type: "error",
                })
                claimFormRef.value?.resetFields
                break
            default:
                ElMessage({
                    message: "出现异常！联系管理员LRay-iu",
                    type: "error",
                })
                claimFormRef.value?.resetFields
                break
        }
    } catch (error) {
        console.error("Claim failed:", error)
        // 注册失败的处理
    }
}
const combinedClick = async () => {
    //这里先经过表单验证，通过后会像后端开始发送json报文
    const isValid = await formcheck(claimFormRef.value)
    if (isValid) {
        submitClaimForm();
        submitUpload()
    } else {
        return
    }
}
</script>

<style scoped>
.Maincontent {
    margin-bottom: 100px;
}
</style>
