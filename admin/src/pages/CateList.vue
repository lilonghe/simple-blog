<script setup lang="ts">
import { NDataTable, NPopconfirm } from 'naive-ui';
import { h, onMounted } from 'vue';
import { useGlobalStore } from '@/stores/global';
import { RouterLink } from 'vue-router';

const globalStore = useGlobalStore();

onMounted(() => {
    globalStore.getCates()
})

const handleDelete = (id: number) => {
    globalStore.deleteCate(id)
}

const columns = [
    {
        title: 'Name',
        key: 'name'
    },
    {
        title: 'Path',
        key: 'pathname'
    },
    {
        title: 'Action',
        key: 'action',
        render: (row: any) => {
            return h(
                'div',
                {
                    class: 'flex gap-4',
                },
                [
                    h(
                        RouterLink,
                        {
                            to: '/cate/edit/' + row.id,
                            class: 'link'
                        },
                        {
                            default: () => 'Edit',
                        }
                    ),
                    h(
                        NPopconfirm,
                        {
                        onPositiveClick: () => handleDelete(row.id),

                        },
                        {
                        trigger: () => h(
                            'a',
                            {
                            class: 'link cursor-pointer',
                            }, {
                            default: () => 'Delete',
                        }),
                        default: () => `Are you sure to delete ${row.name}?`,
                        }
                    )
                ]
            )
        }
    },
]
</script>
<template>
<n-data-table 
    :columns="columns"
    :data="globalStore.cates"
    :pagination="{pageSize: 10}" />
</template>