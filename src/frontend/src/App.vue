<script setup lang="ts">
import { RouterView } from "vue-router";
import NavItem from "./components/NavItem.vue";
</script>

<template>
  <div class="h-24 flex">
    <div
      class="w-1/3 max-w-xs flex justify-center items-center w-1/2 min-w-[7rem] px-2"
    >
      <img src="@/assets/logo.svg" alt="logo" />
    </div>
    <div class="flex flex-1 items-center justify-end">
      <div class="mr-8">
        <font-awesome-icon icon="user" class="text-darkgray" />
        <span class="text-darkgray ml-2 font-semibold">Peter Pan</span>
      </div>
    </div>
  </div>
  <div class="flex flex-1">
    <div class="sm:w-1/3 max-w-xs flex">
      <div class="flex flex-1 flex-col items-center">
        <div class="inline-block pt-8 px-4">
          <NavItem to="/dashboard" icon="calendar-check">Reservations</NavItem>
          <NavItem to="/config/app" icon="cogs"
            >Application Configuration</NavItem
          >
          <NavItem to="/config/restaurant" icon="building"
            >Restaurant Configuration</NavItem
          >
        </div>
      </div>
    </div>
    <div class="bg-secondary flex flex-1"><RouterView /></div>
  </div>
</template>

<style>
@import "@/assets/index.css";
</style>
