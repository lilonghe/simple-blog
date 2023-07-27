<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
const route = useRoute()
</script>
<template>
<h2 v-if="route.meta.title" class="font-normal mt-1 mb-3">{{ route.meta.title }}</h2>
<router-view :key="route.fullPath"></router-view>
</template>