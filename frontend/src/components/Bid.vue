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
  }, 1000)
})
</script>

<template>
  <a href="">
    <div class="bg-gray-700 rounded shadow-md duration-300 hover:scale-105 hover:shadow-xl px-4 py-3">
      <div class="flex items-center justify-between">
        <div>
          <div class="text-white font-bold tracking-wide">{{ domainName }}<span class="text-red-500">.{{ topLevel }}</span></div>
          <div class="text-gray-400 text-sm">From ${{ startingPrice }}</div>
        </div>
        <div class="bg-red-800 px-2 py-1 rounded-sm" v-show="days != 999">
          <div class="text-red-300 text-xs">{{ days }}d {{ hours }}:{{ minutes }}:{{ seconds }}</div>
        </div>
      </div>
    </div>
  </a>
</template>