<script setup lang="ts">
import {RouterView, useRoute} from 'vue-router'
import {onMounted} from "vue";
import {useAuth} from "@/stores/auth";
import {useWorkspaces} from "@/stores/workspaces";
import {useForms} from "@/stores/forms";


const route = useRoute();
const auth = useAuth();
const forms = useForms();
const workspaces = useWorkspaces();

onMounted(async () => {
  await auth.try();
  await workspaces.load();
  if (route.params.hasOwnProperty('workspaceID'))
    await forms.load(parseInt(route.params.workspaceID[0]));
  if (route.params.hasOwnProperty('formID'))
    await forms.select(parseInt(route.params.formID[0]));
})
</script>

<template>
  <RouterView/>
</template>
