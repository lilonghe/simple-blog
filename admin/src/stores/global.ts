import { defineStore } from 'pinia'
import { getCateListReq, getOptionsReq, getTagListReq } from '@/services'

interface IGlobalState {
    options: any;
    cates: any;
    tags: any;
}

export const useGlobalStore = defineStore('global', {
    state: () => {
      return <IGlobalState> {
        options: {},
        cates: [],
        tags: [],
      }
    },
    actions: {
      async getOptions() {
        let { data } = await getOptionsReq()
        if (data) {
          let options: any = {};
          data.forEach((item: any) => {
            options[item.key] = item.value
          })
          console.log(options)
          this.options = options
        }
      },
      async getCates() {
        let { data } = await getCateListReq()
        if (data) {
          this.cates = data
        }
      },
      async getTags() {
        let { data } = await getTagListReq()
        if (data) {
          this.tags = data
        }
      }
    },
    getters: {
      tagOptions: (state) => state.tags.map((item: any) => {
          return {
            label: item.name,
            value: item.name
          }
        })
      }
})