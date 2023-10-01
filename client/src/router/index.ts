import {createRouter, createWebHistory} from 'vue-router'
import Dashboard from "@/views/Dashboard/Dashboard.vue";
import Forms from "@/views/Dashboard/Forms/Forms.vue";
import FormsMenu from "@/views/Dashboard/Forms/FormsMenu.vue";
import FormMenu from "@/views/Dashboard/Form/FormMenu.vue";
import Submissions from "@/views/Dashboard/Form/FormSubmissions/Submissions.vue";
import FormSettings from "@/views/Dashboard/Form/FormSettings/FormSettings.vue";
import HomeView from "@/views/HomeView.vue";
import LoginView from "@/views/Auth/Login/LoginView.vue";
import RegisterView from "@/views/Auth/Register/RegisterView.vue";
import FormBox from "@/views/Sharing/FormBox/FormBox.vue";
import FormEditor from "@/views/Dashboard/Form/FormEditor/FormEditor.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: LoginView
        },
        {
            path: '/register',
            name: 'register',
            component: RegisterView
        },
        {
            path: '/',
            component: HomeView,
        },
        {
            path: '/workspaces/:workspaceID',
            name: 'forms',
            component: Dashboard,
            children: [
                {
                    path: '',
                    components: {
                        default: Forms,
                        menu: FormsMenu,
                    },
                },
            ],
        },
        {
            path: '/workspaces/:workspaceID/forms/:formID',
            name: 'form',
            component: Dashboard,
            children: [
                {
                    path: '',
                    name:'design',
                    components: {
                        default: FormEditor,
                        menu: FormMenu,
                    },
                },
                {
                    path: 'submissions',
                    name:'submissions',
                    components: {
                        default: Submissions,
                        menu: FormMenu,
                    },
                },
                {
                    path: 'settings',
                    name:'form_settings',
                    components: {
                        default: FormSettings,
                        menu: FormMenu,
                    },
                },
            ],
        },
        {
            path: '/f/:secret',
            name: 'submit',
            component: FormBox
        }
    ]
})

export default router
