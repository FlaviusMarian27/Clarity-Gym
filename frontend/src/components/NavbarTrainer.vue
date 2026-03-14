<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-8 py-4 flex justify-between items-center">

      <!-- Logo + Nume -->
      <div @click="goHome" class="flex items-center gap-3 cursor-pointer">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-12 w-12 object-contain" />
        <span class="text-2xl font-bold text-text">Clarity Gym</span>
      </div>

      <!-- Links + Clopot + Avatar -->
      <div class="flex items-center gap-8">
        <span @click="goHome" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Acasă</span>
        <span @click="navigateTo('despre')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Despre</span>
        <span @click="navigateTo('cereri')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Cereri</span>
        <span @click="navigateTo('clienti')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Clienții Mei</span>
        <span @click="navigateTo('recenzii')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Recenzii</span>
        <span @click="navigateTo('locatie')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Locație</span>
        <span @click="navigateTo('suport')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Suport</span>
        <span @click="navigateTo('contact')" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium cursor-pointer">Contact</span>

        <div class="flex items-center gap-4">

          <!-- Clopot -->
          <div @click="navigateTo('cereri')" class="relative cursor-pointer">
            <button class="w-11 h-11 rounded-full bg-bg flex items-center justify-center hover:bg-primary hover:text-white transition-colors duration-200 text-text">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
            </button>
            <span
              v-if="pendingCount > 0"
              class="absolute -top-1 -right-1 w-5 h-5 bg-red-400 text-white text-xs rounded-full flex items-center justify-center font-bold">
              {{ pendingCount }}
            </span>
          </div>

          <!-- Avatar -->
          <div class="relative">
            <button
              @click="toggleDropdown"
              class="w-11 h-11 rounded-full bg-secondary text-white font-bold text-lg flex items-center justify-center hover:bg-primary transition-colors duration-200 overflow-hidden"
            >
              <img v-if="avatarUrl" :src="avatarUrl.startsWith('http') ? avatarUrl : 'http://localhost:8080' + avatarUrl" class="w-full h-full object-cover" />
              <span v-else>{{ initial }}</span>
            </button>
            <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-card rounded-xl shadow-lg py-2 z-50">
              <button @click="router.push('/profile')" class="w-full text-left px-4 py-2 text-text hover:bg-bg transition-colors duration-200">
                Profilul meu
              </button>
              <button
                @click="handleLogout"
                class="w-full text-left px-4 py-2 text-red-400 hover:bg-bg transition-colors duration-200"
              >
                Deconectare
              </button>
            </div>
          </div>

        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import axios from 'axios'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const dropdownOpen = ref(false)
const initial = ref('T')
const pendingCount = ref(0)
const avatarUrl = ref('')

onMounted(async () => {
  try {
    const [requestsRes, profileRes] = await Promise.all([
      axios.get('/api/collaborations/my'),
      axios.get('/api/profile')
    ])
    pendingCount.value = requestsRes.data.filter(r => r.status === 'pending').length
    avatarUrl.value = profileRes.data.avatar_url
    initial.value = profileRes.data.name ? profileRes.data.name[0].toUpperCase() : 'T'
  } catch (err) {
    console.error(err)
  }
})

function navigateTo(section) {
  if (route.path === '/trainer') {
    document.getElementById(section)?.scrollIntoView({ behavior: 'smooth' })
  } else {
    router.push('/trainer#' + section)
  }
}

function goHome() {
  if (route.path === '/trainer') {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } else {
    router.push('/trainer')
  }
}

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function handleLogout() {
  authStore.logout()
  router.push('/')
}
</script>