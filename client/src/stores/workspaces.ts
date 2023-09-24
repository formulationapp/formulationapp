import {defineStore} from "pinia";
import {api} from "@/api";
import type Workspace from "@/models/workspace";

interface WorkspacesState {
    workspaces: Workspace[]
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