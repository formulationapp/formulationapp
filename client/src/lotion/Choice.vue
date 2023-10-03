<template>
  <div class="flex flex-col gap-2 mb-4">

    <div class="flex justify-between">
      <div class="flex ">
        <div>
          <Editor v-if="!readonly" placeholder="Enter label..." v-model="props.block.details.label"
                  class="font-semibold"/>
          <span v-if="readonly" class="font-semibold">{{ props.block.details.label }}</span>
        </div>

        <span class="ml-2 flex w-6 h-6 text-2xl rounded-full justify-center align-middle"
              :class="{'bg-gray-200 opacity-50':!props.block.details.required,'bg-indigo-200':props.block.details.required, 'invisible':!props.block.details.required}"
        >*</span>

        <span class="italic text-red-500 ml-2" v-if="!valid && sharing.submitted">This field is required.</span>
      </div>

      <HoverCard open-delay="100" close-delay="100" v-if="!readonly">
        <HoverCardTrigger>
          <Button variant="outline" class="ml-auto">
            <v-icon name="bi-gear-fill" class=" h-4 w-4 m-0"/>
          </Button>
        </HoverCardTrigger>
        <HoverCardContent class="p-3 space-y-2">
          <div class="flex justify-between">
            Required
            <Switch :checked="props.block.details.required"
                    @update:checked="props.block.details.required=!props.block.details.required"/>
          </div>
          <div class="flex justify-between">
            Multiple answer
            <Switch :checked="props.block.details.multiple"
                    @update:checked="props.block.details.multiple=!props.block.details.multiple"/>
          </div>
        </HoverCardContent>
      </HoverCard>

    </div>

    <div class="choice" v-for="(item,i) in props.block.details.choices">
      <div class="w-fit">
        <Button
            @click="select(props.block.details.choices[i])"
            :class="{'bg-indigo-100 brightness-90':selected[props.block.details.choices[i]]}" type="button"
            variant="outline" class="h-10 align-middle p-2">
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
import {computed, PropType, ref} from 'vue'
import {types} from "@dashibase/lotion";
import Editor from "@/lotion/Editor.vue";
import Button from "@/components/ui/button/Button.vue";
import Badge from "@/components/ui/badge/Badge.vue";
import {OhVueIcon} from "oh-vue-icons";
import {alphabet} from "@/lib/utils";
import Switch from "@/components/ui/switch/Switch.vue";
import HoverCard from "@/components/ui/hover-card/HoverCard.vue";
import HoverCardTrigger from "@/components/ui/hover-card/HoverCardTrigger.vue";
import HoverCardContent from "@/components/ui/hover-card/HoverCardContent.vue";
import {useSharing} from "@/stores/sharing";

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
  props.block.details.choices.push('');
}

const selected = ref({});
const sharing = useSharing();

function select(choice) {
  if(!props.readonly) return;
  selected.value[choice] = !selected.value[choice];
  if (!props.block.details.multiple)
    for (const key in selected.value)
      if (key != choice)
        selected.value[key] = false;

  sharing.setAnswer(props.block.id, JSON.stringify(selected.value), valid.value);
}

const valid = computed(() => {
  if (!props.block.details.required)
    return true;
  return Object.values(selected.value).some(v => v);
});

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