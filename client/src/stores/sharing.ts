import {defineStore} from "pinia";
import {api} from "@/api";

export const useSharing = defineStore('sharing', {
    state: () => ({
        form: {},
    }),
    actions: {
        async loadForm(secret: string) {
            this.form = await api.get('/forms/' + secret);
        }
    }
})