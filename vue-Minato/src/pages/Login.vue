<template>
    <div style="background-color: #fbfbfe;">
        <el-container style="padding:30px 0px 20px 15px">
            <el-image style="width: 180px; height: 45px;border-radius: 15px;" src="/src/assets/logo/login_logo.png"
                fit="fill" />
        </el-container>
        <el-container class="loginPage">
            <el-card class="login">
                <el-header style="height: auto;">
                    <el-container direction="vertical"
                        style=" justify-content: center;align-items: center;margin-bottom: 15px;">
                        <div
                            style="color:rgba(0,0,0,0.75);font-size: 30px;font-weight: 700;letter-spacing: 20px;margin-bottom: 8px;">
                            欢迎来到
                        </div>
                        <div style="color:rgba(0,0,0,0.75);font-size: 20px;font-weight: 600;letter-spacing: 15px;">
                            水门车险
                        </div>
                    </el-container>
                </el-header>
                <el-main style="width:450px;height:350px">
                    <el-container style=" justify-content: center;align-items: center;" direction="vertical">
                        <el-input v-model="userid" placeholder="请输入用户账号（身份证号）" clearable @keydown.enter="handleEnter" />
                        <el-input v-model="password" type="password" placeholder="请输入密码" clearable show-password
                            @keydown.enter="handleEnter" />
                        <el-button color="#0060df" dark="#03459d" @click="Login">
                            <div style="font-weight: 600;">登录</div>
                        </el-button>
                        <el-divider>
                            <div style="font-size: 16;color:rgba(0,0,0,0.4)">或</div>
                        </el-divider>
                        <el-button @click="register">
                            <div style="color:rgba(0,0,0,0.6)">注册</div>
                        </el-button>
                    </el-container>
                </el-main>
                <el-footer>
                    <el-container direction="vertical"
                        style="align-items: center;font-size: smaller;color:rgb(158, 158, 158);">
                        继续操作则表示，您同意我们的服务条款和隐私声明。
                    </el-container>
                </el-footer>
            </el-card>
        </el-container>
    </div>
</template>

<script lang='ts' setup name='Login'>
import { ref } from 'vue'
import axios from "axios";
const userid = ref('')
const password = ref('')
import { useRouter } from 'vue-router';
const router = useRouter()
function register() {
    router.push('/register')
}
//将登录信息发送至后端
import { useLoginStore } from '@/store/loginStore'
import { ElMessage } from 'element-plus'
const loginStore = useLoginStore()
async function Login() {
    let loginMessage = {
        user_id: userid.value,
        password: password.value
    }
    try {
        const response = await axios.post('http://localhost:8080/login', loginMessage);
        // console.log(loginMessage)
        console.log('send successful:', response.data.data);
        switch (response.data.code) {
            case 200:
                // 登录成功后的处理
                loginStore.Login(response.data.data)
                loginStore.$subscribe((mutate, state) => {
                    //mutate表示发生变化的内容
                    console.log('loginStore内的数据发生了变化', mutate, state)
                })
                ElMessage({
                    message: response.data.msg,
                    type: 'success',
                })
                //0.2秒后跳转至登陆界面
                setTimeout(() => {
                    router.push('/')
                }, 200);
                break
            default:
                //登陆失败或者异常时的处理
                ElMessage({
                    message: response.data.msg,
                    type: 'error',
                })
                userid.value = ''
                password.value = ""
                break
        }
    } catch (error) {
        console.error('Login failed:', error);
        // 登录失败的处理
    }
}
const handleEnter = () => {
    // 按下回车键时执行登录操作
    Login();
};
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
    box-shadow: 0px 0px 20px 1px rgba(0, 0, 0, 0.1);
    color: black;
}

.login .el-input {
    width: 414px;
    height: 45px;
    margin: 8px;
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
