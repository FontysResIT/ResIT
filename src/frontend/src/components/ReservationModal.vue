<script setup lang="ts">
import AxiosService from "@/AxiosService";
import type { Guest } from "@/models/Guest";
import type { ReservationCreate } from "@/models/ReservationCreate";
import { ref, type Ref } from "vue";
import GuestPersonas from "./GuestPersonas.vue";

let reservation: Ref<ReservationCreate> = ref({
  guest_persona: [] as Guest[],
  restaurant_id: "4543fe3b42a029933ca90432",
  dts_id: "624ee1fe3b42a029933c7a90",
  is_cancelled: false,
  is_rescheduled: false,
  guest_count: 1,
} as ReservationCreate);

function submit() {
  console.log(reservation.value);
  AxiosService.createReservation(reservation.value);
}
</script>
<template>
  <div
    id="authentication-modal"
    tabindex="-1"
    aria-hidden="true"
    class="hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 w-full md:inset-0 h-modal md:h-full justify-center items-center"
  >
    <div class="relative p-4 w-full max-w-4xl h-full md:h-auto">
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <button
          type="button"
          class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
          data-modal-toggle="authentication-modal"
        >
          <svg
            class="w-5 h-5"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            ></path>
          </svg>
        </button>
        <div class="py-6 px-6 lg:px-8">
          <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
            Create a New Reservation
          </h3>
          <form class="space-y-6" @submit.prevent>
            <div class="flex flex-row space-x-2">
              <div class="flex flex-col flex-1">
                <label
                  for="firstname"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Your firstname</label
                >
                <input
                  v-model="reservation.first_name"
                  type="text"
                  name="firstname"
                  id="firstname"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  placeholder="John"
                  required
                />
              </div>
              <div class="flex flex-col flex-1">
                <label
                  for="password"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Your lastname</label
                >
                <input
                  v-model="reservation.last_name"
                  type="text"
                  name="lastname"
                  id="lastname"
                  placeholder="Doe"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  required
                />
              </div>
            </div>
            <div class="flex flex-row space-x-2">
              <div class="flex flex-col flex-1">
                <label
                  for="email"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Your Email</label
                >
                <input
                  v-model="reservation.email"
                  type="email"
                  name="email"
                  id="email"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  placeholder="info@example.com"
                  required
                />
              </div>
              <div class="flex flex-col flex-1">
                <label
                  for="phone"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Your Phone</label
                >
                <input
                  v-model="reservation.phone_number"
                  type="text"
                  name="phone"
                  id="phone"
                  placeholder="+31 6 37837388"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  required
                />
              </div>
            </div>
            <div class="flex flex-row space-x-2">
              <div class="flex flex-col flex-1">
                <label
                  for="remark"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Your Remark</label
                >
                <textarea
                  v-model="reservation.remark"
                  rows="1"
                  name="remark"
                  id="remark"
                  placeholder="It's my birthday..."
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                />
              </div>
            </div>
            <div class="flex flex-row space-x-2">
              <div class="flex flex-col flex-1">
                <label
                  for="remark"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Date</label
                >
                <input
                  type="date"
                  id="start"
                  name="trip-start"
                  value="2022-07-22"
                  min="2022-01-01"
                  max="2022-12-31"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                />
              </div>
              <div class="flex flex-col flex-1">
                <label
                  for="remark"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Timeslot</label
                >
                <input
                  type="time"
                  id="start"
                  name="trip-start"
                  value="18:00"
                  min="09:00"
                  max="22:00"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                />
              </div>
            </div>
            <GuestPersonas :guests="reservation.guest_persona" />
            <button
              type="submit"
              @click="submit"
              class="w-full text-white bg-primary hover:bg-muted focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
            >
              Create
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
