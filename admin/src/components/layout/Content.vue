<script setup lang="ts">
import { computed } from 'vue'
import { RouterView, useRoute, RouterLink } from 'vue-router'
import BackIcon from '@/assets/back.svg'
const route = useRoute()
const pageSection = computed(() => (route.meta.section as string) || 'Workspace')
const pageDescription = computed(() => route.meta.description as string | undefined)
</script>
<template>
<section class="page-shell">
  <header v-if="route.meta.title" class="page-heading">
    <router-link v-if="route.meta.back" :to="route.meta.back as string" class="page-heading__back">
      <img :src="BackIcon" class="h-[18px] w-[18px]" />
    </router-link>
    <div class="page-heading__copy">
      <p class="page-eyebrow">{{ pageSection }}</p>
      <h1 class="page-title">{{ route.meta.title }}</h1>
      <p v-if="pageDescription" class="page-description">{{ pageDescription }}</p>
    </div>
  </header>
  <router-view :key="route.fullPath"></router-view>
</section>
</template>
