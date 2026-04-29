<script setup lang="ts">
import { NButton, NLayout, NLayoutSider, NLayoutContent, NLayoutHeader } from 'naive-ui'
import Sider from '@/components/layout/Sider.vue'
import Content from '@/components/layout/Content.vue'
import { computed, onMounted } from 'vue'
import { useGlobalStore } from '@/stores/global'
import { useSesssionStore } from '@/stores/session'
import { logoutReq } from '@/services'
import { useRouter } from 'vue-router'

const globalStore = useGlobalStore()
const sessionStore = useSesssionStore()
const router = useRouter()
const siteTitle = computed(() => globalStore.options.title || 'Simple Blog')
const siteMark = computed(() => siteTitle.value.slice(0, 1).toUpperCase())
const userName = computed(() => sessionStore.user?.name || 'Editor')

onMounted(() => {
  globalStore.getOptions()
})

const handleLogout = async () => {
  const { code } = await logoutReq()
  if (!code) {
    sessionStore.clearUser()
    window.$message.success('Logout success')
    router.push('/login')
  }
}
</script>

<template>
<n-layout class="admin-shell h-screen">
  <n-layout-header bordered class="admin-shell__header flex flex-row items-center">
    <div class="admin-brand">
      <div class="admin-brand__mark">{{ siteMark }}</div>
      <div class="admin-brand__copy">
        <span class="admin-brand__title">{{ siteTitle }}</span>
        <span class="admin-brand__subtitle">Publishing workspace</span>
      </div>
    </div>
    <div class="admin-shell__header-actions">
      <div class="admin-user">
        <span class="admin-user__label">Signed in</span>
        <span class="admin-user__name">{{ userName }}</span>
      </div>
      <n-button tertiary type="primary" @click="handleLogout">Logout</n-button>
    </div>
  </n-layout-header>
  <n-layout has-sider class="admin-shell__body">
    <n-layout-sider
      bordered
      class="admin-shell__sider"
      :width="236"
    >
      <div class="admin-shell__nav">
        <sider />
      </div>
    </n-layout-sider>
    <n-layout-content class="admin-shell__content">
      <content />
    </n-layout-content>
  </n-layout>
</n-layout>
</template>
