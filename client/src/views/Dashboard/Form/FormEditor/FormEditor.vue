<script setup lang="ts">

import {onMounted, ref, watch} from "vue";
import {useForms} from "@/stores/forms";
import {useRoute} from "vue-router";
import Button from "@/components/ui/button/Button.vue";
import {v4 as uuidv4} from 'uuid';
import {Lotion} from '@dashibase/lotion'
import '@dashibase/lotion/lib/style.css'
import {useDebouncedRefHistory} from "@vueuse/core";
import AlertDialog from "@/components/ui/alert-dialog/AlertDialog.vue";
import AlertDialogTrigger from "@/components/ui/alert-dialog/AlertDialogTrigger.vue";
import AlertDialogContent from "@/components/ui/alert-dialog/AlertDialogContent.vue";
import AlertDialogHeader from "@/components/ui/alert-dialog/AlertDialogHeader.vue";
import AlertDialogTitle from "@/components/ui/alert-dialog/AlertDialogTitle.vue";
import AlertDialogDescription from "@/components/ui/alert-dialog/AlertDialogDescription.vue";
import AlertDialogFooter from "@/components/ui/alert-dialog/AlertDialogFooter.vue";
import AlertDialogCancel from "@/components/ui/alert-dialog/AlertDialogCancel.vue";
import AlertDialogAction from "@/components/ui/alert-dialog/AlertDialogAction.vue";
import Input from "@/components/ui/input/Input.vue";
import {createToaster} from "@meforma/vue-toaster";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger,} from '@/components/ui/dropdown-menu'

const forms = useForms();
const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);

const page = ref({
  name: '',
  blocks: [{
    id: uuidv4(),
    type: 'TEXT',
    details: {
      value: ''
    },
  }],
});


watch(page, async (page) => {
  await forms.setName(page.name)
  await forms.setBlocks(page);
}, {deep: true});

const submitLabel = ref('Submit');
const showSubmitLabelDialog = ref(false);

async function saveSubmitLabel() {
  if (submitLabel.value.length === 0) {
    return;
  }
  await forms.setSubmitLabel(submitLabel.value);
  await forms.save(forms.form);
  showSubmitLabelDialog.value = false;
  toaster.show(`Form saved!`, {
    type: 'success',
    duration: 2000,
  });
}

const toaster = createToaster({ /* options */});

async function saveForm() {
  await forms.save(forms.form);
  toaster.show(`Form saved!`, {
    type: 'success',
    duration: 2000,
  });
}


let undo = () => {
};

onMounted(async () => {
  await forms.load(workspaceID);
  forms.select(parseInt(route.params.formID));

  page.value = forms.form.data.blocks;
  submitLabel.value = forms.form.data.submit;

  const h = useDebouncedRefHistory(page, {deep: true, debounce: 1000});
  undo = h.undo;

  document.addEventListener('keyup', (event) => {
    if (event.keyCode == 90 && (event.ctrlKey || event.metaKey)) {
      console.log('undo');
      h.undo();
    }
  });
});

function insertQuestion(type: string) {
  if (page.value.blocks[page.value.blocks.length - 1].type === "TEXT") {
    page.value.blocks[page.value.blocks.length - 1] = {
      id: uuidv4(),
      type: type,
      details: {
        value: ''
      },
    };
  } else {
    page.value.blocks.push({
      id: uuidv4(),
      type: type,
      details: {
        value: ''
      },
    });
  }
}

</script>

<template>
  <div class="grid gap-4">
    <div class="container ">
      <div class="remove-all">
        <Lotion :page="page"/>
      </div>

      <div style="width:650px;" class="mx-auto mb-4">
        <DropdownMenu>
          <DropdownMenuTrigger>
            <Button variant="outline">Add question</Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuItem @click="insertQuestion('SHORT_ANSWER')">Short answer</DropdownMenuItem>
            <DropdownMenuItem @click="insertQuestion('CHOICE')">Choice</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <div style="width:650px;" class="mx-auto">

        <AlertDialog :open="showSubmitLabelDialog" @update:open="args => showSubmitLabelDialog=args">
          <AlertDialogTrigger as-child>
            <Button size="lg">
              {{ submitLabel }}
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Change submit label</AlertDialogTitle>
              <AlertDialogDescription class="mt-4">
                <Input v-model="submitLabel" autofocus @keyup.enter="saveSubmitLabel"/>
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Cancel</AlertDialogCancel>
              <AlertDialogAction @click="saveSubmitLabel" style="margin-top: 0">Save</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>

        <Button variant="ghost" size="lg" class="ml-4" @click="saveForm">Save form</Button>

      </div>
    </div>
  </div>
</template>
