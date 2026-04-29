<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { loginReq } from '@/services'
import { NButton, NForm, NFormItem, NInput, type FormInst, type FormRules } from 'naive-ui'
import { useSesssionStore } from '@/stores/session'
import { useGlobalStore } from '@/stores/global'

interface ModelType {
  name: string
  password: string
}

const sessionStore = useSesssionStore()
const globalStore = useGlobalStore()
const formRef = ref<FormInst | null>(null)
const model = ref<ModelType>({ name: '', password: '' })
const rules: FormRules = {
  name: [{ required: true, message: 'Please enter your username.', trigger: ['input', 'blur'] }],
  password: [{ required: true, message: 'Please enter your password.', trigger: ['input', 'blur'] }],
}

const router = useRouter()

onMounted(() => {
  globalStore.getOptions()
})

const login = async () => {
  formRef.value?.validate(async (errors: any) => {
    if (!errors) {
      const { code } = await loginReq(model.value.name, model.value.password)
      if (!code) {
        await sessionStore.getUser()
        router.replace('/')
      }
    }
  })
}
</script>

<template>
  <div class="auth-shell">
    <section class="auth-panel">
      <div class="auth-panel__copy">
        <div>
          <span class="auth-panel__eyebrow">Simple Blog Admin</span>
          <h1 class="auth-panel__title">
            {{ globalStore.options.title || 'A cleaner desk for your publishing work.' }}
          </h1>
          <p class="auth-panel__body">
            Write, revise and keep the blog organized from one calm workspace. The layout stays
            quiet so the content can stay front and center.
          </p>
        </div>

        <div class="auth-panel__meta">
          <span class="auth-meta-pill">Focused writing</span>
          <span class="auth-meta-pill">Structured publishing</span>
          <span class="auth-meta-pill">Minimal overhead</span>
        </div>
      </div>

      <div class="auth-panel__form">
        <div>
          <h2 class="auth-form__heading">Sign in</h2>
          <p class="auth-form__copy">Use your admin account to continue into the workspace.</p>
        </div>

        <n-form ref="formRef" :rules="rules" :model="model" label-placement="top">
          <n-form-item label="Username" path="name">
            <n-input v-model:value="model.name" placeholder="Enter username" />
          </n-form-item>
          <n-form-item label="Password" path="password">
            <n-input
              v-model:value="model.password"
              type="password"
              placeholder="Enter password"
              @keydown.enter.prevent="login"
            />
          </n-form-item>
          <div class="form-actions">
            <n-button block type="primary" @click="login">Enter workspace</n-button>
          </div>
        </n-form>
      </div>
    </section>
  </div>
</template>
