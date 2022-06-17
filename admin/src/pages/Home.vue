<script setup lang="ts">
import { getSystemInfoReq, ISystemInfoResponse } from '@/services'
import { onMounted, ref } from 'vue'
import { NCard, NTag, NStatistic } from 'naive-ui';

const systemInfo = ref()

onMounted(() => {
    loadData()
})

const loadData = async () => {
    let { data } = await getSystemInfoReq()
    if (data) {
        systemInfo.value = data
    }
}
</script>

<template>
<div v-if="systemInfo">
  <n-card title="Summary">
    <div class="flex gap-6">
        <n-statistic
            label="Posts"
            :value="systemInfo.count.posts" />
        <n-statistic
            label="Categories"
            :value="systemInfo.count.cates" />
        <n-statistic
            label="Tags"
            :value="systemInfo.count.tags" />
    </div>
  </n-card>
  <n-card title="System">
    <n-tag>System: {{systemInfo.versions.platform}}</n-tag>
    <n-tag class="ml-1">MySQL: {{systemInfo.versions.mysqlVersion}}</n-tag>
  </n-card>
</div>
</template>

<style>

  
</style>
