import { defineStore } from 'pinia'
import { getCateListReq, getOptionsReq, getTagListReq, deleteCateReq, createOrEditCateReq, deleteTagReq, createOrEditTagReq } from '@/services'

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
      },
      async deleteCate(id: number) {
        let { code } = await deleteCateReq(id)
        if (!code) {
          window.$message.success('Delete success')
          this.getCates()
        }
      },
      async createOrEditCate(params: any) {
        let { code } = await createOrEditCateReq(params)
        if (!code) {
          if (params.id) {
            window.$message.success('Edit success')
          } else {
            window.$message.success('Create success')
          }
          this.getCates()
          return true
        }
        return false
      },
      async deleteTag(id: number) {
        let { code } = await deleteTagReq(id)
        if (!code) {
          window.$message.success('Delete success')
          this.getTags()
        }
      },
      async createOrEditTag(params: any) {
        let { code } = await createOrEditTagReq(params)
        if (!code) {
          if (params.id) {
            window.$message.success('Edit success')
          } else {
            window.$message.success('Create success')
          }
          this.getTags()
          return true
        }
        return false
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