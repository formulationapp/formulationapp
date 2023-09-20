<script setup lang="ts">

import AlertDescription from "@/components/ui/alert/AlertDescription.vue";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import AlertTitle from "@/components/ui/alert/AlertTitle.vue";
import Badge from "@/components/ui/badge/Badge.vue";
import Alert from "@/components/ui/alert/Alert.vue";
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID[0]);
const forms = useForms();
forms.load(workspaceID);
</script>

<template>
  <h1 class="title text-2xl font-bold">Forms</h1>

  <Alert @click="$router.push('/workspaces/'+workspaceID+'/forms/'+form.ID)" v-for="form in forms.forms"
         class="hover:brightness-95 cursor-pointer flex justify-between">
    <div>
      <AlertTitle class="font-semibold">{{form.name}}
        <Badge class="ml-2">Szkic</Badge>
      </AlertTitle>
      <AlertDescription>
        Ostatnio edytowane we wtorek
      </AlertDescription>
    </div>

    <DropdownMenu>
      <DropdownMenuTrigger class="font-bold text-xl mr-4">
        ⋮
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuSeparator/>
        <DropdownMenuItem>Rename</DropdownMenuItem>
        <DropdownMenuItem>Duplicate</DropdownMenuItem>
        <DropdownMenuItem>Delete</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  </Alert>
</template>

<style scoped>

</style>