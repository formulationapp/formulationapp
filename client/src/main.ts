// import './assets/main.css'
import './assets/index.css'
import './assets/unreset.css'
import './utils'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'

import {addIcons, OhVueIcon} from "oh-vue-icons";
import {
    BiArrowDown,
    BiArrowUp,
    BiCardChecklist,
    BiHr,
    BiQuote,
    BiTextLeft,
    BiTypeH1,
    BiTypeH2,
    BiTypeH3,
    HiPlus,
    HiTrash,
    IoTextOutline,
    MdDragindicator
} from "oh-vue-icons/icons";
import {registerBlock} from "@dashibase/lotion";
import ShortAnswer from "@/lotion/ShortAnswer.vue";
import Choice from "@/lotion/Choice.vue";

addIcons(HiTrash, HiPlus,
    MdDragindicator, BiTextLeft, BiTypeH1, BiTypeH2, BiTypeH3, BiHr,
    BiQuote, IoTextOutline, BiCardChecklist, BiArrowDown, BiArrowUp);

registerBlock('SHORT_ANSWER', 'Short answer', ShortAnswer, 'io-text-outline')
registerBlock('CHOICE', 'Choice', Choice, 'bi-card-checklist')

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.component("v-icon", OhVueIcon);

app.mount('#app')
