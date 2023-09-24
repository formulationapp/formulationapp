import {defineStore} from "pinia";
import {api} from "@/api";
import type Form from "@/models/form";
import {v4} from "uuid";

interface SharingState {
    form: Form,
    answers: { [name: string]: string },
    submissionID: string
}

export const useSharing = defineStore('sharing', {
    state: (): SharingState => ({
        form: {},
        answers: {},
        submissionID: ''
    }),
    actions: {
        async loadForm(secret: string) {
            if (localStorage.getItem('submissionID') !== null) {
                this.submissionID = localStorage.getItem('submissionID')
            } else {
                this.submissionID = v4();
                localStorage.setItem('submissionID', this.submissionID)
            }
            this.form = await api.get('/forms/' + secret);
        },
        setAnswer(field: string, value: string) {
            this.answers[field] = value;
        },
        async sendAnswers() {
            for (const [field, value] of Object.entries(this.answers)) {
                console.log(field);
                console.log(value);
                await api.put('/forms/' + this.form.secret + '/answers/' + this.submissionID, {
                    field, value
                });
            }
        }
    }
})