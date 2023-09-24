<script setup lang="ts">

import type {RouteLocationRaw} from "vue-router";

const props = defineProps<{
  link: RouteLocationRaw
}>();
</script>

<template>
  <router-link :to="props.link"
               class="text-sm font-medium transition-colors hover:text-primary">
    <slot/>
  </router-link>
</template>

<style scoped>

</style>