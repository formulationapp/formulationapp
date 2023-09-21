import {defineStore} from "pinia";
import {api} from "@/api";

interface Form {
    ID: number,
    name: string
    definition: string
    workspaceID: number
    secret: string
}

interface FormsState {
    forms: Form[]
    form: Form,
}


export const useForms = defineStore('forms', {
    state: (): FormsState => ({
        forms: [],
        form: {
            ID: 0,
            workspaceID: 0,
            name: ''
        },
    }),
    actions: {
        async load(workspaceID: number) {
            this.forms = await api.get('workspaces/' + workspaceID + '/forms')
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
        async setDefinition(definition: string) {
            this.form.definition = definition;
            await this.save(this.form);
        },
        async delete(workspaceID: number, formID: number) {
            await api.delete('workspaces/' + workspaceID + '/forms/' + formID);
            this.forms = this.forms.filter(form => form.ID != formID);
        },
        async create(workspaceID: number, name: string) {
            const form = await api.post('workspaces/' + workspaceID + '/forms', {
                name,
                definition: JSON.stringify([{
                    id: 'header',
                    type: 'header',
                    data: {
                        text: name,
                        level: 1
                    }
                }])
            } as Form);
            this.forms.push(form);
            await this.select(form.ID);
            return form;
        }
    }
});