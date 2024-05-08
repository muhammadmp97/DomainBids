<script setup>
import { computed } from 'vue'
import Timer from './Timer.vue'
import { isFuture } from '../helpers/time'

const props = defineProps(['id', 'domainName', 'topLevel', 'startingPrice', 'endsAt'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.startingPrice)
})

const isClosed = !isFuture(props.endsAt)
</script>

<template>
  <a :href="'#/a/' + id">
    <div class="bg-gray-700 rounded shadow-md duration-300 hover:scale-105 hover:shadow-xl px-4 py-3">
      <div class="flex items-center justify-between">
        <div>
          <div :class="{ 'line-through': isClosed }" class="text-white font-bold tracking-wide'">
            {{ domainName }}
            <span class="text-red-500 tracking-widest pl-0.5">.{{ topLevel }}</span>
          </div>
          <div class="text-gray-400 text-sm mt-0.5">From ${{ startingPrice }}</div>
        </div>
        <timer :time="endsAt"></timer>
      </div>
    </div>
  </a>
</template>