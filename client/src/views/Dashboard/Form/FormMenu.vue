<script setup lang="ts">

import Search from "@/views/Dashboard/Search.vue";
import MenuItem from "@/components/MenuItem.vue";
import Button from "@/components/ui/button/Button.vue";
import UserNav from "@/views/Dashboard/UserNav.vue";
import {cn} from "@/lib/utils";
import {useRoute} from "vue-router";
import {useForms} from "@/stores/forms";
import Badge from "@/components/ui/badge/Badge.vue";
import DarkModeSwitcher from "@/components/DarkModeSwitcher.vue";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
const formID = parseInt(route.params.formID);

const forms = useForms();

</script>

<template>

  <nav
      :class="cn('flex items-center space-x-4 lg:space-x-6', $attrs.class ?? '')">
    <MenuItem :link="'/workspaces/'+workspaceID">
      Forms
    </MenuItem>
  </nav>

  <div class="flex mx-auto border rounded-full px-4 py-2 shadow">
    <nav class="flex space-x-4 lg:space-x-6">
      <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID" :class="{'opacity-50':route.name=='design'}">
        Design
      </MenuItem>
      <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/submissions'" :class="{'opacity-50':route.name=='submissions'}">
        Submissions
      </MenuItem>
      <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/settings'" :class="{'opacity-50':route.name=='form_settings'}">
        Settings
      </MenuItem>
    </nav>
  </div>

  <div class="ml-auto flex items-center space-x-4">
    <Button @click="$router.push('/workspaces/'+workspaceID+'/forms/'+formID+'/settings')">Publish now!</Button>
<!--    <Search/>-->
    <UserNav/>
    <DarkModeSwitcher/>
  </div>
</template>