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
import { getRealPath } from './utils'

const routes = [
    { path: '/', component: Home },
    { path: '/post', component: PostList, meta: { title: 'Post Manage'} },
    { path: '/post/create', component: CreatePost, meta: { title: 'Create Post', back: '/post' }  },
    { path: '/post/edit/:id', component: CreatePost, meta: { title: 'Edit Post', back: '/post' }  },
    { path: '/login', component: Login },
    { path: '/category', component: CateList, meta: { title: 'Category Manage'}  },
    { path: '/category/create', component: CreateCate, meta: { title: 'Create Category', back: '/category' }  },
    { path: '/category/edit/:id', component: CreateCate, meta: { title: 'Edit Category', back: '/category' }  },
    { path: '/tag', component: TagList, meta: { title: 'Tag Manage' }  },
    { path: '/tag/create', component: CreateTag, meta: { title: 'Create Tag', back: '/tag' }  },
    { path: '/tag/edit/:id', component: CreateTag, meta: { title: 'Edit Tag', back: '/tag' } },

    { path: '/options/general', component: OptionsGeneral, meta: { title: 'Basic Setting'}  },
    { path: '/options/visit/list', component: VisitList, meta: { title: 'Visit History'}  },
]

const router = createRouter({
    history: createWebHistory(getRealPath()),
    routes,
})

export default router