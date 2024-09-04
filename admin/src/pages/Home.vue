<script setup lang="ts">
import { getPostListReq, getSystemInfoReq } from '@/services';
import { SVGGraph } from 'calendar-graph';
import { NCard, NStatistic, NTag } from 'naive-ui';
import { onMounted, ref } from 'vue';

const systemInfo = ref()

onMounted(() => {
    loadData()
    loadAllPost();
})

const loadData = async () => {
    let { data } = await getSystemInfoReq()
    if (data) {
        systemInfo.value = data
    }
}

const loadAllPost = async () => {
  let { data } = await getPostListReq({pageSize: 999999999, page: 1, archive: true});
  if (data) {
    const { list } = data;
    let graphData: any[] = [];
    let current = new Date()
    let year = current.getFullYear().toString();
    let month = current.getMonth() + 1;
    let day = current.getDate();
    list.
      filter((item: any) => new Date(item.create_time).toJSON().startsWith(year)).
      map((item: any) => {
        let day = new Date(item.create_time).toJSON().substring(0, 10);
        let match = graphData.findIndex((d: any) => d.date === day);
        if (match !== -1) {
          graphData[match].count += 1;
        } else {
          graphData.push({
            date: day,
            count: 1,
          })
        }
      });
    new SVGGraph('#svg-root', graphData, {
      startDate: new Date(`${year}-01-01`),
      endDate: new Date(`${year}-${month}-${day}`),
    });
  }
}

</script>

<template>
<div>
  <div id="svg-root"></div>
  <div id="tooltip"></div>
</div>
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
  <n-card title="System" class="mt-3">
    <n-tag>System: {{systemInfo.versions.platform}}</n-tag>
    <n-tag class="ml-1">MySQL: {{systemInfo.versions.mysqlVersion}}</n-tag>
  </n-card>
  <n-card title="Visit top" class="mt-3">
    <div class="flex flex-col gap-2">
      <div v-for="item in systemInfo.count.weekVisitTop" class="flex gap-2 items-center">
        <n-tag size="small">{{ item.count }}</n-tag><a :href="`/post/${item.pathname}.html`" target="_blank" >{{ item.title }}</a>
      </div>
    </div>
    
  </n-card>
</div>
</template>

<style>

  
</style>
