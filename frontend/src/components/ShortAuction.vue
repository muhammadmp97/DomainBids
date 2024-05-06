<script setup>
import { onMounted, ref } from 'vue';
import { computed } from 'vue';

const props = defineProps(['id', 'domainName', 'topLevel', 'startingPrice', 'endsAt'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.startingPrice)
})

const timer = ref({
  days: 999,
  hours: 0,
  minutes: 0,
  seconds: 0
})

onMounted(() => {
  var delay = 5
  let intervalId = setInterval(() => {
    let diff = (new Date(`${props.endsAt} UTC`) - new Date()) / 1000
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
</script>

<template>
  <a :href="'#/a/' + id">
    <div class="bg-gray-700 rounded shadow-md duration-300 hover:scale-105 hover:shadow-xl px-4 py-3">
      <div class="flex items-center justify-between">
        <div>
          <div :class="{ 'line-through': timer.seconds < 0 }" class="text-white font-bold tracking-wide'">
            {{ domainName }}
            <span class="text-red-500 tracking-widest pl-0.5">.{{ topLevel }}</span>
          </div>
          <div class="text-gray-400 text-sm mt-0.5">From ${{ startingPrice }}</div>
        </div>
        <div class="bg-red-800/70 px-2 py-2 rounded select-none" v-show="timer.days != 999">
          <div class="text-red-300 text-xs" v-if="timer.seconds >= 0">
            {{ timer.days }}d {{ timer.hours }}:{{ timer.minutes }}:{{ timer.seconds }}
          </div>
          <div class="text-red-300 text-xs uppercase tracking-widest" v-if="timer.seconds < 0">closed</div>
        </div>
      </div>
    </div>
  </a>
</template>