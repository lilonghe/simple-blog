<script setup lang="ts">
import { useGlobalStore } from '@/stores/global'
import { NForm, NFormItem, NInput, NButton, NImage } from 'naive-ui'
import { ref, onMounted } from 'vue'
import { updateOptionsReq } from '@/services'

interface GeneralFormModel {
  title?: string
  logo_url?: string
  description?: string
  keywords?: string
  github_url?: string
  twitter_url?: string
  site_url?: string
  favicon_url?: string
}

const formModel = ref<GeneralFormModel>({})
const formRef = ref()
const globalStore = useGlobalStore()

onMounted(() => {
  globalStore.getOptions().then(() => {
    formModel.value = { ...globalStore.options }
  })
})

const handleSubmit = () => {
  formRef.value.validate(async (error: any) => {
    if (!error) {
      const { code } = await updateOptionsReq(formModel.value)
      if (!code) {
        globalStore.getOptions()
        window.$message.success('Update options success')
      }
    }
  })
}

const rules = {
  title: {
    required: true,
  },
  description: {
    required: true,
  },
}
</script>

<template>
  <div class="page-stack">
    <section class="page-surface page-surface--form">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Identity</p>
          <h2 class="section-title">Site presentation and public links</h2>
          <p class="section-copy">
            These fields define how the blog is named, described and linked across the public
            surface.
          </p>
        </div>
      </div>

      <n-form ref="formRef" :model="formModel" :rules="rules" label-placement="top">
        <div class="form-grid">
          <n-form-item label="Title" path="title">
            <n-input v-model:value="formModel.title" placeholder="Blog title" />
          </n-form-item>

          <n-form-item label="Site URL" path="site_url">
            <n-input v-model:value="formModel.site_url" placeholder="https://example.com" />
          </n-form-item>

          <n-form-item label="Description" path="description">
            <n-input v-model:value="formModel.description" placeholder="Short site description" />
          </n-form-item>

          <n-form-item label="Keywords" path="keywords">
            <n-input v-model:value="formModel.keywords" placeholder="keyword, keyword, keyword" />
          </n-form-item>

          <n-form-item label="Logo URL" path="logo_url">
            <div class="w-full">
              <n-input v-model:value="formModel.logo_url" placeholder="Logo image URL" />
              <div v-if="formModel.logo_url" class="preview-image">
                <n-image :src="formModel.logo_url" object-fit="contain" />
              </div>
            </div>
          </n-form-item>

          <n-form-item label="Favicon URL" path="favicon_url">
            <div class="w-full">
              <n-input v-model:value="formModel.favicon_url" placeholder="Favicon image URL" />
              <div v-if="formModel.favicon_url" class="preview-image">
                <n-image :src="formModel.favicon_url" object-fit="contain" />
              </div>
            </div>
          </n-form-item>

          <n-form-item label="GitHub URL" path="github_url">
            <n-input v-model:value="formModel.github_url" placeholder="https://github.com/..." />
          </n-form-item>

          <n-form-item label="Twitter URL" path="twitter_url">
            <n-input v-model:value="formModel.twitter_url" placeholder="https://x.com/..." />
          </n-form-item>
        </div>

        <div class="form-actions">
          <n-button type="primary" @click="handleSubmit">Save settings</n-button>
        </div>
      </n-form>
    </section>
  </div>
</template>
