<script setup>
import { computed, ref } from 'vue';
import axios from 'axios'

const form = ref({
  sld: '',
  tld: '',
  description: '',
  starting_price: null,
})

const estimatedPrice = ref()

const key = computed(() => {
  return localStorage.getItem('db_token')?.substring(0, 7)
})

const validateForm = () => {
  if (form.value.sld.length < 3) {
    return false
  }

  if (form.value.tld.length < 2) {
    return false
  }

  if (form.value.starting_price < 200) {
    return false
  }

  return true
}

const estimatePrice = () => {
  if (!form.value.sld || !form.value.tld) {
    return
  }

  axios
    .get(`http://127.0.0.1:8000/estimate-price/${form.value.sld}.${form.value.tld}`)
    .then(res => {
      estimatedPrice.value = res.data.data.price
      if (estimatedPrice.value < 200) {
        estimatedPrice.value = 200
      }
    })
}

const confirm = async () => {
  axios
    .post(`http://127.0.0.1:8000/auctions`, form.value)
    .then(res => {
      let gameId = res.data.data.id
      location.href = `/#/a/${gameId}`
    })
    .catch(err => {
        alert(err.response.data.error)
    })
}
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7">
    <div class="bg-gray-700 rounded shadow-md px-4 py-3">
      <h2 class="font-bold text-lg text-white mb-1">Start an auction</h2>
      <div class="">
        <div class="w-100">
          <input class="input w-60" type="text" placeholder="example" v-model.trim="form.sld" @blur="estimatePrice()">
          <span class="text-white font-bold ml-2">.</span>
          <input class="input w-20 ml-2 mt-3 text-center" type="text" placeholder="com" v-model.trim="form.tld" @blur="estimatePrice()">
        </div>

        <div class="w-100 mt-1">
          <p class="text-sm text-white/50">You should have this TXT record to be allowed to start an auction: <span class="bg-white/50 text-gray-800 rounded">{{ key }}</span>.</p>
        </div>

        <div class="w-100 mt-3 relative">
          <span class="starting-price__dollar | text-white/50 select-none">$</span>
          <input class="starting-price__input | w-40 input" type="text" :placeholder="estimatedPrice || 200" v-model.number="form.starting_price" :disabled="estimatedPrice === null">
        </div>

        <div class="w-100 mt-4">
          <textarea class="input w-100" rows="5" maxlength="1000" placeholder="Description here" v-model.trim="form.description"></textarea>
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