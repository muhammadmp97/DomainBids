<script setup>
import { onMounted, computed } from 'vue';

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
  <div class="bg-gray-700 rounded shadow-md px-4 py-3">
    <div class="flex items-center justify-between border-b border-gray-600 pb-2">
      <div>
        <div :class="(seconds < 0 ? 'line-through ' : '') + 'text-white font-bold tracking-wide text-lg'">
          {{ domainName }}
          <span class="text-red-500 tracking-widest pl-0.5">.{{ topLevel }}</span>
        </div>
        <div class="text-gray-400 text-sm">From ${{ startingPrice }}</div>
      </div>
      <div class="bg-red-800/70 px-2 py-2 rounded select-none" v-show="days != 999">
        <div class="text-red-300 text-xs" v-if="seconds >= 0">{{ days }}d {{ hours }}:{{ minutes }}:{{ seconds }}</div>
        <div class="text-red-300 text-xs uppercase tracking-widest" v-if="seconds < 0">closed</div>
      </div>
    </div>

    <div>
      <p class="text-white font-light mt-2 mb-1">In publishing and graphic design, Lorem ipsum is a placeholder text
        commonly used to demonstrate the visual form of a document or a typeface without relying on meaningful content.
        Lorem ipsum may be used as a placeholder before the final copy is available.</p>
    </div>
  </div>
</template>