import {defineStore} from "pinia";
import {api} from "@/api";

export const useSubmissions = defineStore('submissions', {
    state: () => ({
        form: {}
    }),
    actions: {
        async loadForm(secret: string) {
            this.form = await api.get('/forms/' + secret);
        }
    }
})