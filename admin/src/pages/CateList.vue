<script setup lang="ts">
import { NDataTable, NPopconfirm, NInput, NButton } from 'naive-ui';
import { computed, h, onMounted, ref } from 'vue';
import { useGlobalStore } from '@/stores/global';
import { RouterLink } from 'vue-router';

const globalStore = useGlobalStore();
const keyword = ref()

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
        title: 'Post count',
        key: 'post_count'
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
                            to: '/category/edit/' + row.id,
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

const list = computed(() => {
    if (keyword.value) {
        return globalStore.cates.filter((cate: any) => {
            return cate.name.includes(keyword.value) || cate.pathname.includes(keyword.value)
        })
    }
    return globalStore.cates
})
</script>
<template>
<div>
    <div class="mb-2 w-1/3">
        <n-input placeholder="Search" v-model:value="keyword" />
    </div>
    <div class="mb-2">
        <router-link to="/category/create"><n-button type="primary">Create</n-button></router-link>
    </div>
    <n-data-table 
        :columns="columns"
        :data="list"
        :pagination="{pageSize: 10}" />
</div>

</template>