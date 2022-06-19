<script setup lang="ts">
import { useGlobalStore } from '@/stores/global';
import { NForm, NFormItem, NButton, NInput, NSpin } from 'naive-ui';
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getCateReq } from '@/services'

interface CateFormModel {
    id?: number;
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
            let success = await globalStore.createOrEditCate(form.value)
            if (success) {
                router.push('/cate/list')
            }
        }
    })
}

const getTargetCate = async (id: number) => {
    let { data } = await getCateReq(id)
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

const loaded = computed(() => !route.params.id || route.params.id && targetCate.value?.id)
</script>
<template>
<n-spin :show="!loaded">
    <n-form ref="formRef" class="w-1/3" :model="form">
        <n-form-item label="Id" v-if="targetCate">
            <n-input :value="(targetCate.id as any)" disabled />
        </n-form-item>
        <n-form-item label="Name" path="name" :rule="[{required: true, trigger: ['blur', 'input']}]">
            <n-input v-model:value="form.name" placeholder="Cate name" />
        </n-form-item>
        <n-form-item label="Path" path="pathname" :rule="[{required: true, trigger: ['blur', 'input']}]">
            <n-input v-model:value="form.pathname" placeholder="Cate path" />
        </n-form-item>
        <n-button type="primary" v-on:click="handleSubmit">Submit</n-button>
    </n-form>
</n-spin>
</template>