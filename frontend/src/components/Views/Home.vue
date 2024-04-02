<script setup>
import Auction from '../ShortAuction.vue'
</script>

<template>
    <section class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-2 gap-y-7 gap-x-7 mt-10 mb-5 px-10 lg:px-32">
        <auction :startingPrice="2000" domainName="best-seller" topLevel="net" endsAt="2024-03-30 20:07:00"></auction>
        <auction :startingPrice="150" domainName="ilovecomupters" topLevel="com" endsAt="2024-04-04 16:00:00"></auction>
        <auction :startingPrice="30000" domainName="auto" topLevel="ai" endsAt="2024-04-06 10:00:00"></auction>
        <auction :startingPrice="35000" domainName="inter" topLevel="net" endsAt="2024-04-06 12:00:00"></auction>
    </section>
</template>