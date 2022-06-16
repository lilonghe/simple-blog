import { defineStore } from 'pinia'

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
      saveUser(user: User) {
        this.user = user
      },
      clearUser() {
        this.user = null
      }
    },
})