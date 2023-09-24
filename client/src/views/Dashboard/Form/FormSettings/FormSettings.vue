<script setup lang="ts">

import Input from "@/components/ui/input/Input.vue";
import {ref} from "vue";
import {useForms} from "@/stores/forms";

const forms = useForms();

const link = ref('http://' + window.location.host + '/f/' + forms.form.secret);

async function openLink() {
  window.open(link.value, '_blank');
}
</script>

<template>
  Form settings<br/>

  <Input v-model="link" readonly @click="openLink" class="cursor-pointer"></Input>
</template>

<style scoped>

</style>