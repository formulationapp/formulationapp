import {defineStore} from "pinia";
import {api} from "@/api";

export const useSubmissions = defineStore('submissions', {
    state: () => ({}),
    actions: {
        getForm(secret:string){
            api.get('/')
        }
    }
})