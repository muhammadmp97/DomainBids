<script setup>
import { onMounted, ref } from 'vue'
import Auction from '../ShortAuction.vue'
import axios from 'axios'

const auctions = ref()

onMounted(function () {
  axios
    .get('http://127.0.0.1:8000/auctions')
    .then(res => {
      auctions.value = res.data.data
    })
})
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7">
    <section class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-7 gap-x-7 mt-10 mb-5">
      <auction v-for="auction of auctions" v-bind:key="auction.id" :auction="auction"></auction>
    </section>
  </div>
</template>