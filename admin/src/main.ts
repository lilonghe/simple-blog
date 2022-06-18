import { createApp, } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import './index.css'
import MavonEditor from "mavon-editor";
import 'mavon-editor/dist/css/index.css'

const app = createApp(App);
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(MavonEditor)
app.mount('#app')