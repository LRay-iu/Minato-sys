<template>
    <div style="background-color: #fbfbfe;">
        <el-container style="padding:30px 0px 20px 23px">
            <el-image style="width: 180px; height: 45px;border-radius: 15px;" src="/src/assets/logo/login_logo.png"
                fit="fill" />
        </el-container>
        <el-container class="loginPage">
            <el-card class="login">
                <el-header style="height: auto;">
                    <el-container direction="vertical"
                        style=" justify-content: center;align-items: center;margin-bottom: 15px;">
                        <div style="color:rgba(0,0,0,0.75);font-size: 15px;">
                            水门车险&nbsp;&nbsp;账户
                        </div>
                        <div style="color:rgba(0,0,0,0.75);font-size: 25px;font-weight: 700;margin-bottom: 8px;">
                            请输入账户和密码
                        </div>
                    </el-container>
                </el-header>
                <el-main style="width:450px;">
                    <el-container style=" justify-content: center;align-items: center;" direction="vertical">
                        <el-input v-model="username" placeholder="请输入用户名" clearable />
                        <el-input v-model="password" type="password" placeholder="请输入密码" clearable show-password />
                        <el-input v-model="primaryKey" type="password" placeholder="请输入Metamask账户私钥" @focus="check"
                            clearable show-password />
                        <el-button color="#0060df" dark="#03459d">
                            <div style="font-weight: 600;">注册</div>
                        </el-button>
                        <el-divider>
                            <div style="font-size: 16;color:rgba(0,0,0,0.4)">或</div>
                        </el-divider>
                        <el-button @click="login">
                            <div style="color:rgba(0,0,0,0.6)">已有帐户？点我登录</div>
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

<script lang='ts' setup name='Register'>
import { ref } from 'vue'
let checkTime = false
const username = ref('')
const password = ref('')
const primaryKey = ref('')
import { useRouter } from 'vue-router';
const router = useRouter()
function login() {
    router.push('/login')
}

function check() {
    if (!checkTime) {
        console.log("good")
        checkTime = true
    } else {
        return
    }
}
</script>

<style scoped>
/* 背景 */
.loginPage {
    height: 85vh;
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
