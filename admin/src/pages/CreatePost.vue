<script setup lang="ts">
import { NGrid, NGi, NInput, NInputGroup, NInputGroupLabel, NButton, NCheckboxGroup, NCheckbox, NSelect, NSwitch, NFormItem, NForm, NSpin, NImage } from 'naive-ui'
import { useGlobalStore } from '@/stores/global'
import { computed, onMounted, ref } from 'vue'
import { createPostReq, getEditPostReq } from '@/services'
import dayjs from 'dayjs';
import { useRoute, useRouter } from 'vue-router';

interface PostFormModel {
  id?: number;
  title?: string
  content?: string
  is_public?: boolean
  status?: number
  tags?: string[]
  allow_comment?: boolean
  cate_list?: string[]
  tag_list?: string[]
  markdown_content?: string
  pathname?: string
  featuredImage?: string
  options?: string
  create_time?: string
}

const globalStore = useGlobalStore();
const form = ref<PostFormModel>({
    is_public: true,
    allow_comment: true,
    status: 0,
})
const targetPost = ref<PostFormModel>()
const formRef = ref()
const router = useRouter()
const route = useRoute()

onMounted(() => {
    globalStore.getCates()
    globalStore.getTags()
    if (route.params.id) {
        getTargetPost(route.params.id as any)
    }
})

const getTargetPost = async (id: number) => {
    let { data } = await getEditPostReq(id)
    if (data) {
        targetPost.value = data
        form.value = data
        if (data.options) {
            let { featuredImage } = JSON.parse(data.options)
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

const submit = async (type?: string) => {
    preProcessFormData()
    let params = {
        ...form.value,
        create_time: form.value.create_time || dayjs().format(),
        update_time: dayjs().format(),
    }
    if (type === 'publish') {
        params.status = 3
    }

    let { code, data } = await createPostReq(params)
    if (!code) {
        if (form.value.status === 0) {
            window.$message.success('Save success')
        } else {
            window.$message.success('Publish success')
        }
        getEditPostReq(data.id)
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

const loaded = computed(() => !route.params.id || (route.params.id && targetPost.value?.id !== undefined))
</script>
<template>
<n-spin :show="!loaded">
<n-form ref="formRef" :model="form">
    <n-grid x-gap="20">
        <n-gi :span="18" class="flex flex-col gap-6">
            <n-form-item
                :show-label="false" 
                :show-feedback="false"
                path="title"
                :rule="[{required: true, trigger: ['blur', 'input']}]">
                <n-input placeholder="Title" v-model:value="form.title" />
            </n-form-item>
            <n-form-item
                :show-label="false" 
                :show-feedback="false"
                path="pathname"
                :rule="[{required: true, trigger: ['blur', 'input']}]">
                <n-input-group>
                    <n-input-group-label>{{globalStore.options.site_url}}/post/</n-input-group-label>
                    <n-input
                        :disabled="targetPost?.id != undefined"
                        placeholder="Path name" 
                        v-model:value="form.pathname" />
                    <n-input-group-label>.html</n-input-group-label>
                </n-input-group>
                <a class="link ml-2" target="_blank" :href="globalStore.options.site_url + `/post/${targetPost?.pathname}.html`">View</a>
            </n-form-item>
            <div v-if="!route.params.id || (route.params.id && targetPost?.id)">
                <mavon-editor 
                    v-model="form.markdown_content"
                    :onChange="handleEditChange"
                    :boxShadow="false" class="h-screen"
                    :ishljs="false" />
            </div>
        </n-gi>
        <n-gi :span="6">
            <div class="flex gap-1">
                <n-button v-on:click="handleSave" v-if="!targetPost?.id">Save</n-button>
                <n-button type="primary" v-on:click="handlePublish">Publish</n-button>
            </div>
            <div>
                <div class="text-xl mt-6 mb-1">Catgory</div>
                <n-form-item :show-label="false" :show-feedback="false">
                    <n-checkbox-group class="flex flex-col" v-model:value="form.cate_list">
                        <n-checkbox v-for="item in globalStore.cates" :value="item.id">{{item.name}}</n-checkbox>
                    </n-checkbox-group>
                </n-form-item>
            </div>
            <div>
                <div class="text-xl mt-4">Tag</div>
                <n-form-item :show-label="false" :show-feedback="false">
                    <n-select multiple tag filterable :options="globalStore.tagOptions" v-model:value="form.tag_list" />
                </n-form-item>
            </div>
            <div>
                <div class="text-xl mt-4">Public</div>
                <n-form-item :show-label="false" :show-feedback="false">
                    <n-switch v-model:value="form.is_public"/>
                </n-form-item>
            </div>
            <div>
                <div class="text-xl mt-4">Allow comment</div>
                <n-form-item :show-label="false" :show-feedback="false">
                    <n-switch v-model:value="form.allow_comment" />
                </n-form-item>
            </div>
            <div>
                <div class="text-xl mt-4">Featured Image</div>
                <n-input placeholder="img url" v-model:value="form.featuredImage" />
                <n-image :src="form.featuredImage" />
            </div>
        </n-gi>
    </n-grid>
</n-form>
</n-spin>
</template>