<script setup lang="ts">

import {Tabs, TabsContent} from "@/components/ui/tabs";
import {onMounted} from "vue";
import EditorJS from "@editorjs/editorjs";
import ShortInput from "@/editor/ShortInput";
import Choice from "@/editor/Choice";
import Header from '@editorjs/header';
import List from '@editorjs/list';
import SimpleImage from '@editorjs/simple-image';
import Delimiter from '@editorjs/delimiter';


onMounted(() => {
  const editor = new EditorJS({
    holder: 'editorjs',
    async onChange(api, event) {
      const s = await api.saver.save();
      console.log(s.blocks);
    },
    placeholder: 'Type something to create your new form',
    data: {
      "time": 1550476186479,
      "blocks": [
        {
          "type": "header",
          "data": {
            "text": "Wycieczka szkolna",
            "level": 1
          }
        },
      ],
      "version": "2.8.1"
    },
    tools: {
      header: Header,
      list: List,
      image: SimpleImage,
      delimiter: Delimiter,
      shortInput: ShortInput,
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