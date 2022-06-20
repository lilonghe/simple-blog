<script setup lang="ts">
import { NLayout, NLayoutSider, NLayoutContent } from 'naive-ui'
import Sider from '@/components/layout/sider.vue'
import Content from '@/components/layout/Content.vue'
import { onMounted } from 'vue'
import { useGlobalStore } from '@/stores/global'

const globalStore = useGlobalStore()

onMounted(() => {
  globalStore.getOptions()
})

</script>

<template>
<n-layout has-sider>
    <n-layout-sider bordered class="mr-3 min-h-screen sticky top-0">
        <sider />
    </n-layout-sider>
    <n-layout-content class="pt-2 pb-2 mr-3">
        <content />
    </n-layout-content>
</n-layout>
</template>
