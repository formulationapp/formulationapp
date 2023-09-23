<template>
  <div class=" flex items-center justify-between">
    <div class="w-6/12">
      <input type="text"
             class="flex my-4 h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50">
    </div>
    <div class="opacity-50 cursor-pointer hover:opacity-100">
      Press ⌥+Q to add title
    </div>
  </div>
</template>
<script setup lang="ts">
import {PropType} from 'vue'
import {types} from "@dashibase/lotion";

const props = defineProps({
  block: {
    type: Object as PropType<types.Block>,
    required: true,
  },
  readonly: {
    type: Boolean,
    default: false,
  },
})

function onSet() {
}

function onUnset() {
}

defineExpose({
  onSet,
  onUnset,
})
</script>