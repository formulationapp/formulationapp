<script setup lang="ts">
import {useSubmissions} from "@/stores/submissions";
import {useRoute} from "vue-router";

const submissions = useSubmissions();
const route = useRoute();

submissions.loadForm(route.params.secret);
</script>

<template>
{{submissions.form}}
</template>

<style scoped>

</style>