<script setup>
import { onMounted, computed, ref, onUnmounted } from 'vue';

const props = defineProps(['auction'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.auction.starting_price)
})

const timer = ref({
  days: 999,
  hours: 0,
  minutes: 0,
  seconds: 0
})

let intervalId = null

onMounted(async () => {
  var delay = 1000
  intervalId = setInterval(() => {
    let diff = (new Date(`${props.auction.ends_at} UTC`) - new Date()) / 1000
    timer.value.days = Math.floor(diff / 86400)
    timer.value.hours = Math.floor(diff / 3600 % 24).toString().padStart(2, '0')
    timer.value.minutes = Math.floor(diff / 60 % 60).toString().padStart(2, '0')
    timer.value.seconds = Math.floor(diff % 60).toString().padStart(2, '0')
    delay = 1000

    if (diff < 0) {
      clearInterval(intervalId)
    }
  }, delay)
})

onUnmounted(() => {
  clearInterval(intervalId)
})
</script>

<template>
  <div class="bg-gray-700 rounded shadow-md px-4 py-3">
    <div class="flex items-center justify-between border-b border-gray-600 pb-2">
      <div>
        <div :class="{ 'line-through': timer.seconds < 0 }" class="text-white font-bold tracking-wide text-lg">
          {{ auction.sld }}
          <span class="text-red-500 tracking-widest pl-0.5">.{{ auction.tld }}</span>
        </div>
        <div class="text-gray-400 text-sm">From ${{ startingPrice }}</div>
      </div>
      <div class="bg-red-800/70 px-2 py-2 rounded select-none">
        <div class="text-red-300 text-xs" v-if="timer.seconds >= 0">{{ timer.days }}d {{ timer.hours }}:{{ timer.minutes }}:{{ timer.seconds }}</div>
        <div class="text-red-300 text-xs uppercase tracking-widest" v-else>closed</div>
      </div>
    </div>

    <div>
      <p class="text-white font-light mt-2 mb-1 whitespace-pre-line" v-html="auction.description"></p>
    </div>
  </div>
</template>