<script setup lang="ts">
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {useSharing} from "@/stores/sharing";

const sharing = useSharing();
const route = useRoute();

sharing.loadForm(route.params.secret);
</script>

<template>
  <div class="mt-8 md:mt-48 md:w-1/2 mx-auto w-11/12">

    <!--    {{ submissions.form.data }}-->

    <div v-for="block in sharing.form.data.blocks" class="mb-6">

      <div v-if="block.type=='header'">
        <h1 v-if="block.data.level==1" class="text-3xl font-semibold">{{ block.data.text }}</h1>
        <h1 v-if="block.data.level==2" class="text-2xl font-semibold">{{ block.data.text }}</h1>
        <h1 v-if="block.data.level==3" class="text-xl font-semibold">{{ block.data.text }}</h1>
        <h1 v-if="block.data.level==4" class="text-md font-semibold">{{ block.data.text }}</h1>
        <h1 v-if="block.data.level==5" class="text-xs font-semibold">{{ block.data.text }}</h1>
        <h1 v-if="block.data.level==6" class="text-2xs font-semibold">{{ block.data.text }}</h1>
      </div>

      <div v-if="block.type=='paragraph'">
        <span v-html="block.data.text"></span>
      </div>

      <div v-if="block.type=='list'" class="ml-12">
        <ol v-if="block.data.style=='unordered'" class="list-disc">
          <li v-for="item in block.data.items" v-html="item"></li>
        </ol>
        <ul v-if="block.data.style=='ordered'" class="list-decimal">
          <li v-for="item in block.data.items" v-html="item"></li>
        </ul>
      </div>

      <div v-if="block.type=='shortAnswer'">
        <span class="font-semibold">{{ block.data.label }}</span>
        <input type="text" :placeholder="block.data.text"
               class="t flex my-4 h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50">
      </div>
    </div>

    <Button size="lg">
      {{ sharing.form.data.submit }}
    </Button>

  </div>
</template>

<style scoped>

</style>