<script setup lang="ts">
interface Props {
  label: string;
  change: (a: any) => void;
}

withDefaults(defineProps<Props>(), {
  label: "Today",
});

const today = new Date().toISOString().slice(0, 10);
</script>
<template>
  <div class="flex justify-around flex-1">
    <div>
      <input
        type="date"
        class="border-none"
        id="datePicker"
        :value="today"
        @change="change"
      />
    </div>
    <div>
      <button
        class="rounded-md px-4 py-2 bg-primary text-white cursor-pointer"
        data-modal-toggle="authentication-modal"
      >
        <font-awesome-icon icon="plus" size="xl" />
        <span class="font-bold ml-1">CREATE</span>
      </button>
    </div>
  </div>
</template>
