<script setup lang="ts">
/* eslint-disable vue/no-mutating-props */
import type { Guest } from "@/models/Guest";

const props = defineProps(["guests"]);
console.log(props.guests);
function addGuest() {
  props.guests.push({
    guest_name: "",
    food_preferences: [] as string[],
    dietary_requirements: [] as string[],
  } as Guest);
}

function removeGuest(guest: Guest) {
  const i = props.guests.indexOf(guest);
  props.guests.splice(i, 1);
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
            v-model="item.dietary_requirements"
            multiple
            name="requirement"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            id="requirement"
          >
            <option value="NONE">None</option>
            <option value="VEGAN">Vegan</option>
            <option value="VEGETARIAN">Vegetarian</option>
            <option value="PEANUT">Tree nut and peanut allergies</option>
            <option value="FISH">Fish and shellfish allergies</option>
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
