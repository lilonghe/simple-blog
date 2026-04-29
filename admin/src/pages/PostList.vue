<script setup lang="ts">
import { getPostListReq, deletePostReq } from '@/services'
import { onMounted, ref, h, computed, watch } from 'vue'
import { NDataTable, NTag, NInput, NPopconfirm, NSelect, NButton } from 'naive-ui'
import { RouterLink } from 'vue-router'
import { formatTime, debounce } from '@/utils'

const paginationParams = ref({ pageSize: 10, page: 1 })
const filterParams = ref({ status: undefined, is_public: undefined, keyword: undefined })
const postList = ref<any[]>([])
const total = ref(0)
const loading = ref(false)

const pagination = computed(() => ({
  ...paginationParams.value,
  itemCount: total.value,
}))

const loadData = async () => {
  loading.value = true
  const { data } = await getPostListReq({ ...paginationParams.value, ...filterParams.value })
  if (data) {
    postList.value = data.list
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

const handleViewPost = (pathname: string) => {
  window.open('/post/' + pathname + '.html')
}

const columns = [
  {
    title: 'Title',
    key: 'title',
    render: (row: any) =>
      h('div', { class: 'table-title-cell' }, [
        h(
          'a',
          {
            onClick: () => row.status !== 0 && handleViewPost(row.pathname),
            class: row.status !== 0 ? 'link cursor-pointer' : 'table-link-disabled',
          },
          { default: () => row.title },
        ),
        h('span', { class: 'table-title-cell__meta' }, `/post/${row.pathname}.html`),
      ]),
  },
  {
    title: 'Status',
    key: 'status',
    width: 110,
    render: (row: any) =>
      h(
        NTag,
        {
          type: row.status === 3 ? 'success' : 'default',
        },
        {
          default: () => (row.status === 3 ? 'Published' : 'Draft'),
        },
      ),
  },
  {
    title: 'Visibility',
    key: 'is_public',
    width: 110,
    render: (row: any) => (row.is_public ? 'Public' : 'Private'),
  },
  {
    title: 'Published',
    key: 'create_time',
    width: 180,
    render: (row: any) => formatTime(row.create_time),
  },
  {
    title: 'Visits',
    key: 'visit_count',
    width: 90,
  },
  {
    title: 'Action',
    key: 'action',
    width: 140,
    render: (row: any) =>
      h('div', { class: 'table-actions' }, [
        h(
          RouterLink,
          {
            to: '/post/edit/' + row.id,
            class: 'table-action',
          },
          {
            default: () => 'Edit',
          },
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: () => handleDelete(row.id),
          },
          {
            trigger: () =>
              h(
                'a',
                {
                  class: 'table-action table-action--danger cursor-pointer',
                },
                { default: () => 'Delete' },
              ),
            default: () => `Delete "${row.title}"?`,
          },
        ),
      ]),
  },
]

const statusList = [
  {
    label: 'Published',
    value: '3',
  },
  {
    label: 'Draft',
    value: '0',
  },
]

const visibilityList = [
  {
    label: 'Public',
    value: 'true',
  },
  {
    label: 'Private link only',
    value: 'false',
  },
]

const handlePageChange = (page: number) => {
  paginationParams.value.page = page
  loadData()
}

const handleDelete = async (id: number) => {
  const { code } = await deletePostReq(id)
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
  <div class="page-stack">
    <section class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Library</p>
          <h2 class="section-title">{{ total || 0 }} posts in the archive</h2>
          <p class="section-copy">
            Keep titles, visibility and publishing state under control from one consistent table.
          </p>
        </div>
        <router-link to="/post/create">
          <n-button type="primary">Create post</n-button>
        </router-link>
      </div>

      <div class="filter-row">
        <div class="filter-field filter-field--wide">
          <n-input v-model:value="filterParams.keyword" placeholder="Search title" />
        </div>
        <div class="filter-field">
          <n-select
            v-model:value="filterParams.status"
            clearable
            placeholder="Status"
            :options="statusList"
          />
        </div>
        <div class="filter-field">
          <n-select
            v-model:value="filterParams.is_public"
            clearable
            placeholder="Visibility"
            :options="visibilityList"
          />
        </div>
      </div>

      <div class="data-table-wrap">
        <n-data-table
          :loading="loading"
          :remote="true"
          :data="postList"
          :columns="columns"
          :pagination="pagination"
          striped
          @update:page="handlePageChange"
        />
      </div>
    </section>
  </div>
</template>
