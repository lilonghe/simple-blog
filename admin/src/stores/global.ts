import { defineStore } from 'pinia'

interface IGlobalState {
    options: object;
}

export const useGlobalStore = defineStore('global', {
    state: () => {
      return <IGlobalState> {
        options: {}
      }
    },
    actions: {
      saveOptions(options: object) {
        this.options = options
      }
    },
})