<script setup>
import { onMounted } from 'vue';
import { computed } from 'vue';

const props = defineProps(['domainName', 'topLevel', 'startingPrice', 'endsAt'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.startingPrice)
})

const days = defineModel('days', { default: 999 })
const hours = defineModel('hours', { default: 0 })
const minutes = defineModel('minutes', { default: 0 })
const seconds = defineModel('seconds', { default: 0 })

onMounted(() => {
  let timer = setInterval(() => {
    let diff = (new Date(props.endsAt) - new Date()) / 1000
    days.value = Math.floor(diff / 86400)
    hours.value = Math.floor(diff / 3600 % 24).toString().padStart(2, '0')
    minutes.value = Math.floor(diff / 60 % 60).toString().padStart(2, '0')
    seconds.value = Math.floor(diff % 60).toString().padStart(2, '0')

    if (diff < 0) {
      clearInterval(timer)
    }
  }, 1000)
})
</script>

<template>
  <a href="">
    <div class="bg-gray-700 rounded shadow-md duration-300 hover:scale-105 hover:shadow-xl px-4 py-3">
      <div class="flex items-center justify-between">
        <div>
          <div :class="(seconds < 0 ? 'line-through ' : '') + 'text-white font-bold tracking-wide'">{{ domainName }}<span class="text-red-500 tracking-widest pl-0.5">.{{ topLevel }}</span></div>
          <div class="text-gray-400 text-sm mt-0.5">From ${{ startingPrice }}</div>
        </div>
        <div class="bg-red-800/70 px-2 py-2 rounded" v-show="days != 999">
          <div class="text-red-300 text-xs" v-if="seconds >= 0">{{ days }}d {{ hours }}:{{ minutes }}:{{ seconds }}</div>
          <div class="text-red-300 text-xs uppercase tracking-widest" v-if="seconds < 0">closed</div>
        </div>
      </div>
    </div>
  </a>
</template>