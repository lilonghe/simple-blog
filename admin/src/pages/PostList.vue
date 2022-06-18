<script setup lang="ts">
import { getPostListReq } from '@/services'
import { onMounted, ref, h, computed, watchEffect, watch } from 'vue'
import { NDataTable, NTag, NInput } from 'naive-ui'
import { RouterLink } from 'vue-router'
import { formatTime, debounce } from '@/utils'

const params = ref({ pageSize: 10, page: 1, status: '', keyword: undefined })
const postList = ref([])
const total = ref(0)
const pagination = computed(() => {
  return {
    ...params.value,
    itemCount: total.value,
  }
})
const loading = ref(false)

const loadData = async () => {
  loading.value = true
  let { data } = await getPostListReq(params.value);
  if (data) {
    const { list, total: totalCount } = data;
    postList.value = list
    total.value = totalCount
  }
  loading.value = false
}

onMounted(() => {
  loadData()
})

watch(() => params.value.keyword, (n, old) => {
  if (n !== old) {
    params.value.page = 1
    debounce(loadData)
  }
})

const columns = [
  {
    title: 'Title',
    key: 'title',
  },
  {
    title: 'Status',
    key: 'status',
    render: (row: any) => {
      return h(
          NTag, 
          {
            type: row.status === 3 ? 'success' : 'default',
          },
          {
            default: row.status === 3 ? 'Published' : 'Draft'
          }
        ) 
    }
  },
  {
    title: 'Visibility',
    key: 'is_public',
    render: (row: any) => {
      return row.is_public ? 'Public' : 'Private'
    }
  },
  {
    title: 'Publish time',
    key: 'create_time',
    render: (row: any) => formatTime(row.create_time)
  },
  {
    title: 'Action',
    key: 'action',
    render: (row: any) => {
      return h(
        RouterLink,
        {
          to: '/post/edit/' + row.id,
          class: 'link'
        },
        {
          default: 'Edit',
        }
      )
    }
  }
]

const handlePageChange = (page: number) => {
  params.value.page = page
  loadData()
}

</script>

<template>
<div>
  <div class="mb-2 w-1/3">
    <n-input placeholder="Search" v-model:value="params.keyword" />
  </div>
  <n-data-table
    :loading="loading"
    :remote="true"
    :data="postList"
    :columns="columns"
    :pagination="pagination"
    @update:page="handlePageChange" />
    
</div>
</template>

<style>

  
</style>
