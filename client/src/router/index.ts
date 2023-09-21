import {createRouter, createWebHistory} from 'vue-router'
import Dashboard from "@/views/Dashboard/Dashboard.vue";
import Forms from "@/views/Forms/Forms.vue";
import FormsMenu from "@/views/Forms/FormsMenu.vue";
import FormDesign from "@/views/Form/FormDesign.vue";
import FormMenu from "@/views/Form/FormMenu.vue";
import Submissions from "@/views/Submissions/Submissions.vue";
import FormSettings from "@/views/FormSettings/FormSettings.vue";
import HomeView from "@/views/HomeView.vue";
import SubmissionView from "@/views/Submission/SubmissionView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/Login/LoginView.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('../views/Register/RegisterView.vue')
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
                    components: {
                        default: FormDesign,
                        menu: FormMenu,
                    },
                },
                {
                    path: 'submissions',
                    components: {
                        default: Submissions,
                        menu: FormMenu,
                    },
                },
                {
                    path: 'settings',
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
            component: SubmissionView
        }
    ]
})

export default router
