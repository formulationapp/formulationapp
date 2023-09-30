import {defineStore} from "pinia";
import {api} from "@/api";
import type Form from "@/models/form";
import {v4 as uuidv4} from 'uuid';
import moment from "moment";

interface FormsState {
    forms: Form[]
    form: Form,
    submissions: object[]
}


export const useForms = defineStore('forms', {
    state: (): FormsState => ({
        forms: [],
        form: {
            ID: 0,
            workspaceID: 0,
            name: ''
        },
        submissions: {}
    }),
    actions: {
        async load(workspaceID: number) {
            this.forms = (await api.get('workspaces/' + workspaceID + '/forms')) ?? [];
            this.forms.sort((a, b) => (moment(a.CreatedAt).unix() < moment(b.CreatedAt).unix()) ? 1 : -1)
        },
        select(formID: number) {
            this.form = this.forms.find(x => x.ID == formID)!;
        },
        async save(form: Form) {
            await api.put('workspaces/' + form.workspaceID + '/forms/' + form.ID, form);
        },
        async setName(name: string) {
            this.form.name = name;
            await this.save(this.form);
        },
        async setBlocks(blocks: any) {
            this.form.data.blocks = blocks;
            await this.save(this.form);
        },
        async setSubmitLabel(label: string) {
            this.form.data.submit = label;
            await this.save(this.form);
        },
        async delete(workspaceID: number, formID: number) {
            await api.delete('workspaces/' + workspaceID + '/forms/' + formID);
            this.forms = this.forms.filter(form => form.ID != formID);
        },
        async create(workspaceID: number, name: string) {
            const form = await api.post('workspaces/' + workspaceID + '/forms', {
                name,
                data: {
                    submit: 'Submit now',
                    blocks: {
                        blocks: [{
                            id: uuidv4(),
                            type: 'TEXT',
                            details: {
                                value: ''
                            },
                        }],
                        name
                    }
                }
            } as Form);
            this.forms.push(form);
            await this.select(form.ID);
            return form;
        },
        async loadSubmissions() {
            this.submissions = await api.get('workspaces/' + this.form.workspaceID + '/forms/' + this.form.ID + '/answers');
            return this.submissions;
        },
        getAliases(): { [key: string]: string } {
            const aliases = {};
            this.form.data.blocks.blocks.forEach(block => {
                if (block.details.hasOwnProperty('label'))
                    aliases[block.id] = block.details.label;
            });
            return aliases;
        }
    }
});