import {defineStore} from "pinia";
import {api} from "@/api";

interface WorkspacesState {
    workspaces: [{
        ID: number
    }]
}

export const useWorkspaces = defineStore('workspaces', {
    state: (): WorkspacesState => ({
        workspaces: []
    }),
    actions: {
        async load() {
            this.workspaces = await api.get('workspaces');
        }
    }
})