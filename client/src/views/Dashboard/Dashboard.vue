<script setup lang="ts">

import {useRoute} from "vue-router";
import CommandBar from "@/components/CommandBar.vue";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
</script>

<template>

<!--  <CommandBar/>-->

  <div class="flex-col md:flex">
    <div class="border-b container">
      <div class="flex h-16 items-center px-4">
        <div @click="$router.push('/workspaces/'+workspaceID)"
             class="relative z-20 flex items-center text-lg font-medium mr-8 cursor-pointer">
          <img src="../../assets/icon.png" class="w-8 h-8 mr-3" alt="Formulation">
          Formulation
        </div>
        <router-view name="menu" class="mx-6 "/>
      </div>
    </div>
    <div class="container flex-1 space-y-4 p-8 pt-6 w-9/12">
      <router-view/>
    </div>
  </div>
</template>

<style scoped>

</style>