<script setup>
import { onMounted } from 'vue';
import axios from 'axios'

const sld = defineModel('sld', { default: '' })
const tld = defineModel('tld', { default: '' })
const description = defineModel('description', { default: '' })
const startingPrice = defineModel('startingPrice', { default: 200 })

onMounted(async () => {
  if (localStorage.getItem('db_token')) {
    axios.defaults.headers.common['Authorization'] = `bearer ${localStorage.getItem('db_token')}`
  }

  axios
    .get(`http://127.0.0.1:8000/auth`)
    .catch(err => {
      if (err.response.status === 401) {
        location.href = '/#/login'
      }
    })
})

const validateForm = () => {
  if (sld.value.length < 3) {
    return false
  }

  if (tld.value.length < 2) {
    return false
  }

  if (startingPrice.value < 200) {
    return false
  }

  return true
}

const confirm = async () => {
  axios
    .post(`http://127.0.0.1:8000/auctions`, {sld: sld.value, tld: tld.value, starting_price: +startingPrice.value, description: description.value})
    .then(res => {
      let gameId = res.data.data.id
      location.href = `/#/a/${gameId}`
    })
    .catch(err => {
      // validation stuff
    })
}
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7">
    <div class="bg-gray-700 rounded shadow-md px-4 py-3">
      <h2 class="font-bold text-lg text-white mb-1">Start an auction</h2>
      <div class="">
        <div class="w-100">
          <input class="input w-60" type="text" placeholder="example" v-model="sld">
          <span class="text-white font-bold ml-2">.</span>
          <input class="input w-20 ml-2 mt-3 text-center" type="text" placeholder="com" v-model="tld">
        </div>

        <div class="w-100 mt-3 relative">
          <span class="starting-price__dollar | text-white/50 select-none">$</span>
          <input class="starting-price__input | w-40 input" type="text" placeholder="200" v-model="startingPrice">
        </div>

        <div class="w-100 mt-3">
          <textarea class="input w-100" rows="5" maxlength="1000" placeholder="Description here" v-model="description"></textarea>
        </div>

        <div class="w-100 mt-5">
          <input class="button" type="button" value="Confirm" @click="confirm" :disabled="!validateForm()">
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.starting-price__dollar {
  position: absolute;
  top: 6px;
  left: 10px;
}

.starting-price__input {
  padding-left: 25px;
}
</style>