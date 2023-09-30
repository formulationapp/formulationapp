<script setup lang="ts">

import {cn} from "@/lib/utils";
import {buttonVariants} from "@/components/ui/button";
import Label from "@/components/ui/label/Label.vue";
import Input from "@/components/ui/input/Input.vue";
import Button from "@/components/ui/button/Button.vue";
import {ref} from "vue";
import {useAuth} from "@/stores/auth";
import {useRouter} from "vue-router";
import {useWorkspaces} from "@/stores/workspaces";

const isLoading = ref(false)

const email = ref('');
const password = ref('');
const error = ref('');

const auth = useAuth();
const workspaces = useWorkspaces();
const router = useRouter();

async function onSubmit(event: Event) {
  event.preventDefault()
  error.value = '';
  isLoading.value = true;
  try {
    await auth.login(email.value, password.value);
    await router.push('/workspaces/' + workspaces.workspaces[0].ID);
  } catch (err) {
    error.value = (err as Error).message;
  }
  isLoading.value = false;
}

// if (localStorage.getItem('token') != null && localStorage.getItem('token')?.length > 10) {
//   await auth.useToken(localStorage.getItem('token'));
//   await router.push('/workspaces/' + workspaces.workspaces[0].ID);
// }
</script>

<template>
  <div class="md:hidden">
    <VPImage
        alt="Authentication"
        width="1280"
        height="1214" class="block" :image="{
        dark: '/examples/authentication-dark.png',
        light: '/examples/authentication-light.png',
      }"
    />
  </div>

  <div style="height: 100vh"
       class=" relative hidden h-full flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
    <router-link
        to="/register"
        :class="cn(
        buttonVariants({ variant: 'ghost' }),
        'absolute right-4 top-4 md:right-8 md:top-8',
      )"
    >
      Create an account
    </router-link>
    <div class="relative hidden h-full flex-col bg-muted p-10 text-white dark:border-r lg:flex">
      <div class="absolute inset-0 bg-zinc-900"/>
      <div class="relative z-20 flex items-center text-lg font-medium">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            class="mr-2 h-6 w-6"
        >
          <path d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"/>
        </svg>
        Formulation
      </div>
      <div class="relative z-20 mt-auto">
        <blockquote class="space-y-2">
          <p class="text-lg">
            &ldquo;Formulation is the easiest way to create forms and instantly gather feedback.&rdquo;
          </p>
          <footer class="text-sm">
            Przemek Sosna, Kalnica Ski Resorts
          </footer>
        </blockquote>
      </div>
    </div>
    <div class="lg:p-8 w-[600px] mx-auto">
      <div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
        <div class="flex flex-col space-y-2 text-center">
          <h1 class="text-2xl font-semibold tracking-tight">
            Sign in
          </h1>
          <p class="text-sm text-muted-foreground">
            Enter your email below to sign in to dashboard
          </p>
        </div>
        <div :class="cn('grid gap-6', $attrs.class ?? '')">
          <form @submit="onSubmit" autocomplete="on">
            <div class="grid gap-2">
              <div class="grid gap-1">
                <Label class="sr-only" for="email">
                  Email
                </Label>
                <Input
                    v-model="email"
                    id="email"
                    autocomplete="email"
                    placeholder="name@example.com"
                    type="email"
                    :disabled="isLoading"
                />
                <Input
                    v-model="password"
                    id="password"
                    autocomplete="password"
                    placeholder="password"
                    type="password"
                    :disabled="isLoading"
                    class="mt-2"
                />
              </div>
              <Button :disabled="isLoading" class="mt-2">
                Sign In
              </Button>
              <span class="text-red-500" v-if="error">{{ error }}</span>
            </div>
          </form>
        </div>
        <p class="px-8 text-center text-sm text-muted-foreground">
          <router-link to="/register"
                       class="underline underline-offset-4 hover:text-primary">
            Create an account
          </router-link>
        </p>
      </div>
    </div>
  </div>
</template>