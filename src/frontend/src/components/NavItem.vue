<script setup lang="ts">
import { RouterLink } from "vue-router";

defineProps<{
  icon?: string;
  to?: string;
}>();
</script>

<template>
  <RouterLink :to="to || '/'"
    ><span
      class="flex flex-1 items-center my-4 justify-center sm:justify-start text-muted hover:scale-110 sm:hover:scale-105 transition-all"
      ><font-awesome-icon
        :icon="icon"
        class="w-8 h-8 sm:scale-75 transition-transform text-muted" /><span
        class="hidden max-w-[10rem] ml-2 font-semibold sm:flex"
        ><slot /></span></span
  ></RouterLink>
</template>

<style>
a.router-link-active * {
  color: #29a0b1;
}
</style>
