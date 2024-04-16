<script setup>
import { computed } from 'vue';

const props = defineProps(['bid', 'index'])

const price = computed(() => {
  return new Intl.NumberFormat().format(props.bid.price)
})

let time = computed(() => {
  let time = new Date(props.bid.created_at)
  let diff = Math.floor((new Date() - time) / 1000)

  if (diff < 60) {
    return `${diff}s`
  }

  if (diff < 3600) {
    diff = Math.floor(diff / 60)
    return `${diff}m`
  }

  if (diff < 3600 * 24) {
    diff = Math.floor(diff / 3600)
    return `${diff}h`
  }

  diff = Math.floor(diff / (3600 * 24))
  return `${diff}d`
})
</script>

<template>
  <div class="bg-gray-700 rounded shadow-md px-4 py-3 relative overflow-hidden">
    <div class="select-none absolute font-black text-8xl -top-2 left-1 text-gray-600/30 z-0">{{ index }}</div>
    <div class="flex items-center justify-between">
      <div class="relative z-1 | text-white">
        <div>{{ bid.user.nickname }}</div>
        <div class="text-sm text-white/50" :title="bid.created_at">{{ time }} ago</div>
      </div>
      <div class="text-white font-bold">${{ price }}</div>
    </div>
  </div>
</template>