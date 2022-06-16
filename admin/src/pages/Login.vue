<script setup lang="ts">
import { useSesssionStore } from '../stores/session'
import { useRouter } from 'vue-router';

const sessionStore = useSesssionStore()
const router = useRouter()

const login = () => {
    sessionStore.saveUser({
        id: 1,
        username: 'admin',
    })
    router.replace('/')
}
</script>

<template>
<div>
  Login <button @click="login">Login</button>
</div>
</template>

<style>

  
</style>
