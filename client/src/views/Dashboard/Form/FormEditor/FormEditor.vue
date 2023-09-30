<script setup lang="ts">

import {onMounted, ref, watch} from "vue";
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {v4 as uuidv4} from 'uuid';
import {Lotion} from '@dashibase/lotion'
import '@dashibase/lotion/lib/style.css'
import {useDebouncedRefHistory} from "@vueuse/core";

const forms = useForms();
const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);

const page = ref({
  name: '',
  blocks: [{
    id: uuidv4(),
    type: 'TEXT',
    details: {
      value: ''
    },
  }],
});


watch(page, async (page) => {
  console.log(page);
  await forms.setName(page.name)
  await forms.setBlocks(page);
}, {deep: true});

const submitLabel = ref('');

async function saveSubmitLabel() {
  await forms.setSubmitLabel(submitLabel.value);
}

let undo = () => {
};

onMounted(async () => {
  await forms.load(workspaceID);
  forms.select(parseInt(route.params.formID));

  page.value = forms.form.data.blocks;
  submitLabel.value = forms.form.data.submit;

  const h = useDebouncedRefHistory(page, {deep: true, debounce: 1000});
  undo = h.undo;

  document.addEventListener('keyup', (event) => {
    if (event.keyCode == 90 && (event.ctrlKey || event.metaKey)) {
      console.log('undo');
      h.undo();
    }
  });
});

</script>

<template>
  <div class="grid gap-4">
    <div class="container ">
      <div class="remove-all">
        <Lotion :page="page"/>
      </div>

      <div style="width:650px;" class="mx-auto">
        <Button size="lg">
          Submit now
        </Button>

      </div>
    </div>
  </div>
</template>
