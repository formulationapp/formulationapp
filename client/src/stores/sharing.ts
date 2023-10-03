import {defineStore} from "pinia";
import {api} from "@/api";
import type Form from "@/models/form";
import {v4} from "uuid";

interface SharingState {
    form: Form,
    answers: { [name: string]: string },
    submissionID: string,
    submitted: boolean,
    valid: { [name: string]: boolean },
}

export const useSharing = defineStore('sharing', {
    state: (): SharingState => ({
        form: {},
        answers: {},
        submissionID: '',
        submitted: false,
        valid: {}
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
            for (const block of this.form.data.blocks.blocks) {
                if (block.type == 'SHORT_ANSWER' || block.type == 'CHOICE') {
                    this.valid[block.id] = !block.details.required;
                } else {
                    this.valid[block.id] = true;
                }
            }
        },
        setAnswer(field: string, value: string, valid: bool) {
            this.answers[field] = value;
            this.valid[field] = valid;
        },
        setSubmitted(submitted: boolean) {
            this.submitted = submitted;
        },
        async sendAnswers() {
            for (const [field, value] of Object.entries(this.answers)) {
                await api.put('/forms/' + this.form.secret + '/answers/' + this.submissionID, {
                    field, value
                });
            }
            localStorage.removeItem('submissionID');
        }
    }
})