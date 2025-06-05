<script setup lang="ts">
import { deletePostReq, getPostListReq } from '@/services'
import { debounce, formatTime } from '@/utils'
import { NButton, NDataTable, NInput, NPopconfirm, NSelect, NTag } from 'naive-ui'
import { computed, h, onMounted, ref, watch } from 'vue'
import { RouterLink } from 'vue-router'

const paginationParams = ref({ pageSize: 10, page: 1,})
const filterParams = ref({ status: undefined, is_public: undefined, keyword: undefined })
const postList = ref([])
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
  let { data } = await getPostListReq({...paginationParams.value, ...filterParams.value, type: 1 });
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
            default: () => row.status === 3 ? 'Published' : 'Draft'
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
    title: 'Visit',
    key: 'visit_count',
  },
  {
    title: 'Action',
    key: 'action',
    render: (row: any) => {
      return h(
        'div',
        {
          class: 'flex gap-4',
        },
        [
          h(
            RouterLink,
            {
              to: '/page/edit/' + row.id,
              class: 'link'
            },
            {
              default: () => 'Edit',
            }
          ),
          h(
            NPopconfirm,
            {
              onPositiveClick: () => handleDelete(row.id),

            },
            {
              trigger: () => h(
                'a',
                {
                  class: 'link cursor-pointer',
                }, {
                  default: () => 'Delete',
              }),
              default: () => `Are you sure to delete ${row.title}?`,
            }
          )
        ]
      )
    }
  }
]

const statusList = [
  {
    label: 'Published',
    value: '3'
  },
  {
    label: 'Draft',
    value: '0'
  },
]

const visibilityList = [
  {
    label: 'Public',
    value: 'true'
  },
  {
    label: 'Viewable With Link(Private)',
    value: 'false'
  }
]

const handlePageChange = (page: number) => {
  paginationParams.value.page = page
  loadData()
}

const handleDelete = async (id: number) => {
  let { code } = await deletePostReq(id)
  if (!code) {
    window.$message.success('Delete success')
    if (postList.value.length === 1) {
      paginationParams.value.page--
    }
    loadData()
  }
}

</script>

<template>
<div>
  <div class="mb-2 flex gap-4">
    <div class="w-1/4" >
      <n-input 
        placeholder="Search title" 
        v-model:value="filterParams.keyword" />
    </div>
    <div class="w-1/4" >
      <n-select 
        placeholder="Filter post status" 
        clearable 
        v-model:value="filterParams.status" 
        :options="statusList" />
    </div>
    <div class="w-1/4" >
      <n-select 
        placeholder="Filter post visibility" 
        clearable 
        v-model:value="filterParams.is_public" 
        :options="visibilityList" />
    </div>
  </div>
  <div class="mb-2">
    <router-link to="/page/create"><n-button type="primary">Create</n-button></router-link>
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
