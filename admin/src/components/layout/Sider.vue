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
    label: "Post management",
    key: 'post-management',
    children: [
      {
        label: () =>
          h(
            RouterLink,
            {
              to: '/post/list'
            },
            { default: () => 'Post list' }
        ),
        key: '/post/list',
      },
      {
        label: () =>
          h(
            RouterLink,
            {
              to: '/post/create'
            },
            { default: () => 'Post create' }
        ),
        key: '/post/create',
      }
    ]
  },
]

</script>

<template>
    <n-menu :options="menuOptions" :value="route.fullPath" />
</template>