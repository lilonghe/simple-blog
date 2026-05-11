<script setup lang="ts">
import {
  NGrid,
  NGi,
  NInput,
  NInputGroup,
  NInputGroupLabel,
  NButton,
  NCheckboxGroup,
  NCheckbox,
  NSelect,
  NSwitch,
  NFormItem,
  NForm,
  NSpin,
  NImage,
  NDatePicker,
} from 'naive-ui'
import { useGlobalStore } from '@/stores/global'
import { computed, onMounted, ref } from 'vue'
import { createPostReq, getEditPostReq, uploadFileReq } from '@/services'
import dayjs from 'dayjs'
import { useRoute } from 'vue-router'

const LOCAL_POST_SAVE_KEY = 'post-content'

interface PostFormModel {
  id?: number
  title?: string
  content?: string
  is_public?: boolean
  status?: number
  allow_comment?: boolean
  cate_list?: string[]
  tag_list?: string[]
  markdown_content?: string
  pathname?: string
  featuredImage?: string
  options?: string
  create_time: number
}

const globalStore = useGlobalStore()
const form = ref<PostFormModel>({
  is_public: true,
  allow_comment: true,
  status: 0,
  create_time: new Date().getTime(),
})
const targetPost = ref<PostFormModel>()
const formRef = ref()
const route = useRoute()
const editorRef = ref()

onMounted(() => {
  globalStore.getCates()
  globalStore.getTags()
  if (route.params.id) {
    getTargetPost(route.params.id as any)
  }
  if (!route.params.id && localStorage.getItem(LOCAL_POST_SAVE_KEY)) {
    form.value.markdown_content = localStorage.getItem(LOCAL_POST_SAVE_KEY) || ''
  }
})

const getTargetPost = async (id: number) => {
  const { data } = await getEditPostReq(id)
  if (data) {
    targetPost.value = data
    form.value = data
    form.value.create_time = new Date(data.create_time).getTime()
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
    create_time: dayjs(form.value.create_time).format(),
    update_time: dayjs().format(),
  }

  if (type === 'publish') {
    params.status = 3
  }

  const { code, data } = await createPostReq(params)
  if (!code) {
    const didPublish = type === 'publish' || params.status === 3
    window.$message.success(didPublish ? 'Publish success' : 'Save success')
    if (!didPublish) {
      localStorage.removeItem(LOCAL_POST_SAVE_KEY)
    }
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
  localStorage.setItem(LOCAL_POST_SAVE_KEY, markdown)
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
                <n-input v-model:value="form.title" placeholder="Write a sharp title" />
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
                    placeholder="post-path"
                  />
                  <n-input-group-label>.html</n-input-group-label>
                </n-input-group>
              </n-form-item>

              <div class="editor-main__fields-row">
                <span class="helper-text">
                  Stable paths keep public links clean and predictable.
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
              <h3 class="setting-group__title">Publishing</h3>
              <p class="setting-group__hint">
                Schedule the timestamp and decide how this entry appears to readers.
              </p>
              <n-form-item label="Publish time" :show-feedback="false">
                <n-date-picker v-model:value="form.create_time" type="datetime" />
              </n-form-item>
            </div>

            <div class="setting-group">
              <h3 class="setting-group__title">Taxonomy</h3>
              <p class="setting-group__hint">
                Categories are broad structure, tags are lightweight context.
              </p>
              <n-grid :cols="1" :x-gap="0">
                <n-gi>
                  <n-form-item label="Categories" :show-feedback="false">
                    <n-checkbox-group v-model:value="form.cate_list" class="flex flex-col gap-2">
                      <n-checkbox
                        v-for="item in globalStore.cates"
                        :key="item.id"
                        :value="item.id"
                      >
                        {{ item.name }}
                      </n-checkbox>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="Tags" :show-feedback="false">
                    <n-select
                      v-model:value="form.tag_list"
                      multiple
                      tag
                      filterable
                      :options="globalStore.tagOptions"
                    />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>

            <div class="setting-group">
              <h3 class="setting-group__title">Reader access</h3>
              <p class="setting-group__hint">
                Keep the content public, or share it privately with a direct link.
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
                Add a cover image if the post needs a stronger visual entry point.
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
