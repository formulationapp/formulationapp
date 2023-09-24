<script setup lang="ts">

import AlertDescription from "@/components/ui/alert/AlertDescription.vue";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import AlertTitle from "@/components/ui/alert/AlertTitle.vue";
import Badge from "@/components/ui/badge/Badge.vue";
import Alert from "@/components/ui/alert/Alert.vue";
import {useForms} from "@/stores/forms";
import {useRoute, useRouter} from "vue-router";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from "@/components/ui/dialog";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import Button from "@/components/ui/button/Button.vue";
import Input from "@/components/ui/input/Input.vue";
import {ref} from "vue";
import type Form from "@/models/form";

const route = useRoute();
const router = useRouter();
const workspaceID = parseInt(route.params.workspaceID);
const forms = useForms();
forms.load(workspaceID);

const name = ref('');

async function create() {
  if (name.value.length == 0) return;
  const form = await forms.create(workspaceID, name.value);
  await router.push('/workspaces/' + workspaceID + '/forms/' + form.ID);
}

const showDeleteDialog = ref(false);

async function askForDeleteForm(form: Form) {
  forms.select(form.ID);
  showDeleteDialog.value = true;
}

async function deleteForm(form: Form) {
  await forms.delete(workspaceID, form.ID);
}
</script>

<template>
  <div class="flex justify-between">
    <h1 class="title text-2xl font-bold">Forms</h1>

    <Dialog>
      <DialogTrigger as-child>
        <Button>
          New form
        </Button>
      </DialogTrigger>
      <DialogContent class="sm:max-w-[425px]" @escape-key-down.prevent>
        <DialogHeader>
          <DialogTitle>New form</DialogTitle>
          <DialogDescription>
            Enter title of your new fantastic form!
          </DialogDescription>
        </DialogHeader>
        <div class="grid gap-4 py-4">
          <div class="grid grid-cols-4 items-center gap-4">
            <Input v-model="name" placeholder="Form name..." id="name" class="col-span-4"/>
          </div>
        </div>
        <DialogFooter>
          <Button type="submit" @click="create">
            Create now
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>

  <Alert @click.self="$router.push('/workspaces/'+workspaceID+'/forms/'+form.ID)" v-for="form in forms.forms"
         class="hover:brightness-95 cursor-pointer flex justify-between">
    <div @click="$router.push('/workspaces/'+workspaceID+'/forms/'+form.ID)">
      <AlertTitle class="font-semibold">{{ form.name }}
        <Badge class="ml-2">Szkic</Badge>
      </AlertTitle>
      <AlertDescription>
        Ostatnio edytowane we wtorek
      </AlertDescription>
    </div>

    <DropdownMenu>
      <DropdownMenuTrigger class="font-bold text-xl mr-4">
        <Button variant="outline">⋮</Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuSeparator/>
        <DropdownMenuItem @click="askForDeleteForm(form)" class="text-red-500 hover:text-red-700">Delete
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  </Alert>

  <AlertDialog :open="showDeleteDialog" @update:open="args => showDeleteDialog=args">
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Are you sure?</AlertDialogTitle>
        <AlertDialogDescription>
          You are about to delete {{ forms.form.name }} form. This action cannot be undone.
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Cancel</AlertDialogCancel>
        <AlertDialogAction @click="deleteForm(forms.form)">Delete now</AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  </AlertDialog>
</template>

<style scoped>

</style>