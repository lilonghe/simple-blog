import { defineStore } from 'pinia'
import { getUserReq } from '@/services'

interface ISessionState {
    user: User | null;
}

export const useSesssionStore = defineStore('session', {
    state: () => {
      return <ISessionState> {
        user: null
      }
    },
    actions: {
      async getUser() {
        let { data } = await getUserReq()
        if (data) {
          this.user = data
        }
      },
      clearUser() {
        this.user = null
      }
    },
})