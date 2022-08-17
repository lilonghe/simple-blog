<script setup lang="ts">
import { getVisitListReq } from '@/services'
import { onMounted, ref, computed, watch, h } from 'vue'
import { NDataTable, NInput } from 'naive-ui'
import { formatTime, debounce } from '@/utils'
import { useGlobalStore } from '@/stores/global'

const globalStore = useGlobalStore()
const paginationParams = ref({ pageSize: 10, page: 1,})
const filterParams = ref({ keyword: undefined })
const dataList = ref([])
const total = ref(0)
const pagination = computed(() => {
  return {
    ...paginationParams.value,
    itemCount: total.value,
  }
})
const loading = ref(false)

const loadData = async () => {
  loading.value = true
  let { data } = await getVisitListReq({...paginationParams.value, ...filterParams.value});
  if (data) {
    const { list, total: totalCount } = data;
    dataList.value = list
    total.value = totalCount
  }
  loading.value = false
}

onMounted(() => {
  loadData()
})

watch(filterParams, (now, old) => {
  if (now.keyword !== old.keyword) {
    paginationParams.value.page = 1
    debounce(loadData)
  } else {
    handlePageChange(1)
  }
}, { deep: true })

const columns = [
  {
    title: 'Create Time',
    key: 'create_time',
    render: (row: any) => formatTime(row.create_time),
    className: 'nowrap',
  },
  {
    title: 'Post Path',
    key: 'pathname',
    render: (row: any) => h(
      'a', 
      { href: globalStore.options.site_url + '/post/' + row.pathname + '.html', class: 'link', target: '_blank' },
      { default: () => row.pathname },
    ),
  },
  {
    title: 'User Agent',
    key: 'user_agent',
  },
  {
    title: 'IP',
    key: 'ip',
    className: 'nowrap',
  },
]

const handlePageChange = (page: number) => {
  paginationParams.value.page = page
  loadData()
}
</script>

<template>
<div>
  <div class="mb-2 flex gap-4">
    <div class="w-1/4" >
      <n-input 
        placeholder="Search User Agent" 
        v-model:value="filterParams.keyword" />
    </div>
  </div>
  <n-data-table
    :loading="loading"
    :remote="true"
    :data="dataList"
    :columns="columns"
    :pagination="pagination"
    @update:page="handlePageChange" />
    
</div>
</template>

<style>

  
</style>
