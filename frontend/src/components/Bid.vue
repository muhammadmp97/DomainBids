<script setup>
import { computed } from 'vue';

const props = defineProps(['index', 'name', 'price', 'createdAt'])

const price = computed(() => {
  return new Intl.NumberFormat().format(props.price)
})
</script>

<template>
  <div class="bg-gray-700 rounded shadow-md px-4 py-3 relative overflow-hidden">
    <div class="select-none absolute font-black text-8xl -top-2 left-1 text-gray-600/30 z-0">{{ index }}</div>
    <div class="flex items-center justify-between">
      <div class="relative z-1 | text-white">
        <div>{{ name }}</div>
        <div class="text-sm text-white/50">5m ago</div>
      </div>
      <div class="text-white font-bold">${{ price }}</div>
    </div>
  </div>
</template>