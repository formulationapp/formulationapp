<script setup lang="ts">
import {useWorkspaces} from "@/stores/workspaces";
import {useRouter} from "vue-router";
import {onMounted} from "vue";

const workspaces = useWorkspaces();
const router = useRouter();
onMounted(async () => {
  await workspaces.load();
  router.push('/workspaces/' + workspaces.workspaces[0].ID)
})
</script>

<template>
  <main>
  </main>
</template>
