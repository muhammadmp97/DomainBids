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
  <a :href="'#/a/' + auction.id">
    <div class="bg-gray-700 rounded shadow-md duration-300 hover:scale-105 hover:shadow-xl px-4 py-3">
      <div class="flex items-center justify-between">
        <div>
          <div :class="{ 'line-through': isClosed }" class="text-white font-bold tracking-wide'">
            {{ auction.sld }}
            <span class="text-red-500 tracking-widest pl-0.5">.{{ auction.tld }}</span>
          </div>
          <div class="text-gray-400 text-sm mt-0.5">From ${{ auction.starting_price }}</div>
        </div>
        <timer :time="auction.ends_at"></timer>
      </div>
    </div>
  </a>
</template>