<script setup lang="ts">
import KPI from "../components/KPI.vue";
</script>

<template>
  <div class="flex flex-1 flex-col gap-4 sm:gap-8 flex-wrap">
    <div
      class="flex h-1/5 flex-1 sm:flex-grow-0 sm:h-auto flex-row px-16 pt-16 gap-4 sm:gap-8 flex-wrap"
    >
      <KPI icon="utensils" count="19" description="Reservations Today" />
      <KPI icon="times" count="2" description="Cancallations Today" />
      <KPI icon="clock" count="5" description="Available Time Slots" />
    </div>
    <div class="flex flex-1 flex-row px-16 pb-16">
      <div class="flex flex-1 bg-white rounded-3xl shadow-sm"></div>
    </div>
  </div>
</template>
