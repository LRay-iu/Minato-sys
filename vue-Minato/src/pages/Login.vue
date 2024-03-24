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
                        <el-input v-model="username" placeholder="请输入用户名" clearable />
                        <el-input v-model="password" type="password" placeholder="请输入密码" clearable show-password />
                        <el-button color="#0060df" dark="#03459d">
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
const username = ref('')
const password = ref('')
import { useRouter } from 'vue-router';
const router = useRouter()
function register() {
    router.push('/register')
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
