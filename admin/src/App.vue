<script setup lang="ts">
import { createDiscreteApi, NLayout, NLayoutSider, NLayoutContent } from 'naive-ui'
import Sider from '@/components/layout/sider.vue'
import { getUserReq } from "@/services"
import { onBeforeMount, onMounted } from 'vue'
import { useSesssionStore } from './stores/session'
import { useRouter, useRoute } from 'vue-router'

const sessionStore = useSesssionStore()
const router = useRouter()
const route = useRoute()

onMounted(() => {
  const { message } = createDiscreteApi(["message"],{})
  window.$message = message
  sessionStore.getUser()
})

</script>

<template>
<div v-if="route.fullPath === '/login'">
  <router-view></router-view>
</div>
<div v-else>
  <n-layout has-sider v-if="sessionStore.user">
    <n-layout-sider>
      <sider />
    </n-layout-sider>
    <n-layout-content class="pt-1">
      <router-view></router-view>
    </n-layout-content>
  </n-layout>
  <div v-else>
    loading
  </div>
</div>

</template>
