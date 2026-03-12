import { ref } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const isLoggedIn = ref(false)
  const role = ref(null)
  const token = ref(null)

  // Axios base URL
  axios.defaults.baseURL = 'http://localhost:8080'

  function login(userData) {
    user.value = userData
    isLoggedIn.value = true
    role.value = userData.role
    token.value = userData.token

    // Salveaza token in localStorage
    localStorage.setItem('token', userData.token)
    localStorage.setItem('role', userData.role)
    localStorage.setItem('name', userData.name)

    // Adauga token la toate request-urile
    axios.defaults.headers.common['Authorization'] = `Bearer ${userData.token}`
  }

  function logout() {
    user.value = null
    isLoggedIn.value = false
    role.value = null
    token.value = null

    localStorage.removeItem('token')
    localStorage.removeItem('role')
    localStorage.removeItem('name')

    delete axios.defaults.headers.common['Authorization']
  }

  // Restaureaza sesiunea la refresh pagina
  function restoreSession() {
    const savedToken = localStorage.getItem('token')
    const savedRole = localStorage.getItem('role')
    const savedName = localStorage.getItem('name')

    if (savedToken && savedRole) {
      token.value = savedToken
      role.value = savedRole
      isLoggedIn.value = true
      user.value = { name: savedName, role: savedRole }
      axios.defaults.headers.common['Authorization'] = `Bearer ${savedToken}`
    }
  }

  return { user, isLoggedIn, role, token, login, logout, restoreSession }
})