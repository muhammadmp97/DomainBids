<script setup>
import { ref, computed } from 'vue'
import Home from './components/Views/Home.vue'
import NotFound from './components/Views/NotFound.vue'
import Auction from './components/Views/Auction.vue';
import Login from './components/Views/Login.vue';
import StartAuction from './components/Views/StartAuction.vue';

const routes = {
  '^\/?$': Home,
  '\/a\/[\\d]+': Auction,
  '^\/login$': Login,
  '^\/start-auction$': StartAuction,
}

const currentPath = ref(window.location.hash)

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
})

const currentView = computed(() => {
  for (let route in routes) {
    if (new RegExp(route).test(currentPath.value.slice(1) || '/')) {
      return routes[route]
    }
  }

  return NotFound
})
</script>

<template>
  <div class="flex w-screen relative -z-50 -top-20 justify-center">
    <div class="bg-blue-500 absolute blur-3xl opacity-15 w-2/4 h-32"></div>
  </div>

  <header class="bg-black/20 py-5">
    <div class="container lg:lg-container flex items-center justify-between mx-auto">
      <a href="">
        <div class="logo select-none text-lg font-black text-white">Domain<span class="bg-white text-gray-800 rounded-sm">Bids</span></div>
      </a>
      <nav>
        <ul class="text-white">
          <a href="/#/start-auction">
            <li class="float-left mr-7 duration-300 hover:text-red-500">Start an auction</li>
          </a>
          <a href="#">
            <li class="float-left mr-7 duration-300 hover:text-red-500">About</li>
          </a>
          <a href="#">
            <li class="float-left duration-300 hover:text-red-500">Contact</li>
          </a>
        </ul>
      </nav>
    </div>
  </header>

  <component :is="currentView" />
</template>

<style scoped></style>
