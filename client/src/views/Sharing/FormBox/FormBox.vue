<script setup lang="ts">
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {useSharing} from "@/stores/sharing";
import {Lotion} from '@dashibase/lotion';

const sharing = useSharing();
const route = useRoute();

sharing.loadForm(route.params.secret);
</script>

<template>
  <div class="mt-8 md:mt-48 md:w-1/2 mx-auto w-11/12">

    <div class="remove-all">
      <Lotion :page="sharing.form.data.blocks" :readonly="true"/>
    </div>

    <div style="width:650px;" class="mx-auto">
      <Button size="lg">
        {{ sharing.form.data.submit }}
      </Button>

    </div>

  </div>
</template>

<style scoped>

</style>