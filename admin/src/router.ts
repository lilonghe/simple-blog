import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import PostList from './pages/PostList.vue'
import Login from './pages/Login.vue'

const routes = [
    { path: '/', component: Home },
    { path: '/post', component: PostList },
    { path: '/login', component: Login },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router