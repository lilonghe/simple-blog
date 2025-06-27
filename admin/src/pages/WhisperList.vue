<template>
  <div>
    <div class="mb-2 flex gap-4">
      <n-input placeholder="搜索内容" v-model:value="filterParams.keyword" style="width: 200px" @change="loadData" />
      <n-select
        placeholder="公开性"
        clearable
        v-model:value="filterParams.is_public"
        :options="[
          { label: '公开', value: 'true' },
          { label: '私密', value: 'false' }
        ]"
        style="width: 120px"
        @update:value="loadData"
      />
      <n-button type="primary" @click="openEdit()">Create</n-button>
    </div>
    <n-data-table
      :loading="loading"
      :remote="true"
      :data="whisperList"
      :columns="columns"
      :pagination="pagination"
      @update:page="handlePageChange"
    />
    <n-modal v-model:show="editVisible" preset="dialog" :show-icon="false" :title="editForm.id ? 'Edit' : 'Create'">
      <n-form ref="formRef" :model="editForm" :rules="rules" label-placement="top" :show-label="false">
        <n-form-item label="内容" path="content">
          <n-input v-model:value="editForm.content" type="textarea" rows="4" placeholder="内容"/>
        </n-form-item>
        <n-form-item label="公开性" path="is_public" :show-feedback="false">
          <n-switch v-model:value="editForm.is_public">
            <template #checked>公开</template>
            <template #unchecked>私密</template>
          </n-switch>
        </n-form-item>
      </n-form>
      <template #action>
        <n-button @click="editVisible=false">取消</n-button>
        <n-button type="primary" @click="save" style="margin-left: 8px;">保存</n-button>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import { NDataTable, NInput, NPopconfirm, NButton, NSwitch, NSelect, NModal, NForm, NFormItem } from 'naive-ui'
import { getWhisperListReq, createWhisperReq, updateWhisperReq, deleteWhisperReq, updateWhisperVisibilityReq } from '@/services'
import { formatTime } from '@/utils'
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

const editVisible = ref(false)
const editForm = ref<Partial<Whisper>>({ id: 0, content: '', is_public: true })

const pagination = computed(() => ({
  page: paginationParams.value.page,
  pageSize: paginationParams.value.pageSize,
  itemCount: total.value,
}))

const formRef = ref<FormInst | null>(null)

const rules: FormRules = {
  content: [
    { required: true, message: '内容不能为空', trigger: ['input', 'blur'] }
  ]
}

const loadData = async () => {
  loading.value = true
  const params: any = {
    page: paginationParams.value.page,
    limit: paginationParams.value.pageSize,
    ...filterParams.value
  }
  const { data } = await getWhisperListReq(params)
  if (data) {
    whisperList.value = data.list
    total.value = data.total
  }
  loading.value = false
}

onMounted(loadData)

const columns = [
  {
    title: '内容',
    key: 'content',
    ellipsis: { tooltip: true },
  },
  {
    title: '公开',
    key: 'is_public',
    width: 80,
    render: (row: Whisper) =>
      h(NSwitch, {
        value: row.is_public,
        onUpdateValue: (val: boolean) => changePublic(row, val),
      })
  },
  {
    title: '创建时间',
    key: 'create_time',
    render: (row: Whisper) => formatTime(row.create_time),
    width: 200,
  },
  {
    title: '操作',
    key: 'action',
    width: 180,
    render: (row: Whisper) => h('div', { class: 'flex gap-4' }, [
      h(NButton, {
        size: 'small',
        onClick: () => openEdit(row)
      }, { default: () => '编辑' }),
      h(NPopconfirm, {
        onPositiveClick: () => handleDelete(row.id),
      }, {
        trigger: () => h(NButton, { size: 'small', type: 'error' }, { default: () => '删除' }),
        default: () => `确定删除？`,
      })
    ])
  }
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
  editVisible.value = true
}

async function save() {
  if (!formRef.value) return
  await formRef.value.validate(async (errors) => {
    if (!errors) {
      let res;
      if (editForm.value.id) {
        res = await updateWhisperReq(editForm.value.id, editForm.value)
      } else {
        res = await createWhisperReq(editForm.value)
      }
      if (!res.code) {
        window.$message.success('保存成功')
        editVisible.value = false
      }
      loadData()
    }
  })
}

async function handleDelete(id: number) {
  const res = await deleteWhisperReq(id)
  if (!res.code) {
    window.$message.success('删除成功')
    if (whisperList.value.length === 1 && paginationParams.value.page > 1) {
      paginationParams.value.page--
    }
  }
  loadData()
}

async function changePublic(row: Whisper, val: boolean) {
  const res = await updateWhisperVisibilityReq(row.id, val)
  if (!res.code) {
    window.$message.success('修改成功')
  }
  loadData()
}
</script> 