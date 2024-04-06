<template>
    <div style="background-color: #fbfbfe">
        <el-container style="padding: 30px 0px 20px 15px">
            <el-image
                style="width: 180px; height: 45px; border-radius: 15px"
                src="/src/assets/logo/login_logo.png"
                fit="fill"
            />
        </el-container>
        <el-container class="loginPage">
            <el-card class="login">
                <el-header style="height: auto">
                    <el-container
                        direction="vertical"
                        style="
                            justify-content: center;
                            align-items: center;
                            margin-bottom: 15px;
                        "
                    >
                        <div
                            style="
                                color: rgba(0, 0, 0, 0.75);
                                font-size: 15px;
                                font-weight: 600;
                            "
                        >
                            水门车险&nbsp;&nbsp;账户
                        </div>
                        <div
                            style="
                                color: rgba(0, 0, 0, 0.75);
                                font-size: 25px;
                                font-weight: 700;
                                margin-bottom: 8px;
                            "
                        >
                            请输入账户和密码
                        </div>
                    </el-container>
                </el-header>
                <el-main style="width: 450px">
                    <el-container
                        style="justify-content: center; align-items: center"
                        direction="vertical"
                    >
                        <!-- --------------------------------------注册表单 ------------------------------------------->
                        <el-form
                            :rules="rules"
                            ref="registerFormRef"
                            :model="registerForm"
                        >
                            <el-form-item prop="userid">
                                <el-input
                                    v-model="registerForm.userid"
                                    placeholder="请输入身份证号"
                                    clearable
                                    @keydown.enter="handleEnter"
                                />
                            </el-form-item>
                            <el-form-item prop="username">
                                <el-input
                                    v-model="registerForm.username"
                                    placeholder="请输入用户名"
                                    clearable
                                    @keydown.enter="handleEnter"
                                />
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input
                                    v-model="registerForm.password"
                                    type="password"
                                    placeholder="请输入密码"
                                    clearable
                                    show-password
                                    @keydown.enter="handleEnter"
                                />
                            </el-form-item>
                            <el-form-item prop="callnumber">
                                <el-input
                                    v-model="registerForm.callnumber"
                                    placeholder="请输入手机号"
                                    clearable
                                    @keydown.enter="handleEnter"
                                />
                            </el-form-item>
                            <el-form-item prop="publickey">
                                <el-input
                                    v-model="registerForm.publickey"
                                    placeholder="请输入Metamask账户公钥"
                                    clearable
                                    @focus="autoFill"
                                    @keydown.enter="handleEnter"
                                />
                            </el-form-item>
                            <el-form-item
                                prop="role"
                                style="text-align: center"
                            >
                                <el-radio-group
                                    v-model="registerForm.role"
                                    style="
                                        width: 200px;
                                        margin: 0 auto;
                                        justify-content: space-between;
                                    "
                                >
                                    <el-radio value="user" size="large"
                                        >用户</el-radio
                                    >
                                    <el-radio value="admin" size="large"
                                        >管理员</el-radio
                                    >
                                </el-radio-group>
                            </el-form-item>
                        </el-form>
                        <!-- --------------------------------------------------------------------->
                        <el-button color="#0060df" @click="combinedClick">
                            <div style="font-weight: 600">注册</div>
                        </el-button>
                        <el-divider>
                            <div
                                style="font-size: 16; color: rgba(0, 0, 0, 0.4)"
                            >
                                或
                            </div>
                        </el-divider>
                        <el-button @click="login">
                            <div style="color: rgba(0, 0, 0, 0.6)">
                                已有帐户？点我登录
                            </div>
                        </el-button>
                    </el-container>
                </el-main>
                <el-footer>
                    <el-container
                        direction="vertical"
                        style="
                            align-items: center;
                            font-size: smaller;
                            color: rgb(158, 158, 158);
                        "
                    >
                        继续操作则表示，您同意我们的服务条款和隐私声明。
                    </el-container>
                </el-footer>
            </el-card>
        </el-container>
    </div>
</template>

