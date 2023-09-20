<script setup lang="ts">

import {Tabs, TabsContent} from "@/components/ui/tabs";
import {onMounted, ref} from "vue";
import EditorJS from "@editorjs/editorjs";
import ShortInput from "@/editor/ShortInput";
import ShortAnswer from "@/editor/ShortAnswer";
import Choice from "@/editor/Choice";
import Header from '@editorjs/header';
import List from '@editorjs/list';
import SimpleImage from '@editorjs/simple-image';
import Delimiter from '@editorjs/delimiter';
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";

const forms = useForms();
const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID[0]);

async function save(blocks: any) {
  console.log(blocks);
  const newName = blocks.filter(block => block.type == 'header')[0].data.text;
  await forms.setName(newName);
  await forms.setDefinition(JSON.stringify(blocks));
}

onMounted(async () => {
  await forms.load(workspaceID);
  forms.select(parseInt(route.params.formID[0]));
  console.log(JSON.parse(forms.form.definition));
  const editor = new EditorJS({
    holder: 'editorjs',
    async onChange(api, event) {
      console.log((await api.saver.save()).blocks);
      await save((await api.saver.save()).blocks);
    },
    placeholder: 'Type something to create your new form',
    data: {
      "time": 1550476186479,
      "blocks": JSON.parse(forms.form.definition),
      "version": "2.8.1"
    },
    tools: {
      header: Header,
      list: List,
      image: SimpleImage,
      delimiter: Delimiter,
      shortInput: ShortInput,
      shortAnswer: ShortAnswer,
      choice: Choice,
    },
  });
});
</script>

<template>
  <Tabs default-value="overview" class="space-y-4">
    <TabsContent value="overview" class="space-y-4">
      <div class="grid gap-4">
        <div class="container ">
          <div id="editorjs" class="no-tailwindcss titles"></div>
        </div>
      </div>
    </TabsContent>
  </Tabs>
</template>

<style scoped>

</style>