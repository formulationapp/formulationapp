<script setup lang="ts">

import {cn} from "@/lib/utils";
import type {RouteLocationRaw} from "vue-router";

const props = defineProps<{
  label: String,
  link: RouteLocationRaw
}>();
</script>

<template>
    <router-link :to="props.link"
                 class="text-sm font-medium transition-colors hover:text-primary">
      {{ props.label }}
    </router-link>
</template>

<style scoped>

</style>