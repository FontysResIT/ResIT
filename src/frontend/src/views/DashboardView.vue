<script setup lang="ts">
import KPI from "../components/KPI.vue";
import TableComponent from "../components/TableComponent.vue";
import SearchInput from "../components/SearchInput.vue";
import DateInput from "../components/DateInput.vue";
import { onMounted, ref } from "vue";
import AxiosService from "@/AxiosService";
import ReservationModalVue from "../components/ReservationModal.vue";
let loading = ref(true);
const data = ref([]);
onMounted(async () => {
  data.value = await AxiosService.getReservationsByDate(
    new Date().toJSON().slice(0, 10).replace(/-/g, "-")
  );
  loading.value = false;
});

async function fetchDates(e: InputEvent) {
  console.log(e);
  loading.value = true;
  data.value = await AxiosService.getReservationsByDate(
    (e.target as HTMLInputElement).value
  );
  loading.value = false;
}
</script>

<template>
  <div class="flex flex-1 flex-col md:gap-6 gap-4 flex-wrap p-6">
    <ReservationModalVue />
    <div
      class="flex flex-1 sm:flex-grow-0 sm:h-auto flex-row md:gap-6 gap-4 flex-wrap"
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
    <div class="flex flex-1 flex-row">
      <div class="flex flex-1 flex-col bg-white rounded-3xl shadow-sm">
        <div class="flex flex-row p-6 pb-0 items-center flex-wrap box-content">
          <div
            class="flex flex-1 justify-start md:justify-center p-2 max-w-[200px]"
          >
            <span class="text-2xl font-bold text-primary">Reservations</span>
          </div>
          <div class="flex flex-[2] justify-center p-2"><SearchInput /></div>
          <div class="flex flex-1 justify-start md:justify-center p-2">
            <DateInput :change="fetchDates" />
          </div>
        </div>
        <div class="flex flex-row p-4">
          <TableComponent :data="data.slice(0, 10)" v-if="!loading" />
        </div>
      </div>
    </div>
  </div>
</template>
