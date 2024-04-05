//store/loginStore.ts
import { defineStore } from "pinia"
import { type LoginResult } from "@/interface/loginResult"
import { ref } from "vue"
export const useLoginStore = defineStore("login", function login() {
    //真正存数据的地方
    const userid = ref(localStorage.getItem("userid") || "")
    const username = ref(localStorage.getItem("username") || "")
    const publicKey = ref(localStorage.getItem("publicKey") || "")
    const konohaToken = ref(localStorage.getItem("konohaToken") || "")
    function Login(loginResult: LoginResult) {
        userid.value = loginResult.user_id
        username.value = loginResult.username
        publicKey.value = loginResult.publicKey
        konohaToken.value = loginResult.konohaToken
    }

    return { userid, username, publicKey, konohaToken, Login }
})
