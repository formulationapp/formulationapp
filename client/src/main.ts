import './assets/main.css'
import './assets/index.css'
import './assets/unreset.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'
import {useAuth} from "@/stores/auth";

const app = createApp(App)
app.use(createPinia())
app.use(router)

const auth = useAuth();
auth.try();

app.mount('#app')
