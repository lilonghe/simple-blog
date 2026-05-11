<script setup lang="ts">
import { createPostReq, getEditPostReq, uploadFileReq } from '@/services'
import dayjs from 'dayjs'
import {
  NButton,
  NForm,
  NFormItem,
  NImage,
  NInput,
  NInputGroup,
  NInputGroupLabel,
  NSpin,
  NSwitch,
} from 'naive-ui'
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useGlobalStore } from '@/stores/global'

interface PostFormModel {
  id?: number
  title?: string
  content?: string
  is_public?: boolean
  status?: number
  allow_comment?: boolean
  markdown_content?: string
  pathname?: string
  featuredImage?: string
  options?: string
  create_time?: string
}

const globalStore = useGlobalStore()
const form = ref<PostFormModel>({
  is_public: true,
  allow_comment: true,
  status: 0,
})
const targetPost = ref<PostFormModel>()
const formRef = ref()
const route = useRoute()
const editorRef = ref()

onMounted(() => {
  globalStore.getOptions()
  if (route.params.id) {
    getTargetPost(route.params.id as any)
  }
})

const getTargetPost = async (id: number) => {
  const { data } = await getEditPostReq(id)
  if (data) {
    targetPost.value = data
    form.value = data
    if (data.options) {
      const { featuredImage } = JSON.parse(data.options)
      form.value.featuredImage = featuredImage
    }
  }
}

const preProcessFormData = () => {
  if (form.value.featuredImage) {
    form.value.options = JSON.stringify({
      featuredImage: form.value.featuredImage,
    })
  } else {
    form.value.options = undefined
  }
}

const submit = async (type?: 'publish') => {
  preProcessFormData()
  const params = {
    ...form.value,
    create_time: form.value.create_time || dayjs().format(),
    update_time: dayjs().format(),
    type: 1,
  }
  if (type === 'publish') {
    params.status = 3
  }

  const { code, data } = await createPostReq(params)
  if (!code) {
    const didPublish = type === 'publish' || params.status === 3
    window.$message.success(didPublish ? 'Publish success' : 'Save success')
    getTargetPost(data.id)
  }
}

const handleSave = async () => {
  formRef.value.validate(async (error: any) => {
    if (!error) {
      submit()
    }
  })
}

const handlePublish = async () => {
  formRef.value.validate(async (error: any) => {
    if (!error) {
      submit('publish')
    }
  })
}

const handleEditChange = (markdown: string, html: string) => {
  form.value.markdown_content = markdown
  form.value.content = html
}

const handleUploadImage = (pos: string, file: any) => {
  const formData = new FormData()
  formData.append('file', file)
  uploadFileReq(formData).then((res: any) => {
    if (!res.code) {
      editorRef.value.$img2Url(pos, res.data)
    } else {
      editorRef.value.$refs.toolbar_left.$imgDel(pos)
    }
  })
}

const loaded = computed(() => !route.params.id || targetPost.value?.id !== undefined)
</script>

<template>
  <n-spin :show="!loaded">
    <div class="page-stack">
      <n-form ref="formRef" :model="form" label-placement="top">
        <div class="editor-layout">
          <section class="page-surface editor-main">
            <div class="editor-main__fields">
              <n-form-item
                path="title"
                :show-feedback="false"
                :rule="[{ required: true, trigger: ['blur', 'input'] }]"
              >
                <n-input v-model:value="form.title" placeholder="Name the page" />
              </n-form-item>

              <n-form-item
                path="pathname"
                :show-feedback="false"
                :rule="[{ required: true, trigger: ['blur', 'input'] }]"
              >
                <n-input-group>
                  <n-input-group-label>{{ globalStore.options.site_url }}/post/</n-input-group-label>
                  <n-input
                    v-model:value="form.pathname"
                    :disabled="targetPost?.id != undefined"
                    placeholder="page-path"
                  />
                  <n-input-group-label>.html</n-input-group-label>
                </n-input-group>
              </n-form-item>

              <div class="editor-main__fields-row">
                <span class="helper-text">
                  Use pages for evergreen content such as About, Colophon or resource hubs.
                </span>
                <a
                  v-if="targetPost?.pathname"
                  class="link"
                  target="_blank"
                  :href="
                    globalStore.options.site_url +
                    `/post/${targetPost.pathname}.html${targetPost.status === 0 ? '?preview=true' : ''}`
                  "
                >
                  View current version
                </a>
              </div>
            </div>

            <div v-if="!route.params.id || targetPost?.id" class="editor-canvas">
              <mavon-editor
                ref="editorRef"
                v-model="form.markdown_content"
                @imgAdd="handleUploadImage"
                @change="handleEditChange"
                :box-shadow="false"
                :ishljs="false"
              />
            </div>
          </section>

          <aside class="page-surface editor-sidebar">
            <div class="editor-sidebar__actions">
              <n-button v-if="targetPost?.status != 3" @click="handleSave">Save draft</n-button>
              <n-button type="primary" @click="handlePublish">Publish</n-button>
            </div>

            <div class="setting-group">
              <h3 class="setting-group__title">Reader access</h3>
              <p class="setting-group__hint">
                Standalone pages often stay public, but you can still keep them private when needed.
              </p>
              <div class="flex flex-col gap-4">
                <div class="switch-row">
                  <span>Public</span>
                  <n-switch v-model:value="form.is_public" />
                </div>
                <div class="switch-row">
                  <span>Allow comments</span>
                  <n-switch v-model:value="form.allow_comment" />
                </div>
              </div>
            </div>

            <div class="setting-group">
              <h3 class="setting-group__title">Feature image</h3>
              <p class="setting-group__hint">
                Optional, but useful if the page should carry a stronger editorial feel.
              </p>
              <n-form-item :show-feedback="false">
                <n-input v-model:value="form.featuredImage" placeholder="Image URL" />
              </n-form-item>
              <div v-if="form.featuredImage" class="preview-image">
                <n-image :src="form.featuredImage" object-fit="cover" />
              </div>
            </div>
          </aside>
        </div>
      </n-form>
    </div>
  </n-spin>
</template>
