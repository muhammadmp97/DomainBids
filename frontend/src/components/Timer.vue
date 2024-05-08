<script setup>
import { onMounted, onUnmounted, ref } from 'vue';

const props = defineProps(['time'])

const timer = ref({
  days: 999,
  hours: 0,
  minutes: 0,
  seconds: 0
})

const updateTimer = () => {
  let diff = (new Date(`${props.time} UTC`) - new Date()) / 1000
  timer.value.days = Math.floor(diff / 86400)
  timer.value.hours = Math.floor(diff / 3600 % 24).toString().padStart(2, '0')
  timer.value.minutes = Math.floor(diff / 60 % 60).toString().padStart(2, '0')
  timer.value.seconds = Math.floor(diff % 60).toString().padStart(2, '0')

  return diff
}

let intervalId = null

onMounted(() => {
  updateTimer()
  intervalId = setInterval(() => {
    if (updateTimer() < 0) {
      clearInterval(intervalId)
    }
  }, 1000)
})

onUnmounted(() => {
  clearInterval(intervalId)
})
</script>

<template>
  <div class="bg-red-800/70 px-2 py-2 rounded select-none">
    <div class="text-red-300 text-xs" v-if="timer.seconds >= 0">
      {{ timer.days }}d {{ timer.hours }}:{{ timer.minutes }}:{{ timer.seconds }}
    </div>
    <div class="text-red-300 text-xs uppercase tracking-widest" v-else>closed</div>
  </div>
</template>