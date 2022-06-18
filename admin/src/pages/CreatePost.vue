<script setup lang="ts">
import { NGrid, NGi, NInput, NInputGroup, NInputGroupLabel, NButton, NCheckboxGroup, NCheckbox, NSelect, NSwitch, NFormItem, NForm } from 'naive-ui'
import { useGlobalStore } from '@/stores/global'
import { onMounted, ref } from 'vue'
import { createPostReq } from '@/services'
import dayjs from 'dayjs';
import { useRouter } from 'vue-router';

interface PostFormModel {
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
}

const globalStore = useGlobalStore();
const form = ref<PostFormModel>({
    is_public: true,
    allow_comment: true,
    status: 0,
})
const formRef = ref()
const router = useRouter()

onMounted(() => {
    globalStore.getCates()
    globalStore.getTags()
})

const preProcessFormData = () => {
    if (form.value.featuredImage) {
        form.value.options = JSON.stringify({
            featuredImage: form.value.featuredImage,
        })
    } else {
        form.value.options = undefined
    }
}

const submit = async () => {
    preProcessFormData()
    let params = {
        ...form.value,
        create_time: dayjs().format(),
        update_time: dayjs().format(),
    }
    let { code } = await createPostReq(params)
    if (!code) {
        if (form.value.status === 0) {
            $message.success('Save success')
            router.push('/post/list')
        } else {
            $message.success('Create post success')
            router.push('/post/list')
        }
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
            form.value.status = 3
            submit()
        }
    })
}

const handleEditChange = (markdown: string, html: string) => {
    form.value.markdown_content = markdown
    form.value.content = html
}

</script>
<template>
<n-form ref="formRef" :model="form">
    <n-grid x-gap="10">
        <n-gi :span="18" class="flex flex-col gap-6">
            <n-form-item 
                        :show-label="false" 
                        :show-feedback="false"
                        path="title"
                        :rule="[{required: true, trigger: ['blur', 'input']}]">
                <n-input placeholder="Title" v-model:value="form.title" />
            </n-form-item>
            <div class="flex items-center gap-1">
                <n-form-item 
                        :show-label="false" 
                        :show-feedback="false"
                        path="pathname"
                        :rule="[{required: true, trigger: ['blur', 'input']}]">
                    <n-input-group>
                        <n-input-group-label>{{globalStore.options.site_url}}/post/</n-input-group-label>
                        <n-input
                            :style="{ width: '33%' }" 
                            placeholder="Path name" 
                            v-model:value="form.pathname" />
                        <n-input-group-label>.html</n-input-group-label>
                    </n-input-group>
                </n-form-item>
            </div>
            <div>
                <mavon-editor 
                    :value="form.markdown_content"
                    :onChange="handleEditChange"
                    :boxShadow="false" class="h-screen"
                    :ishljs="false" />
            </div>
        </n-gi>
        <n-gi :span="6">
            <div class="flex gap-1">
                <n-button v-on:click="handleSave">Save</n-button>
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
            </div>
        </n-gi>
    </n-grid>
</n-form>
</template>