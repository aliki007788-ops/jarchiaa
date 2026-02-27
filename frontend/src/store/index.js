import { reactive } from 'vue'

export const store = reactive({
  state: {
    currentRole: 'user',
    isDarkMode: true,
    isDrawerOpen: false,
    user: {
      name: 'علی رضایی',
      wallet: '۲,۴۵۰,۰۰۰'
    }
  },

  mutations: {
    toggleTheme(state) {
      state.isDarkMode = !state.isDarkMode
      document.body.classList.toggle('light-mode', !state.isDarkMode)
    },
    toggleDrawer(state) {
      state.isDrawerOpen = !state.isDrawerOpen
    },
    setRole(state, role) {
      state.currentRole = role
    }
  },

  actions: {
    switchRole({ state, mutations }) {
      if (state.currentRole === 'user') {
        mutations.setRole('business')
      } else if (state.currentRole === 'business') {
        mutations.setRole('admin')
      } else {
        mutations.setRole('user')
      }
    }
  }
})

export function useStore() {
  return store
}
