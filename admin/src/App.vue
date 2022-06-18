<script setup lang="ts">
import { createDiscreteApi, NElement } from 'naive-ui'
import MainLayout from '@/components/layout/Index.vue'
import { onMounted } from 'vue'
import { useSesssionStore } from './stores/session'
import { useRoute } from 'vue-router'

const sessionStore = useSesssionStore()
const route = useRoute()

onMounted(() => {
  const { message } = createDiscreteApi(["message"],{})
  window.$message = message
  sessionStore.getUser()
})

</script>

<template>
<n-element>
  <div v-if="route.fullPath === '/login'">
    <router-view></router-view>
  </div>
  <div v-else>
    <main-layout v-if="sessionStore.user" />
    <div v-else class="flex justify-center mt-20">
      Loading
    </div>
  </div>
</n-element>
</template>
