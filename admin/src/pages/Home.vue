<script setup lang="ts">
import { getPostListReq, getSystemInfoReq } from '@/services'
import { SVGGraph } from 'calendar-graph'
import { NButton, NTag } from 'naive-ui'
import { computed, onMounted, ref } from 'vue'

const systemInfo = ref<any>()
const currentYear = new Date().getFullYear()

onMounted(() => {
  loadData()
  loadAllPost()
})

const loadData = async () => {
  const { data } = await getSystemInfoReq()
  if (data) {
    systemInfo.value = data
  }
}

const loadAllPost = async () => {
  const { data } = await getPostListReq({ pageSize: 999999999, page: 1, archive: true })
  if (!data) return

  const { list } = data
  const graphData: Array<{ date: string; count: number }> = []
  const current = new Date()
  const year = current.getFullYear().toString()
  const month = current.getMonth() + 1
  const day = current.getDate()

  list
    .filter((item: any) => new Date(item.create_time).toJSON().startsWith(year))
    .forEach((item: any) => {
      const currentDay = new Date(item.create_time).toJSON().substring(0, 10)
      const match = graphData.findIndex((record) => record.date === currentDay)
      if (match !== -1) {
        graphData[match].count += 1
        return
      }

      graphData.push({
        date: currentDay,
        count: 1,
      })
    })

  new SVGGraph('#svg-root', graphData, {
    startDate: new Date(`${year}-01-01`),
    endDate: new Date(`${year}-${month}-${day}`),
  })
}

const metrics = computed(() => {
  if (!systemInfo.value) return []

  return [
    {
      label: 'Posts',
      value: systemInfo.value.count.posts,
      hint: 'Long-form entries across draft and published states.',
    },
    {
      label: 'Categories',
      value: systemInfo.value.count.cates,
      hint: 'Primary buckets that keep the archive structured.',
    },
    {
      label: 'Tags',
      value: systemInfo.value.count.tags,
      hint: 'Secondary labels for related writing and discovery.',
    },
  ]
})
</script>

<template>
  <div class="page-stack">
    <section class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">Daily pulse</p>
          <h2 class="section-title">A quiet control room for your writing.</h2>
          <p class="section-copy">
            Keep an eye on output, reader interest and publishing momentum without turning
            the dashboard into a wall of cards.
          </p>
        </div>
        <div class="quick-actions">
          <router-link to="/post/create">
            <n-button type="primary">Write post</n-button>
          </router-link>
          <router-link to="/page/create">
            <n-button secondary>Create page</n-button>
          </router-link>
        </div>
      </div>

      <div v-if="metrics.length" class="metric-grid">
        <article v-for="item in metrics" :key="item.label" class="metric-card">
          <span class="metric-card__label">{{ item.label }}</span>
          <span class="metric-card__value">{{ item.value }}</span>
          <span class="metric-card__hint">{{ item.hint }}</span>
        </article>
      </div>
    </section>

    <div class="panel-grid">
      <section class="page-surface page-surface--padded panel-block">
        <div>
          <p class="section-kicker">Activity</p>
          <h2 class="section-title">{{ currentYear }} publishing rhythm</h2>
          <p class="section-copy">
            The contribution map keeps the homepage grounded in writing frequency instead of
            vanity metrics.
          </p>
        </div>
        <div id="svg-root"></div>
        <div id="tooltip"></div>
      </section>

      <section v-if="systemInfo" class="page-surface page-surface--padded panel-block">
        <div>
          <p class="section-kicker">Readers</p>
          <h2 class="section-title">What people opened this week</h2>
          <p class="section-copy">Use recent attention as a signal for what is resonating.</p>
        </div>

        <ul class="list-plain">
          <li
            v-for="item in systemInfo.count.weekVisitTop"
            :key="item.pathname"
            class="list-row"
          >
            <div>
              <a
                :href="`/post/${item.pathname}.html`"
                class="list-row__title link"
                target="_blank"
              >
                {{ item.title }}
              </a>
              <span class="list-row__meta">/post/{{ item.pathname }}.html</span>
            </div>
            <n-tag size="small">{{ item.count }} views</n-tag>
          </li>
        </ul>
      </section>
    </div>

    <section v-if="systemInfo" class="page-surface page-surface--padded">
      <div class="page-toolbar">
        <div>
          <p class="section-kicker">System</p>
          <h2 class="section-title">Environment and blog identity</h2>
          <p class="section-copy">
            Keep a compact view of platform details close to the content workflow.
          </p>
        </div>
      </div>

      <div class="system-pills">
        <n-tag>System: {{ systemInfo.versions.platform }}</n-tag>
        <n-tag>MySQL: {{ systemInfo.versions.mysqlVersion }}</n-tag>
      </div>
    </section>
  </div>
</template>
