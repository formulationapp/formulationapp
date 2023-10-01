<script setup lang="ts">

import Search from "@/views/Dashboard/Search.vue";
import UserNav from "@/views/Dashboard/UserNav.vue";
import Button from "@/components/ui/button/Button.vue";
import MenuItem from "@/components/MenuItem.vue";
import {cn} from "@/lib/utils";
import {useRoute} from "vue-router";
import DarkModeSwitcher from "@/components/DarkModeSwitcher.vue";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
</script>

<template>
  <nav
      :class="cn('flex items-center space-x-4 lg:space-x-6', $attrs.class ?? '')">
    <MenuItem :link="'/workspaces/'+workspaceID">
      Forms
    </MenuItem>
  </nav>

  <div class="ml-auto flex items-center space-x-4">
<!--    <Search/>-->
    <UserNav/>
    <DarkModeSwitcher/>
  </div>
</template>
