import {defineStore} from "pinia";
import {api} from "@/api";
import {capitalize} from "@/utils";
import jwtDecode from "jwt-decode";
import {useWorkspaces} from "@/stores/workspaces";
import type User from "@/models/user";

interface AuthState {
    user: User
}

export const useAuth = defineStore('auth', {
    state: (): AuthState => <AuthState>({
        user: {}
    }),
    actions: {
        async try() {
            if (localStorage.getItem('token') !== null) {
                await this.useToken(localStorage.getItem('token'));
            }
        },
        async login(email: string, password: string) {
            if (email.length == 0 || password.length == 0) return;
            try {
                const response = await api.post('auth/login', {email, password}) as { token: string };
                await this.useToken(response.token);
            } catch (err: any) {
                throw new Error(capitalize(err.body.message));
            }
        },
        async register(email: string, password: string) {
            if (email.length == 0 || password.length == 0) return;
            try {
                const response = await api.post('auth/register', {email, password}) as { token: string };
                await this.useToken(response.token);
            } catch (err: any) {
                throw new Error(capitalize(err.body.message));
            }
        },
        async useToken(token: string) {
            localStorage.setItem('token', token);
            this.user = jwtDecode(token);

            api.options.headers.Authorization = 'Bearer ' + token;

            const workspaces = useWorkspaces();
            await workspaces.load();
        },
        logout() {
            localStorage.removeItem('token');
        }
    }
});