<script lang="ts" setup name="Register">
import { ref, reactive } from "vue"
let checkTime = true
//---------------------响应式组件数据获取---------------------------
let registerForm = reactive({
    userid: "",
    username: "",
    password: "",
    callnumber: "",
    publickey: "",
    role: "",
})
//-------------------------编程式路由导航----------------------------
import { useRouter } from "vue-router"
const router = useRouter()
function login() {
    router.push("/login")
    // console.log(password)
}
//---------------------------------------表单验证---------------------
import type { FormInstance, FormRules } from "element-plus"
import { checkID } from "@/scripts/Codecheck"
const registerFormRef = ref<FormInstance>()
//表单验证规则
const rules = reactive<FormRules>({
    username: [
        { required: true, message: "请输入用户名！", trigger: "blur" },
        { min: 1, max: 10, message: "请输入真实用户名！", trigger: "blur" },
    ],
    userid: [
        { required: true, message: "请输入身份证号！", trigger: "blur" },
        { validator: checkID, trigger: "blur" },
    ],
    callnumber: [
        { required: true, message: "请输入电话号码！", trigger: "blur" },
        { max: 11, message: "请输入有效电话号码", trigger: "blur" },
    ],
    password: [
        { required: true, message: "请输入密码！", trigger: "blur" },
        { min: 7, message: "请输入安全系数高的密码", trigger: "blur" },
    ],
    publickey: [{ required: true, message: "请输入公钥！", trigger: "blur" }],
    role: [{ required: true, message: "请选择身份！", trigger: "blur" }],
})
//下面这段我也看不太懂，大概意思就是一键验证表单
// 反正粘过来直接用就行了
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
//---------------------------------------方法---------------------------------
// This function detects most providers injected at window.ethereum.
import axios from "axios"
import { ElMessage } from "element-plus"
//自动填充公钥
async function autoFill() {
    if (checkTime) {
        connect()
        fund()
        checkTime = false
    } else {
        fund()
    }
}
//像后端发送注册请求
async function register() {
    let registerMessage = {
        user_id: registerForm.userid,
        username: registerForm.username,
        password: registerForm.password,
        callnumber: registerForm.callnumber,
        publicKey: registerForm.publickey,
        role: registerForm.role,
    }
    try {
        const response = await axios.post(
            "http://localhost:8080/register",
            registerMessage,
        )
        console.log("send successful:", response.data.data)
        switch (response.data.code) {
            case 200:
                ElMessage({
                    message: "注册成功！即将返回登录界面",
                    type: "success",
                })
                //3秒后跳转至登陆界面
                setTimeout(() => {
                    router.push("/login")
                }, 3000)
                // 可以添加其他情况的处理
                break
            case 1007:
                ElMessage({
                    message: "用户已存在！",
                    type: "error",
                })
                registerFormRef.value?.resetFields
                break
            case 1008:
                ElMessage({
                    message: "注册失败！联系管理员LRay-iu",
                    type: "error",
                })
                registerFormRef.value?.resetFields
                break
            default:
                ElMessage({
                    message: "出现异常！联系管理员LRay-iu",
                    type: "error",
                })
                registerFormRef.value?.resetFields
                break
        }
    } catch (error) {
        console.error("Register failed:", error)
        // 注册失败的处理
    }
}
//组合式按键，先验证表单，再实现注册
const combinedClick = async () => {
    const isValid = await formcheck(registerFormRef.value)
    if (isValid) {
        register()
    } else {
        return
    }
}
const handleEnter = () => {
    // 按下回车键时执行登录操作
    combinedClick()
}
//----------------------------------用以检查浏览器有没有安装Metamask扩展---------------------------------
async function connect() {
    if (typeof (window as any).ethereum !== "undefined") {
        await (window as any).ethereum.request({
            method: "eth_requestAccounts",
        })
    } else {
        console.log("No metamask!")
    }
}
import { ethers } from "@/scripts/ethers-5.7.esm.min.js"
async function fund() {
    // console.log(ethAccount)
    if (typeof (window as any).ethereum !== "undefined") {
        //连接到区块链
        //一个具有gas的人/钱包。用于实际发送交易
        //能与之交互的合约
        //合约的ABI
        const provider = new ethers.providers.Web3Provider(
            (window as any).ethereum,
        ) //从metamask找到了http端点，将其作为provider
        const signer = provider.getSigner() //返回metamask当前连接的钱包是哪一个
        const address = await signer.getAddress()
        // console.log(address)
        registerForm.publickey = address.toString()
    }
}
</script>

<style scoped>
/* 背景 */
.loginPage {
    height: 90vh;
    /* display: flex; */
    justify-content: center;
    align-items: center;
}

.loginPage .login {
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 20px;
    background-color: rgba(255, 255, 255, 1);
    box-shadow: 0px 0px 0px 2px 14px rgba(0, 0, 0, 0.1);
    color: black;
}

.login .el-input {
    width: 414px;
    height: 42px;
    margin: 3px;
    font-size: large;
    border-radius: 6px;
    /* border: 1px solid rgb(158, 158, 158);
    --el-input-hover-border-color: rgb(158, 158, 158);
    --el-input-border-color: rgb(158, 158, 158);
    --el-border-color: rgb(158, 158, 158); */
}

.loginPage .login .el-button {
    width: 414px;
    height: 45px;
    margin: 8px;
    font-size: large !important;
    color: white;
    border-radius: 6px;
}

.loginPage .login .register {
    width: 330px;
    height: 45px;
    margin: 8px;
    font-size: large !important;
    color: rgba(0, 0, 0, 0.8);
    border-radius: 6px;
    border: 1px solid rgb(0, 0, 0);
}
</style>
