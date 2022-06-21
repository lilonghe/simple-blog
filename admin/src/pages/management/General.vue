<script setup lang="ts">
import { useGlobalStore } from '@/stores/global';
import { NForm, NFormItem, NInput, NButton, NImage } from 'naive-ui';
import { ref, computed, h, onMounted } from 'vue';
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
        formModel.value = {...globalStore.options}
    })
})

const handleSubmit = () => {
    formRef.value.validate(async (error: any) => {
        if (!error) {
            let { code } = await updateOptionsReq(formModel.value)
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
    }
}

</script>
<template>
<n-form 
    class="w-1/3"
    :model="formModel" 
    :rules="rules"
    ref="formRef">
    <n-form-item label="Title" path="title">
        <n-input v-model:value="formModel.title" placeholder="Title" />
    </n-form-item>
    <n-form-item label="Logo" path="logo_url">
        <n-input v-model:value="formModel.logo_url" placeholder="Logo url" />
        <n-image :src="formModel.logo_url" />
    </n-form-item>
    <n-form-item label="Description" path="description">
        <n-input v-model:value="formModel.description" placeholder="Description" />
    </n-form-item>
    <n-form-item label="Site URL" path="site_url">
        <n-input v-model:value="formModel.site_url" placeholder="Site URL" />
    </n-form-item>
    <n-form-item label="Favicon URL" path="favicon_url">
        <n-input v-model:value="formModel.favicon_url" placeholder="Favicon URL" />
        <n-image :src="formModel.favicon_url" />
    </n-form-item>
    <n-form-item label="Keywords" path="keywords">
        <n-input v-model:value="formModel.keywords" placeholder="Keywords" />
    </n-form-item>
    <n-form-item label="Github address" path="github_url">
        <n-input v-model:value="formModel.github_url" placeholder="Github address" />
    </n-form-item>
    <n-form-item label="Twitter address" path="twitter_url">
        <n-input v-model:value="formModel.twitter_url" placeholder="Twitter address" />
    </n-form-item>

    <n-button v-on:click="handleSubmit" type="primary">Submit</n-button>
</n-form>
</template>