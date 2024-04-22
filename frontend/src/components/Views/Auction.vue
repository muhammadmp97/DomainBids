<script setup>
import { onMounted } from 'vue';
import axios from 'axios'
import Auction from '../Auction.vue'
import Bid from '../Bid.vue'

const auction = defineModel('auction', { default: {} })

const id = window.location.hash.slice(4)
onMounted(async () => {
  loadAuction()
})

const loadAuction = () => {
  axios
    .get(`http://127.0.0.1:8000/auctions/${id}`)
    .then(res => {
      auction.value = res.data.data
    })
    .catch(err => {
      if (err.response.status === 404) {
        location.href = '/#/404'
      }
    })
}

const placeBid = () => {
  const price = prompt("Enter the new price:")

  if (localStorage.getItem('db_token')) {
    axios.defaults.headers.common['Authorization'] = `bearer ${localStorage.getItem('db_token')}`
  }

  axios
    .post(`http://127.0.0.1:8000/auctions/${id}/bids`, { price: +price })
    .then(res => {
      loadAuction()
    })
    .catch(err => {
      alert(err.response.data.error)
    })
}
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7" v-if="auction">
    <auction :auction="auction"></auction>

    <div id="bids"
      class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-5 gap-x-5 mt-5 | pt-5 border-t-4 border-dotted border-gray-700">
      <bid v-for="(bid, index) of auction.bids" :key="bid.id" :index="index + 1" :bid="bid"></bid>

      <div @click="placeBid()"
        class="relative overflow-hidden flex justify-center items-center bg-gray-600 rounded shadow-md | transition duration-300 hover:shadow-lg hover:bg-white/25 hover:scale-105 active:bg-gray-700 cursor-pointer"
        style="min-height:60px;">
        <span class="text-white font-bold tracking-wide">Place a bid!</span>
      </div>
    </div>
  </div>
</template>