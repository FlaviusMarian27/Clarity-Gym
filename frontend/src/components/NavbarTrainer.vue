<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-4 md:px-8 py-4 flex justify-between items-center">

      <!-- Logo + Nume -->
      <div @click="goHome" class="flex items-center gap-3 cursor-pointer">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-10 w-10 md:h-12 md:w-12 object-contain" />
        <span class="text-xl md:text-2xl font-bold text-text">Clarity Gym</span>
      </div>

      <!-- Desktop Links -->
      <div class="hidden lg:flex items-center gap-6">
        <span @click="goHome" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Acasă</span>
        <span @click="navigateTo('despre')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Despre</span>
        <span @click="navigateTo('cereri')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Cereri</span>
        <span @click="navigateTo('clienti')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Clienții Mei</span>
        <span @click="navigateTo('recenzii')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Recenzii</span>
        <span @click="navigateTo('locatie')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Locație</span>
        <span @click="navigateTo('suport')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Suport</span>
        <span @click="navigateTo('contact')" class="text-text hover:text-primary transition-colors duration-200 font-medium cursor-pointer">Contact</span>

        <!-- Clopot -->
        <div @click="navigateTo('cereri')" class="relative cursor-pointer">
          <button class="w-10 h-10 rounded-full bg-bg flex items-center justify-center hover:bg-primary hover:text-white transition-colors duration-200 text-text">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <span v-if="pendingCount > 0" class="absolute -top-1 -right-1 w-5 h-5 bg-red-400 text-white text-xs rounded-full flex items-center justify-center font-bold">{{ pendingCount }}</span>
        </div>

        <!-- Avatar -->
        <div class="relative">
          <button @click="toggleDropdown" class="w-10 h-10 rounded-full bg-secondary text-white font-bold text-lg flex items-center justify-center hover:bg-primary transition-colors duration-200 overflow-hidden">
            <img v-if="avatarUrl" :src="avatarUrl.startsWith('http') ? avatarUrl : 'http://localhost:8080' + avatarUrl" class="w-full h-full object-cover" />
            <span v-else>{{ initial }}</span>
          </button>
          <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-card rounded-xl shadow-lg py-2 z-50">
            <button @click="router.push('/profile')" class="w-full text-left px-4 py-2 text-text hover:bg-bg transition-colors duration-200">Profilul meu</button>
            <button @click="handleLogout" class="w-full text-left px-4 py-2 text-red-400 hover:bg-bg transition-colors duration-200">Deconectare</button>
          </div>
        </div>
      </div>

      <!-- Mobile: Clopot + Avatar + Hamburger -->
      <div class="flex lg:hidden items-center gap-3">
        <div @click="navigateTo('cereri')" class="relative cursor-pointer">
          <button class="w-9 h-9 rounded-full bg-bg flex items-center justify-center text-text">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <span v-if="pendingCount > 0" class="absolute -top-1 -right-1 w-4 h-4 bg-red-400 text-white text-xs rounded-full flex items-center justify-center font-bold">{{ pendingCount }}</span>
        </div>
        <button @click="toggleDropdown" class="w-9 h-9 rounded-full bg-secondary text-white font-bold flex items-center justify-center overflow-hidden">
          <img v-if="avatarUrl" :src="avatarUrl.startsWith('http') ? avatarUrl : 'http://localhost:8080' + avatarUrl" class="w-full h-full object-cover" />
          <span v-else class="text-sm">{{ initial }}</span>
        </button>
        <button @click="menuOpen = !menuOpen" class="w-9 h-9 flex items-center justify-center rounded-xl bg-bg text-text">
          <svg v-if="!menuOpen" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" /></svg>
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" /></svg>
        </button>
      </div>

    </div>

    <!-- Dropdown Avatar Mobile -->
    <div v-if="dropdownOpen" class="lg:hidden bg-card border-t border-bg px-4 py-2">
      <button @click="router.push('/profile'); dropdownOpen = false" class="w-full text-left px-2 py-2 text-text hover:bg-bg transition-colors duration-200">Profilul meu</button>
      <button @click="handleLogout" class="w-full text-left px-2 py-2 text-red-400 hover:bg-bg transition-colors duration-200">Deconectare</button>
    </div>

    <!-- Mobile Menu -->
    <div v-if="menuOpen" class="lg:hidden bg-card border-t border-bg px-4 py-4 flex flex-col gap-2">
      <span @click="goHome; menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Acasă</span>
      <span @click="navigateTo('despre'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Despre</span>
      <span @click="navigateTo('cereri'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Cereri</span>
      <span @click="navigateTo('clienti'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Clienții Mei</span>
      <span @click="navigateTo('recenzii'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Recenzii</span>
      <span @click="navigateTo('locatie'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Locație</span>
      <span @click="navigateTo('suport'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Suport</span>
      <span @click="navigateTo('contact'); menuOpen = false" class="text-text hover:text-primary font-medium cursor-pointer py-2">Contact</span>
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
const menuOpen = ref(false)
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
  if (dropdownOpen.value) menuOpen.value = false
}

function handleLogout() {
  authStore.logout()
  router.push('/')
}
</script>