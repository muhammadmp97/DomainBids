<script setup>
import Bid from './components/Bid.vue';
</script>

<template>
  <header class="flex items-center justify-between bg-black/20 px-10 lg:px-32 py-5">
    <a href="">
      <div class="logo select-none text-lg font-black text-white">Domain<span class="bg-white text-gray-800 rounded-sm">Bids</span></div>
    </a>
    <nav>
      <ul class="text-white">
        <li class="float-left mr-7 duration-300 hover:text-red-500"><a href="#">Submit a bid</a></li>
        <li class="float-left mr-7 duration-300 hover:text-red-500"><a href="#">About</a></li>
        <li class="float-left duration-300 hover:text-red-500"><a href="#">Contact</a></li>
      </ul>
    </nav>
  </header>

  <section id="Bids"
    class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-4 gap-x-10 mt-10 mb-5 px-10 lg:px-32">
    <bid :startingPrice="2000" domainName="best-seller" topLevel="net" endsAt="2024-04-03 12:00:00"></bid>
    <bid :startingPrice="150" domainName="ilovecomupters" topLevel="com" endsAt="2024-04-04 16:00:00"></bid>
    <bid :startingPrice="30000" domainName="auto" topLevel="ai" endsAt="2024-04-06 10:00:00"></bid>
  </section>
</template>

<style scoped></style>
