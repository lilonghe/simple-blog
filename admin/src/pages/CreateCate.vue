<script setup lang="ts">
import { useGlobalStore } from '@/stores/global'
import { NForm, NFormItem, NButton, NInput, NSpin } from 'naive-ui'
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getCateReq } from '@/services'

interface CateFormModel {
  id?: number
  name?: string
  pathname?: string
}

const form = ref<CateFormModel>({})
const formRef = ref()
const globalStore = useGlobalStore()
const router = useRouter()
const targetCate = ref<CateFormModel>()
const route = useRoute()

const handleSubmit = () => {
  formRef.value.validate(async (error: any) => {
    if (!error) {
      const success = await globalStore.createOrEditCate(form.value)
      if (success) {
        router.push('/category')
      }
    }
  })
}

const getTargetCate = async (id: number) => {
  const { data } = await getCateReq(id)
  if (data) {
    targetCate.value = data
    form.value = data
  }
}

onMounted(() => {
  if (route.params.id) {
    getTargetCate(route.params.id as any)
  }
})

const loaded = computed(() => !route.params.id || (route.params.id && targetCate.value?.id))
</script>

<template>
  <n-spin :show="!loaded">
    <div class="page-stack">
      <section class="page-surface page-surface--form">
        <div class="page-toolbar">
          <div>
            <p class="section-kicker">Taxonomy</p>
            <h2 class="section-title">
              {{ targetCate ? 'Update the category label and path' : 'Create a new category' }}
            </h2>
            <p class="section-copy">
              Categories should stay broad and stable so the archive remains easy to navigate over
              time.
            </p>
          </div>
        </div>

        <n-form ref="formRef" :model="form" label-placement="top">
          <div class="form-grid form-grid--single">
            <n-form-item v-if="targetCate" label="ID">
              <n-input :value="String(targetCate.id)" disabled />
            </n-form-item>
            <n-form-item
              label="Name"
              path="name"
              :rule="[{ required: true, trigger: ['blur', 'input'] }]"
            >
              <n-input v-model:value="form.name" placeholder="Category name" />
            </n-form-item>
            <n-form-item
              label="Path"
              path="pathname"
              :rule="[{ required: true, trigger: ['blur', 'input'] }]"
            >
              <n-input v-model:value="form.pathname" placeholder="category-path" />
            </n-form-item>
          </div>

          <div class="form-actions">
            <n-button type="primary" @click="handleSubmit">
              {{ targetCate ? 'Save category' : 'Create category' }}
            </n-button>
          </div>
        </n-form>
      </section>
    </div>
  </n-spin>
</template>
