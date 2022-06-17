<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router'
import { loginReq } from '@/services'
import { NButton, NForm, NFormItem, NInput, FormInst, FormRules } from 'naive-ui';

interface ModelType {
    name: string
    password: string
}

const formRef = ref<FormInst | null>();
const model = ref<ModelType>({ name: '', password: '' });
const rules: FormRules = {
    name: [{ required: true, message: 'please input username', trigger: ['input', 'blur'] }],
    password: [{required: true, message: 'please input password', trigger: ['input', 'blur']}]
}

const router = useRouter()

const login = (e: any) => {
    e.preventDefault()
    formRef?.value?.validate(async (errors: any) => {
        if (!errors) {
            let { code } = await loginReq(model.value.name, model.value.password)
            if (!code) {
                router.replace('/')
            }
        }
    })
    
}
</script>

<template>
<div class="form">
    <n-form
    ref="formRef"
    :label-width="80"
    :rules="rules"
    :model="model"
  >
    <n-form-item label="Name" path="name">
      <n-input v-model:value="model.name" placeholder="Username" />
    </n-form-item>
    <n-form-item label="Password" path="password">
      <n-input type="password" v-model:value="model.password" placeholder="Password" />
    </n-form-item>
    <n-form-item class="flex justify-end">
      <n-button attr-type="button" @click="login">Login</n-button>
    </n-form-item>
  </n-form>
</div>
</template>

<style>
.form {
    position: fixed;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 300px;
}
  
</style>
