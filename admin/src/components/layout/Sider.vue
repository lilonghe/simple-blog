<script setup lang="ts">
import { NMenu, MenuOption } from 'naive-ui'
import { h } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: '/',
          activeClass: 'active'
        },
        { default: () => 'Dashboard' }
    ),
    key: '/',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: '/post'
        },
        { default: () => 'Post' }
    ),
    key: '/post',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: '/category'
        },
        { default: () => 'Category' }
    ),
    key: '/category',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: '/tag'
        },
        { default: () => 'Tag' }
    ),
    key: '/tag',
  },
  {
    label: "System",
    key: 'system-management',
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: '/options/general'
            },
            { default: () => 'Basic setting' }
        ),
        key: '/options/general',
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: '/options/visit/list'
            },
            { default: () => 'Visit List' }
        ),
        key: '/options/visit/list',
      },
    ]
  },
]

</script>

<template>
    <n-menu :options="menuOptions" :value="route.fullPath" />
</template>