<script setup lang="ts">
import { NDataTable, NPopconfirm, NInput, NButton } from 'naive-ui'
import { computed, h, onMounted, ref } from 'vue'
import { useGlobalStore } from '@/stores/global'
import { RouterLink } from 'vue-router'

const globalStore = useGlobalStore()
const keyword = ref()

onMounted(() => {
  globalStore.getCates()
})

const handleDelete = (id: number) => {
  globalStore.deleteCate(id)
}

const columns = [
  {
    title: 'Name',
    key: 'name',
  },
  {
    title: 'Path',
    key: 'pathname',
  },
  {
    title: 'Posts',
    key: 'post_count',
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
            to: '/category/edit/' + row.id,
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
            default: () => `Delete "${row.name}"?`,
          },
        ),
      ]),
  },
]

const list = computed(() => {
  if (keyword.value) {
    return globalStore.cates.filter((cate: any) => {
      return cate.name.includes(keyword.value) || cate.pathname.includes(keyword.value)
    })
  }
  return globalStore.cates
})
</script>

<template>
  <div class="page-stack">
    <section class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Structure</p>
          <h2 class="section-title">{{ list.length }} category labels in use</h2>
          <p class="section-copy">
            Categories define the broad reading map, so keep them sparse and easy to scan.
          </p>
        </div>
        <router-link to="/category/create">
          <n-button type="primary">Create category</n-button>
        </router-link>
      </div>

      <div class="filter-row">
        <div class="filter-field filter-field--wide">
          <n-input v-model:value="keyword" placeholder="Search name or path" />
        </div>
      </div>

      <div class="data-table-wrap">
        <n-data-table :columns="columns" :data="list" :pagination="{ pageSize: 10 }" striped />
      </div>
    </section>
  </div>
</template>
