<script setup lang="ts">
import { NLayout, NLayoutSider, NLayoutContent, NLayoutHeader } from 'naive-ui'
import Sider from '@/components/layout/sider.vue'
import Content from '@/components/layout/Content.vue'
import { onMounted } from 'vue'
import { useGlobalStore } from '@/stores/global'
import { useSesssionStore } from '@/stores/session'
import { logoutReq } from '@/services'
import { useRouter } from 'vue-router'

const globalStore = useGlobalStore()
const sessionStore = useSesssionStore()
const router = useRouter()

onMounted(() => {
  globalStore.getOptions()
})

const handleLogout = async () => {
    const { code } = await logoutReq();
    if (!code) {
        window.$message.success('Logout success')
        router.push('/login')
    }
}

</script>

<template>
<n-layout class="h-screen">
    <n-layout-header bordered class="h-12 flex flex-row items-center">
        <div class="ml-3">
            <span class="font-bold text-lg">{{globalStore.options.title}}</span>
        </div>
        <div class="ml-auto mr-4 flex flex-row gap-4">
            <span>{{sessionStore.user?.name}}</span>
            <a class="link cursor-pointer" v-on:click="handleLogout">Logout</a>
        </div>
    </n-layout-header>
    <n-layout has-sider style="height: calc(100vh - 3rem)">
        <n-layout-sider bordered>
            <sider />
        </n-layout-sider>
        <n-layout-content class="pt-2 pb-2 mr-3 ml-3">
            <content />
        </n-layout-content>
    </n-layout>
</n-layout>

</template>
