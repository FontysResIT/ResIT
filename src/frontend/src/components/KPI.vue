<script setup lang="ts">
defineProps<{
  icon?: string;
  description?: string;
  count?: number;
  type?: number;
}>();
</script>

<template>
  <div
    class="flex flex-1 bg-white rounded-3xl shadow-sm justify-center items-center"
  >
    <div class="flex flex-row p-4">
      <div
        class="h-16 min-w-[4rem] rounded-full flex justify-center items-center"
        :class="[
          type === 0 && 'bg-kpibg1',
          type === 1 && 'bg-kpibg2',
          type === 2 && 'bg-kpibg3',
        ]"
      >
        <font-awesome-icon
          :icon="icon"
          class="w-8 h-8"
          :class="[
            type === 0 && 'text-kpifront1',
            type === 1 && 'text-kpifront2',
            type === 2 && 'text-kpifront3',
          ]"
        />
      </div>
      <div class="flex flex-col ml-2">
        <span class="font-bold text-3xl text-darkgray">{{ count }}</span>
        <span class="min-w-[8rem] text-muted">{{ description }}</span>
      </div>
    </div>
  </div>
</template>
