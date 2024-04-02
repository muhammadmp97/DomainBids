<script setup>
import Auction from '../Auction.vue'
import Bid from '../Bid.vue'
</script>

<template>
  <div class="container mx-auto px-5 pt-7">
    <auction :startingPrice="30000" domainName="auto" topLevel="ai" endsAt="2024-04-06 10:00:00"></auction>

    <div id="bids" class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-5 gap-x-5 mt-5 | pt-5 border-t-4 border-dotted border-gray-700">
      <bid v-for="index in 4" :key="index" :index="index" name="John Doe" :price="2500 * index" createdAt="2024-03-04 20:25"></bid>

      <div class="relative overflow-hidden flex justify-center items-center bg-gray-600 rounded shadow-md | transition duration-300 hover:shadow-lg hover:bg-white/25 hover:scale-105 active:bg-gray-700 cursor-pointer">
        <span class="text-white font-bold tracking-wide">Place a bid!</span>
      </div>
    </div>
  </div>
</template>