<script setup lang="ts">

import Search from "@/views/Dashboard/Search.vue";
import MenuItem from "@/components/MenuItem.vue";
import Button from "@/components/ui/button/Button.vue";
import UserNav from "@/views/Dashboard/UserNav.vue";
import {cn} from "@/lib/utils";
import {useRoute} from "vue-router";
import {useForms} from "@/stores/forms";
import Badge from "@/components/ui/badge/Badge.vue";

const route = useRoute();
const workspaceID = parseInt(route.params.workspaceID);
const formID = parseInt(route.params.formID);

const forms = useForms();

</script>

<template>

  <nav
      :class="cn('flex items-center space-x-4 lg:space-x-6', $attrs.class ?? '')">
    <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID">
      Form
    </MenuItem>
    <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/submissions'">
      Submissions
      <Badge v-if="forms.submissions.length>0"> {{ forms.submissions.length }}</Badge>
    </MenuItem>
    <MenuItem :link="'/workspaces/'+workspaceID+'/forms/'+formID+'/settings'">
      Settings
    </MenuItem>
  </nav>

  <div class="ml-auto flex items-center space-x-4">
    <Button @click="$router.push('/workspaces/'+workspaceID+'/forms/'+formID+'/settings')">Publish now!</Button>
<!--    <Search/>-->
    <UserNav/>
  </div>
</template>