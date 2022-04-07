<script setup lang="ts">
interface Props {
  label: string;
}

withDefaults(defineProps<Props>(), {
  label: "Today",
});

const today = new Date().toISOString().slice(0, 10);
</script>
<template>
  <div>
    <input type="date" id="datePicker" :value="today" />
  </div>
</template>
