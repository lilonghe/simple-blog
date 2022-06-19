<script setup lang="ts">
import { useGlobalStore } from '@/stores/global';
import { NForm, NFormItem, NButton, NInput, NSpin } from 'naive-ui';
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getTagReq } from '@/services'

interface TagFormModel {
    id?: number;
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
            let success = await globalStore.createOrEditTag(form.value)
            if (success) {
                router.push('/tag/list')
            }
        }
    })
}

const getTargetTag = async (id: number) => {
    let { data } = await getTagReq(id)
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

const loaded = computed(() => !route.params.id || route.params.id && targetTag.value?.id)
</script>
<template>
<n-spin :show="!loaded">
    <n-form ref="formRef" class="w-1/3" :model="form">
        <n-form-item label="Id" v-if="targetTag">
            <n-input :value="(targetTag.id as any)" disabled />
        </n-form-item>
        <n-form-item label="Name" path="name" :rule="[{required: true, trigger: ['blur', 'input']}]">
            <n-input v-model:value="form.name" placeholder="Tag name" />
        </n-form-item>
        <n-form-item label="Path" path="pathname" :rule="[{required: true, trigger: ['blur', 'input']}]">
            <n-input v-model:value="form.pathname" placeholder="Tag path" />
        </n-form-item>
        <n-button type="primary" v-on:click="handleSubmit">Submit</n-button>
    </n-form>
</n-spin>
</template>