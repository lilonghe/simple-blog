<script setup lang="ts">
import { deletePostReq, getPostListReq } from '@/services'
import { debounce, formatTime } from '@/utils'
import { NButton, NDataTable, NInput, NPopconfirm, NSelect, NTag } from 'naive-ui'
import { computed, h, onMounted, ref, watch } from 'vue'
import { RouterLink } from 'vue-router'

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
  const { data } = await getPostListReq({
    ...paginationParams.value,
    ...filterParams.value,
    type: 1,
  })
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

const columns = [
  {
    title: 'Title',
    key: 'title',
    render: (row: any) =>
      h('div', { class: 'table-title-cell' }, [
        h('span', { class: 'link' }, row.title),
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
            to: '/page/edit/' + row.id,
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
          <p class="section-kicker">Evergreen</p>
          <h2 class="section-title">{{ total || 0 }} standalone pages</h2>
          <p class="section-copy">
            Keep long-lived pages separate from posts while using the same editing rhythm.
          </p>
        </div>
        <router-link to="/page/create">
          <n-button type="primary">Create page</n-button>
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
