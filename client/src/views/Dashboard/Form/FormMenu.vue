<script setup lang="ts">

import Search from "@/views/Dashboard/Search.vue";
import MenuItem from "@/components/MenuItem.vue";
import Button from "@/components/ui/button/Button.vue";
import UserNav from "@/views/Dashboard/UserNav.vue";
import {cn} from "@/lib/utils";
import {useRoute} from "vue-router";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
const formID = parseInt(route.params.formID);
</script>

<template>

  <nav
      :class="cn('flex items-center space-x-4 lg:space-x-6', $attrs.class ?? '')">
    <MenuItem label="Form" :link="'/workspaces/'+workspaceID+'/forms/'+formID"/>
    <MenuItem label="Submissions" :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/submissions'"/>
    <MenuItem label="Settings" :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/settings'"/>
  </nav>

  <div class="ml-auto flex items-center space-x-4">
    <Button>Publish!</Button>
    <Search/>
    <UserNav/>
  </div>
</template>