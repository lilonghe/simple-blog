<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { NMenu } from 'naive-ui'
import { computed, h } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

interface NavItem {
  label: string
  key: string
}

interface NavGroup {
  type: 'group'
  label: string
  key: string
  children: NavItem[]
}

const renderLink = (label: string, to: string) => () =>
  h(
    RouterLink,
    {
      to,
      class: 'admin-nav__link',
    },
    { default: () => label },
  )

const navigation: Array<NavItem | NavGroup> = [
  {
    label: 'Overview',
    key: '/',
  },
  {
    type: 'group',
    label: 'Publishing',
    key: 'publishing',
    children: [
      {
        label: 'Posts',
        key: '/post',
      },
      {
        label: 'Pages',
        key: '/page',
      },
      {
        label: 'Whispers',
        key: '/whisper',
      },
    ],
  },
  {
    type: 'group',
    label: 'Taxonomy',
    key: 'taxonomy',
    children: [
      {
        label: 'Categories',
        key: '/category',
      },
      {
        label: 'Tags',
        key: '/tag',
      },
    ],
  },
  {
    type: 'group',
    label: 'Settings',
    key: 'settings',
    children: [
      {
        label: 'General',
        key: '/options/general',
      },
      {
        label: 'Visit history',
        key: '/options/visit/list',
      },
    ],
  },
]

const menuOptions: MenuOption[] = navigation.map((item) => {
  if ('children' in item) {
    return {
      ...item,
      children: item.children.map((child) => ({
        ...child,
        label: renderLink(child.label, child.key),
      })),
    }
  }

  return {
    ...item,
    label: renderLink(item.label, item.key),
  }
})

const routeKeys = navigation.flatMap((item) => ('children' in item ? item.children : item)).map((item) => item.key)

const selectedKey = computed(() => {
  return routeKeys.findLast((key) => route.path === key || route.path.startsWith(`${key}/`)) ?? '/'
})
</script>

<template>
<n-menu class="admin-nav" :options="menuOptions" :value="selectedKey" />
</template>
