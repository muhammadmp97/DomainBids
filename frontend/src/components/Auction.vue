<script setup>
import { computed } from 'vue'
import Timer from './Timer.vue'
import { isFuture } from '../helpers/time'

const props = defineProps(['auction'])

const startingPrice = computed(() => {
  return new Intl.NumberFormat().format(props.auction.starting_price)
})

const isClosed = !isFuture(new Date(`${props.auction.ends_at} UTC`))
</script>

<template>
  <div class="bg-gray-700 rounded shadow-md px-4 py-3">
    <div class="flex items-center justify-between border-b border-gray-600 pb-2">
      <div>
        <div :class="{ 'line-through': isClosed }" class="text-white font-bold tracking-wide text-lg">
          {{ auction.sld }}
          <span class="text-red-500 tracking-widest pl-0.5">.{{ auction.tld }}</span>
        </div>
        <div class="text-gray-400 text-sm">From ${{ startingPrice }}</div>
      </div>
      <timer :time="auction.ends_at"></timer>
    </div>

    <div>
      <p class="text-white font-light mt-2 mb-1 whitespace-pre-line" v-html="auction.description"></p>
    </div>
  </div>
</template>