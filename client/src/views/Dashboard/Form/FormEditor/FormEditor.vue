<script setup lang="ts">

import {onMounted, ref} from "vue";
import EditorJS from "@editorjs/editorjs";
import ShortAnswer from "@/editor/ShortAnswer";
import Header from '@editorjs/header';
import List from '@editorjs/list';
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";

const forms = useForms();
const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);

async function save(blocks: any) {
  const name = blocks.filter(block => block.type == 'header')[0].data.text;
  await forms.setName(name);
  await forms.setBlocks(blocks);
}

const submitLabel = ref('');

async function saveSubmitLabel() {
  await forms.setSubmitLabel(submitLabel.value);
}

onMounted(async () => {
  await forms.load(workspaceID);
  forms.select(parseInt(route.params.formID));

  submitLabel.value = forms.form.data.submit;

  const editor = new EditorJS({
    holder: 'editorjs',
    async onChange(api, event) {
      console.log((await api.saver.save()).blocks);
      await save((await api.saver.save()).blocks);
    },
    placeholder: 'Type something to create your new form',
    data: {
      "time": 1550476186479,
      "blocks": forms.form.data.blocks,
      "version": "2.8.1"
    },
    tools: {
      header: Header,
      list: List,
      shortAnswer: ShortAnswer,
    },
  });
});
import contenteditable from 'vue-contenteditable';

</script>

<template>

  Secret: {{ forms.form.secret }}

  <div class="grid gap-4">
    <div class="container ">
      <div id="editorjs" class="no-tailwindcss titles"></div>

      <div style="width:650px;" class="mx-auto">
        <Button size="lg">
          <contenteditable tag="div"
                           :contenteditable="true"
                           @update:modelValue="saveSubmitLabel"
                           v-model="submitLabel"/>
        </Button>

      </div>
    </div>
  </div>
</template>
