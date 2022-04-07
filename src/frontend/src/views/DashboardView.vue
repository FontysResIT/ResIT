<script setup lang="ts">
import KPI from "../components/KPI.vue";
import TableComponent from "../components/TableComponent.vue";
import SearchInput from "../components/SearchInput.vue";
import DateInput from "../components/DateInput.vue";
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
      <div class="flex flex-1 flex-col bg-white rounded-3xl shadow-sm">
        <div class="flex flex-row p-4 h-1/6 items-center flex-wrap box-content">
          <div class="flex flex-1 justify-center px-2 max-w-[200px]">
            <span class="text-2xl font-bold text-primary">Reservations</span>
          </div>
          <div class="flex flex-[2] justify-center px-2"><SearchInput /></div>
          <div class="flex flex-1 justify-center px-2"><DateInput /></div>
        </div>
        <div class="flex flex-row p-4">
          <TableComponent />
        </div>
      </div>
    </div>
  </div>
</template>
