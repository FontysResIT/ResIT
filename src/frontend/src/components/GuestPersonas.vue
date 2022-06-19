<script setup lang="ts">
import type { Guest } from "@/models/Guest";
import { ref, type Ref } from "vue";

let guests: Ref<Guest[]> = ref([]);

function addGuest() {
  guests.value.push({ guest_name: "" } as Guest);
}

function removeGuest(guest: Guest) {
  const i = guests.value.indexOf(guest);
  guests.value.splice(i, 1);
}
</script>
<template>
  <button
    @click="addGuest"
    type="button"
    class="text-white bg-kpifront1 hover:bg-muted focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
  >
    <font-awesome-icon icon="plus" />
    Add Guest
  </button>
  <div class="overflow-y-auto max-h-48">
    <div class="flex flex-col space-y-2" v-for="(item, i) in guests" :key="i">
      <div class="flex flex-col flex-1">
        <label
          for="email"
          class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
          >Guest Name</label
        >
        <input
          v-model="item.guest_name"
          type="text"
          name="guestname"
          id="guestname"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
          placeholder="Sarah"
          required
        />
      </div>
      <div class="flex flex-1 flex-row items-end space-x-2">
        <div class="flex flex-col flex-1">
          <label
            for="remark"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
            >Dietary Requirement</label
          >
          <select
            name="requirement"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            id="requirement"
          >
            <option value="volvo">None</option>
            <option value="saab">Vegan</option>
            <option value="opel">Vegetarian</option>
            <option value="audi">Tree nut and peanut allergies</option>
            <option value="audi">Fish and shellfish allergies</option>
          </select>
        </div>
        <div class="flex flex-col">
          <button
            @click="() => removeGuest(item)"
            type="button"
            class="text-white bg-kpifront2 hover:bg-muted focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Remove
          </button>
        </div>
      </div>
      <hr class="py-4" />
    </div>
  </div>
</template>
