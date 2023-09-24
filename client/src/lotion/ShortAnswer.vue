<template>
  <div class="flex">

    <Editor v-if="!readonly" placeholder="Enter label..." v-model="props.block.details.label" class="font-semibold"/>
    <span v-if="readonly" class="font-semibold">{{ props.block.details.label }}</span>

    <span class="ml-1 flex w-6 h-6 text-2xl rounded-full justify-center align-middle"
          :class="{'bg-gray-200 opacity-50':!props.block.details.required,'bg-indigo-200':props.block.details.required, 'cursor-pointer':!readonly, 'invisible':readonly&&!props.block.details.required}"
          @click="toggleRequired()"
    >*</span>
  </div>

  <input :required="props.block.details.required" type="text" v-model="value" @change="setAnswer"
         class="flex my-4 h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50">
</template>
<script setup lang="ts">
import {PropType, ref} from 'vue'
import {types} from "@dashibase/lotion";
import Editor from "@/lotion/Editor.vue";
import {useSharing} from "@/stores/sharing";

const sharing = useSharing();

const value = ref();

const props = defineProps({
  block: {
    type: Object as PropType<types.Block>,
    required: true,
    default: {
      details: {required: true}
    }
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

function toggleRequired() {
  if (props.readonly) return;
  props.block.details.required = !props.block.details.required;
}

function setAnswer() {
  sharing.setAnswer(props.block.id, value.value);
}

defineExpose({
  onSet,
  onUnset,
})
</script>