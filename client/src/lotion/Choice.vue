<template>
  <div class="flex flex-col gap-2 mb-4">

    <div class="flex">
      <Editor v-if="!readonly" placeholder="Enter label..." v-model="props.block.details.label" class="font-semibold"/>
      <span v-if="readonly" class="font-semibold">{{ props.block.details.label }}</span>

      <span class="ml-1 flex px-1 text-md rounded-lg justify-center align-middle"
            :class="{'bg-gray-200 opacity-50':!props.block.details.multiple,'bg-orange-200':props.block.details.multiple, 'cursor-pointer':!readonly, 'invisible':readonly&&!props.block.details.multiple}"
            @click="toggleMultipleAnswer()"
      >Multiple answer</span>

      <span class="ml-1 flex w-6 h-6 text-2xl rounded-full justify-center align-middle"
            :class="{'bg-gray-200 opacity-50':!props.block.details.required,'bg-indigo-200':props.block.details.required, 'cursor-pointer':!readonly, 'invisible':readonly&&!props.block.details.required}"
            @click="toggleRequired()"
      >*</span>


    </div>

    <div class="choice" v-for="(item,i) in props.block.details.choices">
      <div class="w-fit">
        <Button variant="outline" class="h-10 align-middle p-2">
          <Badge class="mr-3 rounded-md w-6 h-6 justify-center">{{ alphabet[i] }}</Badge>
          <Editor v-if="!readonly" placeholder="Enter option..." class="font-semibold mr-2"
                  v-model="props.block.details.choices[i]"/>
          <span v-if="readonly" class="font-semibold mr-2">{{ props.block.details.choices[i] }}</span>
        </Button>
        <span v-if="!readonly && props.block.details.choices.length>1" class="tip ml-4"
              @click="props.block.details.choices.swap(i,i-1)"><OhVueIcon name="bi-arrow-up"/></span>
        <span v-if="!readonly && props.block.details.choices.length>1" class="tip ml-4"
              @click="props.block.details.choices.swap(i,i+1)"><OhVueIcon
            name="bi-arrow-down"/></span>
        <span v-if="!readonly && props.block.details.choices.length>1" class="tip ml-4 text-red-500"
              @click="props.block.details.choices.splice(i,1)">Delete</span>
      </div>
    </div>

    <span @click="addChoice" v-if="!readonly" class="opacity-50 cursor-pointer hover:opacity-100">Press ⌥ + G to add new option</span>
  </div>
</template>
<script setup lang="ts">
import {PropType} from 'vue'
import {types} from "@dashibase/lotion";
import Editor from "@/lotion/Editor.vue";
import Button from "@/components/ui/button/Button.vue";
import Badge from "@/components/ui/badge/Badge.vue";
import {OhVueIcon} from "oh-vue-icons";
import {alphabet} from "../lib/utils";

const props = defineProps({
  block: {
    type: Object as PropType<types.Block>,
    required: true,
    default: {
      details: {
        required: true,
        choices: ['']
      },
    }
  },
  readonly: {
    type: Boolean,
    default: false,
  },
})

if (!props.block.details.hasOwnProperty('choices'))
  props.block.details = {
    required: true,
    choices: ['']
  }

function addChoice() {
  console.log(props.block);
  props.block.details.choices.push('');
}

function onSet() {
}

function onUnset() {
}

function toggleRequired() {
  if (props.readonly) return;
  props.block.details.required = !props.block.details.required;
}

function toggleMultipleAnswer(){
  if (props.readonly) return;
  props.block.details.multiple = !props.block.details.multiple;
}

defineExpose({
  onSet,
  onUnset,
})
</script>

<style scoped>
.choice .tip {
  opacity: 0.7;
  display: none;
}

.choice:hover .tip {
  display: inline;
  cursor: pointer;
}

.choice:hover .tip:hover {
  opacity: 1;
  cursor: pointer;
}
</style>