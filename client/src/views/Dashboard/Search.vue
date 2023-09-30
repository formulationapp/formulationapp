<script setup lang="ts">
import { Input } from '@/components/ui/input'
import Button from "@/components/ui/button/Button.vue";
import CommandBar from "@/components/CommandBar.vue";
</script>

<template>
  <div>
    <Input
        type="search"
        placeholder="Search..."
        class="md:w-[100px] lg:w-[230px]"
    />
  </div>
</template>