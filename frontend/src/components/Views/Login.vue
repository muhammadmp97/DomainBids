<script setup>
import { onMounted, ref } from 'vue';
import axios from 'axios'
import { Store } from '../../store'

const form = ref({
  username: '',
  password: '',
})

onMounted(async () => {
  if (Store.authenticated) {
    location.href = '/'
  }
})

const login = async () => {
  axios
    .post(`http://127.0.0.1:8000/auth`, form.value)
    .then(res => {
      localStorage.setItem("db_token", res.data.data.token)
      alert("Welcome here!")
      location.href = '/'
    })
    .catch(err => {
      if (err.response.status === 401) {
        alert(err.response.data.error)
      }
    })
}
</script>

<template>
  <div class="container lg:lg-container mx-auto pt-7">
    <div class="bg-gray-700 rounded shadow-md px-4 py-3">
      <div class="flex items-center justify-between">
        <div class="w-80">
          <input class="input" type="text" placeholder="username" v-model="form.username">
          <input class="input mt-3" type="password" placeholder="password" v-model="form.password">
          <input class="button mt-3" type="button" value="Login/Register" @click="login" :disabled="form.username.length === 0 || form.password.length === 0">
        </div>
      </div>
    </div>
  </div>
</template>