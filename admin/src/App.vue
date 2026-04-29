<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { createDiscreteApi, NConfigProvider, NElement, NGlobalStyle } from 'naive-ui'
import MainLayout from '@/components/layout/Index.vue'
import { useSesssionStore } from './stores/session'
import { useRoute } from 'vue-router'
import { themeOverrides } from './theme'

const sessionStore = useSesssionStore()
const route = useRoute()
const isLoginRoute = computed(() => route.path === '/login')

onMounted(() => {
  const { message } = createDiscreteApi(['message'], {
    configProviderProps: {
      themeOverrides,
    },
  })
  window.$message = message
  sessionStore.getUser()
})
</script>

<template>
<n-config-provider :theme-overrides="themeOverrides">
  <n-global-style />
  <n-element>
    <div class="app-root">
      <div v-if="isLoginRoute" class="app-route">
        <router-view />
      </div>
      <div v-else>
        <main-layout v-if="sessionStore.user" />
        <div v-else class="app-loading">
          <div class="app-loading__dot"></div>
          <p>Loading workspace...</p>
        </div>
      </div>
    </div>
  </n-element>
</n-config-provider>
</template>
