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
  DialogHeader,
  DialogFooter,
  DialogContent,
  DialogTrigger,
  DialogTitle,
  DialogDescription
} from "@/components/ui/dialog";
import Button from "@/components/ui/button/Button.vue";
import Input from "@/components/ui/input/Input.vue";
import {ref} from "vue";

const route = useRoute();
const router = useRouter();
const workspaceID = parseInt(route.params.workspaceID[0]);
const forms = useForms();
forms.load(workspaceID);

const name = ref('');

async function create() {
  if (name.value.length == 0) return;
  const form = await forms.create(workspaceID, name.value);
  await router.push('/workspaces/' + workspaceID + '/forms/' + form.ID);
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

  <Alert @click="$router.push('/workspaces/'+workspaceID+'/forms/'+form.ID)" v-for="form in forms.forms"
         class="hover:brightness-95 cursor-pointer flex justify-between">
    <div>
      <AlertTitle class="font-semibold">{{ form.name }}
        <Badge class="ml-2">Szkic</Badge>
      </AlertTitle>
      <AlertDescription>
        Ostatnio edytowane we wtorek
      </AlertDescription>
    </div>

    <DropdownMenu>
      <DropdownMenuTrigger class="font-bold text-xl mr-4">
        ⋮
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuSeparator/>
        <DropdownMenuItem>Rename</DropdownMenuItem>
        <DropdownMenuItem>Duplicate</DropdownMenuItem>
        <DropdownMenuItem>Delete</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  </Alert>


</template>

<style scoped>

</style>