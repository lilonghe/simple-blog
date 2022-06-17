<script setup lang="ts">
import { useSesssionStore } from './stores/session'
import { useRouter } from 'vue-router'
import { createDiscreteApi } from 'naive-ui'

const sessionStore = useSesssionStore()
const router = useRouter()

const { message } = createDiscreteApi(["message"],{})
window.$message = message

if (sessionStore.user === null) {
  router.push('/login');
}
</script>

<template>
<n-message-provider>
  <MessageAPI />
  <router-view></router-view>
</n-message-provider>
</template>
