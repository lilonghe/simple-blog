<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { NLayoutHeader } from 'naive-ui';
const route = useRoute()
</script>
<template>
<router-view :key="route.fullPath"></router-view>
</template>