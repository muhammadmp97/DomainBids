<script setup>
import { onMounted, computed } from 'vue';

const props = defineProps(['auction'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.auction.starting_price)
})

const description = computed(() => {
  let description = String(props.auction.description)
  return description.replace("\n", '<br>')
})

const days = defineModel('days', { default: 999 })
const hours = defineModel('hours', { default: 0 })
const minutes = defineModel('minutes', { default: 0 })
const seconds = defineModel('seconds', { default: 0 })

onMounted(async () => {
  var delay = 5
  let timer = setInterval(() => {
    let diff = (new Date(`${props.auction.ends_at} UTC`) - new Date()) / 1000
    days.value = Math.floor(diff / 86400)
    hours.value = Math.floor(diff / 3600 % 24).toString().padStart(2, '0')
    minutes.value = Math.floor(diff / 60 % 60).toString().padStart(2, '0')
    seconds.value = Math.floor(diff % 60).toString().padStart(2, '0')
    delay = 1000

    if (diff < 0) {
      clearInterval(timer)
    }
  }, delay)
})
</script>

<template>
  <div class="bg-gray-700 rounded shadow-md px-4 py-3">
    <div class="flex items-center justify-between border-b border-gray-600 pb-2">
      <div>
        <div :class="(seconds < 0 ? 'line-through ' : '') + 'text-white font-bold tracking-wide text-lg'">
          {{ auction.sld }}
          <span class="text-red-500 tracking-widest pl-0.5">.{{ auction.tld }}</span>
        </div>
        <div class="text-gray-400 text-sm">From ${{ startingPrice }}</div>
      </div>
      <div class="bg-red-800/70 px-2 py-2 rounded select-none" v-show="days != 999">
        <div class="text-red-300 text-xs" v-if="seconds >= 0">{{ days }}d {{ hours }}:{{ minutes }}:{{ seconds }}</div>
        <div class="text-red-300 text-xs uppercase tracking-widest" v-if="seconds < 0">closed</div>
      </div>
    </div>

    <div>
      <p class="text-white font-light mt-2 mb-1" v-html="description"></p>
    </div>
  </div>
</template>