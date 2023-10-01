<script setup lang="ts">

import Input from "@/components/ui/input/Input.vue";
import {ref} from "vue";
import {useForms} from "@/stores/forms";
import {copyToClipboard} from "@/utils";

const forms = useForms();

const link = ref('http://' + window.location.host + '/f/' + forms.form.secret);


</script>

<template>
  <div class="container mt-12">
   <div class="w-1/2 mx-auto">
      <h3 class="title text-2xl font-bold ">Form link</h3>
      <p class="text-gray-500">Share this link with your clients</p>
      <br/>

     <Input v-model="link" readonly @click="copyToClipboard(link)" class="cursor-pointer"></Input>
   </div>
  </div>
</template>

<style scoped>

</style>