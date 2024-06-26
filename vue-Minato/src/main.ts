import { createApp } from "vue"

import ElementPlus from "element-plus"
import "element-plus/dist/index.css"

import * as ElementPlusIconsVue from "@element-plus/icons-vue"
import router from "./router"
import App from "./App.vue"
const app = createApp(App)
import { createPinia } from "pinia"
const pinia = createPinia()
app.use(ElementPlus)
app.use(router)
app.mount("#app")
//安装pinia
app.use(pinia)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
