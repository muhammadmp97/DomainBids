<script setup>
import { computed, onMounted, ref } from 'vue';
import axios from 'axios'
import { timeForHuman } from '../../helpers/time'

const user = ref()

const id = window.location.hash.slice(4)
onMounted(async () => {
  axios
    .get(`http://127.0.0.1:8000/users/${id}`)
    .then(res => {
      user.value = res.data.data
    })
    .catch(err => {
      if (err.response.status === 404) {
        location.href = '/#/404'
      }
    })
})

let registrationDateForHuman = computed(() => {
  let time = new Date(user.value.created_at)
  return timeForHuman(time)
})
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7">
    <div class="bg-gray-700 rounded shadow-md px-4 py-3">
      <div class="mb-3">
        <h2 class="font-bold text-lg text-white">{{ user.nickname }}</h2>
        <p class="text-sm text-white/75">{{ user.bio }}</p>
      </div>

      <p class="text-xs text-white">Registered {{ registrationDateForHuman }}</p>
    </div>
  </div>
</template>