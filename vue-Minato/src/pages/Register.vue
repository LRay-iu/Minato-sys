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
                        <div style="color:rgba(0,0,0,0.75);font-size: 15px;font-weight: 600;">
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
                        <el-input v-model="publicKey" placeholder="请输入Metamask账户公钥" clearable @focus="autoFill"/>
                        <el-button color="#0060df" @click="register">
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
let checkTime = true
//--------------------------------------------响应式组件数据获取--------------------------------------
let username = ref('')
let password = ref('')
let publicKey = ref('')
//----------------------------------------------编程式路由导航----------------------------------------
import { useRouter } from 'vue-router';
const router = useRouter()
function login() {
    router.push('/login')
    // console.log(password)
}
//----------------------------------------------------方法---------------------------------------------
// This function detects most providers injected at window.ethereum.
async function autoFill() {
    if (checkTime) {
        connect()
        fund()
        checkTime = false
    } else {
        fund()
    }
}

function register(){
    console.log({username,password})
}
//-----------------------------------------------------------用以检查浏览器有没有安装Metamask扩展----------------------------------------
async function connect() {
    if (typeof (window as any).ethereum !== "undefined") {
        await (window as any).ethereum.request({ method: "eth_requestAccounts" })
    } else {
        console.log("No metamask!")
    }
}
import { ethers } from '@/scripts/ethers-5.7.esm.min.js';
async function fund() {
    // console.log(ethAccount)
    if (typeof (window as any).ethereum !== "undefined") {
        //连接到区块链
        //一个具有gas的人/钱包。用于实际发送交易
        //能与之交互的合约
        //合约的ABI
        const provider = new ethers.providers.Web3Provider((window as any).ethereum) //从metamask找到了http端点，将其作为provider
        const signer = provider.getSigner()//返回metamask当前连接的钱包是哪一个
        const address =await signer.getAddress( ) 
        console.log(address)
        publicKey.value =address.toString(); 
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
