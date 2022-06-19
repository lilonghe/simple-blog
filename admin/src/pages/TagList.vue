<script setup lang="ts">
import { NDataTable, NPopconfirm, NInput } from 'naive-ui';
import { h, onMounted, ref, computed } from 'vue';
import { useGlobalStore } from '@/stores/global';
import { RouterLink } from 'vue-router';

const globalStore = useGlobalStore();
const keyword = ref()

onMounted(() => {
    globalStore.getTags()
})

const handleDelete = (id: number) => {
    globalStore.deleteTag(id)
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
                            to: '/tag/edit/' + row.id,
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
        return globalStore.tags.filter((tag: any) => {
            return tag.name.includes(keyword.value) || tag.pathname.includes(keyword.value)
        })
    }
    return globalStore.tags
})
</script>
<template>
<div>
    <div class="mb-2 w-1/3">
        <n-input placeholder="Search" v-model:value="keyword" />
    </div>
    <n-data-table 
        :columns="columns"
        :data="list"
        :pagination="{pageSize: 10}" />
</div>

</template>