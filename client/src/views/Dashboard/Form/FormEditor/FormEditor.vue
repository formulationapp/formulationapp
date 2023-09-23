<script setup lang="ts">

import {onMounted, ref, watch} from "vue";
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {v4 as uuidv4} from 'uuid';
import {Lotion, registerBlock} from '@dashibase/lotion'
import '@dashibase/lotion/lib/style.css'
import ShortAnswer from "@/lotion/ShortAnswer.vue";

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

onMounted(async () => {
  await forms.load(workspaceID);
  forms.select(parseInt(route.params.formID));

  page.value = forms.form.data.blocks;
  submitLabel.value = forms.form.data.submit;
});

</script>

<template>

  Secret: {{ forms.form.secret }}

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
