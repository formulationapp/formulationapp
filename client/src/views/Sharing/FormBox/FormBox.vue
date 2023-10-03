<script setup lang="ts">
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {useSharing} from "@/stores/sharing";
import {Lotion} from '@dashibase/lotion';
import {ref} from "vue";

const sharing = useSharing();
const route = useRoute();

sharing.loadForm(route.params.secret);
const submitted = ref(false);

async function submit(event) {
  event.preventDefault();
  sharing.setSubmitted(true);
  if (Object.values(sharing.valid).some(x => !x))
    return;
  await sharing.sendAnswers();
  submitted.value = true;
}
</script>

<template>
  <div class="mt-8 md:mt-48 md:w-1/2 mx-auto w-11/12" v-if="!submitted&& sharing.form.data">

    <form @submit="submit">
      <div class="remove-all">
        <Lotion :page="sharing.form.data.blocks" :readonly="true"/>
      </div>

      <div style="width:650px;" class="mx-auto">

        <Button size="lg" type="submit">
          {{ sharing.form.data.submit }}
        </Button>

      </div>
    </form>
  </div>

  <div v-if="submitted">
    <div class="grid h-screen place-items-center text-center">
      <div class="space-y-3">
        <h1 class="text-4xl font-bold">Form submitted!</h1>
        <div>This form was created using <a href="https://formulationapp.com/" class="font-bold hover:underline ">Formulation</a>.
        </div>
        <Button @click="window.location='https://formulationapp.com/'">Create own Form for free!</Button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>