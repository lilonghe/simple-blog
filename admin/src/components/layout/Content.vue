<script setup lang="ts">
import { RouterView, useRoute, RouterLink } from 'vue-router';
import BackIcon from '@/assets/back.svg'
const route = useRoute()
</script>
<template>
<div v-if="route.meta.title" class="flex items-center mt-1 mb-3 gap-2">
    <router-link v-if="route.meta.back" :to="(route.meta.back as string)" class="flex">
        <img :src="BackIcon" class="w-[20px] h-[20px] cursor-pointer" />
    </router-link>
    <h2 class="font-normal m-0">{{ route.meta.title }}</h2>
</div>
<router-view :key="route.fullPath"></router-view>
</template>