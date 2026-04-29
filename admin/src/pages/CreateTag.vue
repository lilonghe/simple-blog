<script setup lang="ts">
import { useGlobalStore } from '@/stores/global'
import { NForm, NFormItem, NButton, NInput, NSpin } from 'naive-ui'
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getTagReq } from '@/services'

interface TagFormModel {
  id?: number
  name?: string
  pathname?: string
}

const form = ref<TagFormModel>({})
const formRef = ref()
const globalStore = useGlobalStore()
const router = useRouter()
const targetTag = ref<TagFormModel>()
const route = useRoute()

const handleSubmit = () => {
  formRef.value.validate(async (error: any) => {
    if (!error) {
      const success = await globalStore.createOrEditTag(form.value)
      if (success) {
        router.push('/tag')
      }
    }
  })
}

const getTargetTag = async (id: number) => {
  const { data } = await getTagReq(id)
  if (data) {
    targetTag.value = data
    form.value = data
  }
}

onMounted(() => {
  if (route.params.id) {
    getTargetTag(route.params.id as any)
  }
})

const loaded = computed(() => !route.params.id || (route.params.id && targetTag.value?.id))
</script>

<template>
  <n-spin :show="!loaded">
    <div class="page-stack">
      <section class="page-surface page-surface--form">
        <div class="page-toolbar">
          <div>
            <p class="section-kicker">Taxonomy</p>
            <h2 class="section-title">
              {{ targetTag ? 'Update the tag label and path' : 'Create a new tag' }}
            </h2>
            <p class="section-copy">
              Tags work best when they stay short, specific and consistently reusable.
            </p>
          </div>
        </div>

        <n-form ref="formRef" :model="form" label-placement="top">
          <div class="form-grid form-grid--single">
            <n-form-item v-if="targetTag" label="ID">
              <n-input :value="String(targetTag.id)" disabled />
            </n-form-item>
            <n-form-item
              label="Name"
              path="name"
              :rule="[{ required: true, trigger: ['blur', 'input'] }]"
            >
              <n-input v-model:value="form.name" placeholder="Tag name" />
            </n-form-item>
            <n-form-item
              label="Path"
              path="pathname"
              :rule="[{ required: true, trigger: ['blur', 'input'] }]"
            >
              <n-input v-model:value="form.pathname" placeholder="tag-path" />
            </n-form-item>
          </div>

          <div class="form-actions">
            <n-button type="primary" @click="handleSubmit">
              {{ targetTag ? 'Save tag' : 'Create tag' }}
            </n-button>
          </div>
        </n-form>
      </section>
    </div>
  </n-spin>
</template>
