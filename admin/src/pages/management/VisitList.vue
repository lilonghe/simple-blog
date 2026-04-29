<script setup lang="ts">
import { getVisitListReq } from '@/services'
import { onMounted, ref, computed, watch, h } from 'vue'
import { NDataTable, NInput } from 'naive-ui'
import { formatTime, debounce } from '@/utils'
import { useGlobalStore } from '@/stores/global'

const globalStore = useGlobalStore()
const paginationParams = ref({ pageSize: 10, page: 1 })
const filterParams = ref({ keyword: undefined })
const dataList = ref<any[]>([])
const total = ref(0)
const loading = ref(false)

const pagination = computed(() => ({
  ...paginationParams.value,
  itemCount: total.value,
}))

const loadData = async () => {
  loading.value = true
  const { data } = await getVisitListReq({ ...paginationParams.value, ...filterParams.value })
  if (data) {
    dataList.value = data.list
    total.value = data.total
  }
  loading.value = false
}

onMounted(() => {
  loadData()
})

watch(
  filterParams,
  (now, old) => {
    if (now.keyword !== old.keyword) {
      paginationParams.value.page = 1
      debounce(loadData)
    } else {
      handlePageChange(1)
    }
  },
  { deep: true },
)

const columns = [
  {
    title: 'Visited',
    key: 'create_time',
    width: 180,
    render: (row: any) => formatTime(row.create_time),
    className: 'nowrap',
  },
  {
    title: 'Path',
    key: 'pathname',
    render: (row: any) =>
      h(
        'a',
        {
          href: globalStore.options.site_url + '/post/' + row.pathname + '.html',
          class: 'link',
          target: '_blank',
        },
        { default: () => row.pathname },
      ),
  },
  {
    title: 'User agent',
    key: 'user_agent',
  },
  {
    title: 'IP',
    key: 'ip',
    width: 150,
    className: 'nowrap',
  },
]

const handlePageChange = (page: number) => {
  paginationParams.value.page = page
  loadData()
}
</script>

<template>
  <div class="page-stack">
    <section class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Traffic trace</p>
          <h2 class="section-title">{{ total || 0 }} visit records</h2>
          <p class="section-copy">
            Search by user agent and inspect what content paths are drawing real attention.
          </p>
        </div>
      </div>

      <div class="filter-row">
        <div class="filter-field filter-field--wide">
          <n-input v-model:value="filterParams.keyword" placeholder="Search user agent" />
        </div>
      </div>

      <div class="data-table-wrap">
        <n-data-table
          :loading="loading"
          :remote="true"
          :data="dataList"
          :columns="columns"
          :pagination="pagination"
          striped
          @update:page="handlePageChange"
        />
      </div>
    </section>
  </div>
</template>
