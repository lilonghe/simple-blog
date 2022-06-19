<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
const route = useRoute()
</script>
<template>
<router-view :key="route.fullPath"></router-view>
</template>