import { createRouter, createWebHistory } from 'vue-router'

import CateList from './pages/CateList.vue'
import CreateCate from './pages/CreateCate.vue'
import CreatePage from './pages/CreatePage.vue'
import CreatePost from './pages/CreatePost.vue'
import CreateTag from './pages/CreateTag.vue'
import Home from './pages/Home.vue'
import Login from './pages/Login.vue'
import OptionsGeneral from './pages/management/General.vue'
import VisitList from './pages/management/VisitList.vue'
import PageList from './pages/PageList.vue'
import PostList from './pages/PostList.vue'
import TagList from './pages/TagList.vue'
import WhisperList from './pages/WhisperList.vue'
import { getRealPath } from './utils'

const routes = [
  {
    path: '/',
    component: Home,
    meta: {
      title: 'Overview',
      description: 'Track writing activity, reader attention and the overall health of your blog in one place.',
      section: 'Workspace',
    },
  },
  {
    path: '/post',
    component: PostList,
    meta: {
      title: 'Posts',
      description: 'Search, filter and manage long-form articles across draft and published states.',
      section: 'Publishing',
    },
  },
  {
    path: '/post/create',
    component: CreatePost,
    meta: {
      title: 'Create Post',
      description: 'Write and publish a new article with the metadata and presentation details it needs.',
      section: 'Publishing',
      back: '/post',
    },
  },
  {
    path: '/post/edit/:id',
    component: CreatePost,
    meta: {
      title: 'Edit Post',
      description: 'Revise the article body, update metadata and republish when everything is ready.',
      section: 'Publishing',
      back: '/post',
    },
  },
  {
    path: '/page',
    component: PageList,
    meta: {
      title: 'Pages',
      description: 'Maintain standalone pages such as About, Notes and other evergreen content.',
      section: 'Publishing',
    },
  },
  {
    path: '/page/create',
    component: CreatePage,
    meta: {
      title: 'Create Page',
      description: 'Compose a dedicated page and shape how it appears to readers before publishing.',
      section: 'Publishing',
      back: '/page',
    },
  },
  {
    path: '/page/edit/:id',
    component: CreatePage,
    meta: {
      title: 'Edit Page',
      description: 'Refine a standalone page and keep its presentation aligned with the rest of the site.',
      section: 'Publishing',
      back: '/page',
    },
  },
  { path: '/login', component: Login },
  {
    path: '/category',
    component: CateList,
    meta: {
      title: 'Categories',
      description: 'Curate your primary content buckets so the blog stays structured and easy to browse.',
      section: 'Taxonomy',
    },
  },
  {
    path: '/category/create',
    component: CreateCate,
    meta: {
      title: 'Create Category',
      description: 'Add a new category with a clear name and stable path for archive pages.',
      section: 'Taxonomy',
      back: '/category',
    },
  },
  {
    path: '/category/edit/:id',
    component: CreateCate,
    meta: {
      title: 'Edit Category',
      description: 'Update a category without breaking the information architecture readers rely on.',
      section: 'Taxonomy',
      back: '/category',
    },
  },
  {
    path: '/tag',
    component: TagList,
    meta: {
      title: 'Tags',
      description: 'Manage the lighter labeling system that helps readers discover related writing.',
      section: 'Taxonomy',
    },
  },
  {
    path: '/tag/create',
    component: CreateTag,
    meta: {
      title: 'Create Tag',
      description: 'Add a new tag with a concise label and a clean public path.',
      section: 'Taxonomy',
      back: '/tag',
    },
  },
  {
    path: '/tag/edit/:id',
    component: CreateTag,
    meta: {
      title: 'Edit Tag',
      description: 'Adjust tag naming and paths while keeping the content graph tidy.',
      section: 'Taxonomy',
      back: '/tag',
    },
  },
  {
    path: '/options/general',
    component: OptionsGeneral,
    meta: {
      title: 'General Settings',
      description: 'Configure the site identity, public links and other global details that shape the blog.',
      section: 'Settings',
    },
  },
  {
    path: '/options/visit/list',
    component: VisitList,
    meta: {
      title: 'Visit History',
      description: 'Review recent traffic details to understand what readers are opening and from where.',
      section: 'Settings',
    },
  },
  {
    path: '/whisper',
    component: WhisperList,
    meta: {
      title: 'Whispers',
      description: 'Capture short updates, toggle visibility and keep your micro-notes close at hand.',
      section: 'Publishing',
    },
  },
]

const router = createRouter({
    history: createWebHistory(getRealPath()),
    routes,
})

export default router
