// import './assets/main.css'
import './assets/index.css'
import './assets/unreset.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'

import { OhVueIcon, addIcons } from "oh-vue-icons";
import * as Icons from "oh-vue-icons/icons";
const IconsSet = Object.values({ ...Icons });
addIcons(...IconsSet);

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.component("v-icon", OhVueIcon);

app.mount('#app')
