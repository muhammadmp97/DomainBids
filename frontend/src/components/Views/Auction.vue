<script setup>
import { onMounted } from 'vue'
import axios from 'axios'
import Auction from '../Auction.vue'
import Bid from '../Bid.vue'
import { isFuture } from '../../helpers/time.js'

const auction = defineModel('auction', { default: {} })

const id = window.location.hash.slice(4)
onMounted(async () => {
  loadAuction()
})

const loadAuction = () => {
  axios
    .get(`http://127.0.0.1:8000/auctions/${id}`)
    .then(res => {
      if (auction.value.id != undefined) {
        res.data.data.bids = res.data.data.bids.map(bid => {
          bid.is_new = !auction.value.bids.some((oldBid) => bid.id == oldBid.id)
          return bid
        })
      }

      auction.value = res.data.data
    })
    .catch(err => {
      if (err.response.status === 404) {
        location.href = '/#/404'
      }
    })
}

const authToken = localStorage.getItem('db_token')
const placeBid = () => {
  const price = prompt("Enter the new price:")

  if (!price) {
    return
  }

  if (authToken) {
    axios.defaults.headers.common['Authorization'] = `bearer ${authToken}`
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

const isOpen = () => {
  return isFuture(new Date(auction.value.ends_at))
}
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7" v-if="auction">
    <auction :auction="auction"></auction>

    <div id="bids"
      class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-5 gap-x-5 mt-5 | pt-5 border-t-4 border-dotted border-gray-700">
      <bid v-for="(bid, index) of auction.bids" :key="bid.id" :index="index + 1" :bid="bid"></bid>

      <div @click="placeBid()"
        v-show="authToken && isOpen()"
        class="relative overflow-hidden flex justify-center items-center bg-gray-600 rounded shadow-md | transition duration-300 hover:shadow-lg hover:bg-white/25 hover:scale-105 active:bg-gray-700 cursor-pointer"
        style="min-height:60px;">
        <span class="text-white font-bold tracking-wide">Place a bid!</span>
      </div>
    </div>

    <p v-show="authToken && isOpen()" class="text-gray-500 mt-5">By participating in an auction, your email will be shared with the starter.</p>
  </div>
</template>