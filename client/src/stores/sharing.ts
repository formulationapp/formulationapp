import {defineStore} from "pinia";
import {api} from "@/api";
import type Form from "@/models/form";

interface SharingState {
    form: Form
}

export const useSharing = defineStore('sharing', {
    state: (): SharingState => ({
        form: {},
    }),
    actions: {
        async loadForm(secret: string) {
            this.form = await api.get('/forms/' + secret);
        }
    }
})