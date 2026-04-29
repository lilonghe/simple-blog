<template>
  <div class="page-stack">
    <section class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Micro notes</p>
          <h2 class="section-title">{{ total || 0 }} whispers collected</h2>
          <p class="section-copy">
            Short thoughts should be quick to write and edit, without interruptive modals.
          </p>
        </div>
        <n-button type="primary" @click="openEdit()">
          {{ isEditorOpen ? 'Editing whisper' : 'New whisper' }}
        </n-button>
      </div>

      <div class="filter-row">
        <div class="filter-field filter-field--wide">
          <n-input v-model:value="filterParams.keyword" placeholder="Search content" />
        </div>
        <div class="filter-field">
          <n-select
            v-model:value="filterParams.is_public"
            clearable
            placeholder="Visibility"
            :options="visibilityOptions"
          />
        </div>
      </div>

      <div v-if="isEditorOpen" class="inline-editor">
        <div class="inline-editor__header">
          <div>
            <h3 class="section-title">{{ editForm.id ? 'Edit whisper' : 'Create whisper' }}</h3>
            <p class="section-copy">
              Keep it brief. These notes work best when they read like clean fragments.
            </p>
          </div>
          <n-button tertiary @click="cancelEdit">Cancel</n-button>
        </div>

        <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="top">
          <n-form-item label="Content" path="content">
            <n-input v-model:value="editForm.content" type="textarea" rows="4" placeholder="Write a short note" />
          </n-form-item>
          <n-form-item label="Visibility" path="is_public" :show-feedback="false">
            <n-switch v-model:value="editForm.is_public">
              <template #checked>Public</template>
              <template #unchecked>Private</template>
            </n-switch>
          </n-form-item>
          <div class="form-actions">
            <n-button @click="cancelEdit">Cancel</n-button>
            <n-button type="primary" @click="save">Save whisper</n-button>
          </div>
        </n-form>
      </div>

      <div class="data-table-wrap">
        <n-data-table
          :loading="loading"
          :remote="true"
          :data="whisperList"
          :columns="columns"
          :pagination="pagination"
          striped
          @update:page="handlePageChange"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted, computed, watch } from 'vue'
import {
  NDataTable,
  NInput,
  NPopconfirm,
  NButton,
  NSwitch,
  NSelect,
  NForm,
  NFormItem,
} from 'naive-ui'
import {
  getWhisperListReq,
  createWhisperReq,
  updateWhisperReq,
  deleteWhisperReq,
  updateWhisperVisibilityReq,
} from '@/services'
import { formatTime, debounce } from '@/utils'
import type { FormInst, FormRules } from 'naive-ui'

interface Whisper {
  id: number
  content: string
  is_public: boolean
  create_time: string
}

const paginationParams = ref({ pageSize: 10, page: 1 })
const filterParams = ref<{ keyword?: string; is_public?: string }>({})
const whisperList = ref<Whisper[]>([])
const total = ref(0)
const loading = ref(false)

const isEditorOpen = ref(false)
const editForm = ref<Partial<Whisper>>({ content: '', is_public: true })

const pagination = computed(() => ({
  page: paginationParams.value.page,
  pageSize: paginationParams.value.pageSize,
  itemCount: total.value,
}))

const formRef = ref<FormInst | null>(null)

const rules: FormRules = {
  content: [{ required: true, message: 'Content cannot be empty.', trigger: ['input', 'blur'] }],
}

const visibilityOptions = [
  { label: 'Public', value: 'true' },
  { label: 'Private', value: 'false' },
]

const loadData = async () => {
  loading.value = true
  const params: any = {
    page: paginationParams.value.page,
    limit: paginationParams.value.pageSize,
    ...filterParams.value,
  }
  const { data } = await getWhisperListReq(params)
  if (data) {
    whisperList.value = data.list
    total.value = data.total
  }
  loading.value = false
}

onMounted(loadData)

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
    title: 'Content',
    key: 'content',
    ellipsis: { tooltip: true },
  },
  {
    title: 'Public',
    key: 'is_public',
    width: 96,
    render: (row: Whisper) =>
      h(NSwitch, {
        value: row.is_public,
        onUpdateValue: (val: boolean) => changePublic(row, val),
      }),
  },
  {
    title: 'Created',
    key: 'create_time',
    render: (row: Whisper) => formatTime(row.create_time),
    width: 180,
  },
  {
    title: 'Action',
    key: 'action',
    width: 160,
    render: (row: Whisper) =>
      h('div', { class: 'table-actions' }, [
        h(
          'a',
          {
            class: 'table-action cursor-pointer',
            onClick: () => openEdit(row),
          },
          { default: () => 'Edit' },
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
            default: () => 'Delete this whisper?',
          },
        ),
      ]),
  },
]

function handlePageChange(page: number) {
  paginationParams.value.page = page
  loadData()
}

function openEdit(row?: Whisper) {
  if (row) {
    editForm.value = { ...row }
  } else {
    editForm.value = { content: '', is_public: true }
  }
  isEditorOpen.value = true
}

function cancelEdit() {
  isEditorOpen.value = false
  editForm.value = { content: '', is_public: true }
}

async function save() {
  if (!formRef.value) return
  await formRef.value.validate(async (errors) => {
    if (!errors) {
      let res
      if (editForm.value.id) {
        res = await updateWhisperReq(editForm.value.id, editForm.value)
      } else {
        res = await createWhisperReq(editForm.value)
      }
      if (!res.code) {
        window.$message.success('Save success')
        cancelEdit()
      }
      loadData()
    }
  })
}

async function handleDelete(id: number) {
  const res = await deleteWhisperReq(id)
  if (!res.code) {
    window.$message.success('Delete success')
    if (whisperList.value.length === 1 && paginationParams.value.page > 1) {
      paginationParams.value.page--
    }
  }
  loadData()
}

async function changePublic(row: Whisper, val: boolean) {
  const res = await updateWhisperVisibilityReq(row.id, val)
  if (!res.code) {
    window.$message.success('Update success')
  }
  loadData()
}
</script>
