<script setup lang="ts">
import { useSesssionStore } from './stores/session'
import { useRouter } from 'vue-router'

const sessionStore = useSesssionStore()
const router = useRouter()

if (sessionStore.user === null) {
  router.push('/login');
}
</script>

<template>
  <router-view></router-view>
</template>

<style>

  
</style>
