import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import PostList from './pages/PostList.vue'
import Login from './pages/Login.vue'
import CreatePost from './pages/CreatePost.vue'
import CateList from './pages/CateList.vue'
import CreateCate from './pages/CreateCate.vue'
import TagList from './pages/TagList.vue'
import CreateTag from './pages/CreateTag.vue'
import OptionsGeneral from './pages/management/General.vue'
import VisitList from './pages/management/VisitList.vue'

const routes = [
    { path: '/', component: Home },
    { path: '/post/list', component: PostList },
    { path: '/post/create', component: CreatePost },
    { path: '/post/edit/:id', component: CreatePost },
    { path: '/login', component: Login },
    { path: '/cate/list', component: CateList },
    { path: '/cate/create', component: CreateCate },
    { path: '/cate/edit/:id', component: CreateCate },
    { path: '/tag/list', component: TagList },
    { path: '/tag/create', component: CreateTag },
    { path: '/tag/edit/:id', component: CreateTag },

    { path: '/options/general', component: OptionsGeneral },
    { path: '/options/visit/list', component: VisitList },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router