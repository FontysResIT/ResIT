<script setup lang="ts">
import type { Reservation } from "@/models/Reservation";

defineProps<{ data: Reservation[] }>();
</script>
<template>
  <div class="w-full">
    <table>
      <tr class="text-darkgray font-bold">
        <th class="hidden sm:table-cell">Id</th>
        <th>Name</th>
        <th class="hidden sm:table-cell">Email</th>
        <th>Guests</th>
        <th>Time Slot</th>
      </tr>
      <tr
        v-for="reservation in data"
        v-bind:key="reservation.id"
        class="hover:bg-secondary cursor-pointer"
      >
        <td class="hidden sm:table-cell">{{ reservation.id }}</td>
        <td>{{ `${reservation.first_name} ${reservation.last_name}` }}</td>
        <td class="hidden sm:table-cell">{{ reservation.email }}</td>
        <td>{{ reservation.guest_count }}</td>
        <td>12:00 - 13:30</td>
      </tr>
    </table>
  </div>
</template>
<style>
table {
  width: 100%;
}

td,
th {
  text-align: center;
  padding: 8px;
  max-width: 20%;
}
</style>
