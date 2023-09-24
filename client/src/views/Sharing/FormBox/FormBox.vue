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

async function submit() {
  await sharing.sendAnswers();
  submitted.value = true;
}
</script>

<template>
  <div class="mt-8 md:mt-48 md:w-1/2 mx-auto w-11/12" v-if="!submitted">

    <div class="remove-all">
      <Lotion :page="sharing.form.data.blocks" :readonly="true"/>
    </div>

    <div style="width:650px;" class="mx-auto">

      <Button size="lg" @click="submit">
        {{ sharing.form.data.submit }}
      </Button>

    </div>

  </div>

  <div v-if="submitted">
    Submitted!
  </div>
</template>

<style scoped>

</style>