import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const isLoggedIn = ref(false)
  const role = ref(null) // 'client', 'trainer', 'admin'

  function login(userData) {
    user.value = userData
    isLoggedIn.value = true
    role.value = userData.role
  }

  function logout() {
    user.value = null
    isLoggedIn.value = false
    role.value = null
  }

  return { user, isLoggedIn, role, login, logout }
})