import {defineStore} from "pinia";
import {api} from "@/api";

interface FormsState {
    forms: {
        ID: number,
        Name: string
    }[]
}


export const useForms = defineStore('forms', {
    state: (): FormsState => ({
        forms: []
    }),
    actions: {
        async load(workspaceID: number) {
            this.forms = await api.get('workspaces/' + workspaceID + '/forms')
        }
    }
});