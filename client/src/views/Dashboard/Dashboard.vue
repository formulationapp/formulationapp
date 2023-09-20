<script setup lang="ts">

import {useRoute} from "vue-router";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
</script>

<template>
  <div class="hidden flex-col md:flex">
    <div class="border-b container">
      <div class="flex h-16 items-center px-4">
        <div @click="$router.push('/workspaces/'+workspaceID)"
             class="relative z-20 flex items-center text-lg font-medium mr-8 cursor-pointer">
          <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              class="mr-2 h-6 w-6"
          >
            <path d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"/>
          </svg>
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