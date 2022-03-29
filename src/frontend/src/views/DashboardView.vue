<script setup lang="ts">
import KPI from "../components/KPI.vue";
</script>

<template>
  <div class="flex flex-1 flex-col gap-4 sm:gap-8 flex-wrap p-4">
    <div
      class="flex flex-1 sm:flex-grow-0 sm:h-auto flex-row gap-4 sm:gap-8 flex-wrap m-4"
    >
      <KPI
        icon="utensils"
        :count="19"
        :type="0"
        description="Reservations Today"
      />
      <KPI
        icon="times"
        :count="2"
        :type="1"
        description="Cancallations Today"
      />
      <KPI
        icon="clock"
        :count="5"
        :type="2"
        description="Available Time Slots"
      />
    </div>
    <div class="flex flex-1 flex-row m-4 mt-0">
      <div class="flex flex-1 bg-white rounded-3xl shadow-sm"></div>
    </div>
  </div>
</template>
