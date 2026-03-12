<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-8 py-4 flex justify-between items-center">

        <!-- Logo + Nume -->
        <div class="flex items-center gap-3">
          <img src="/logo.png" alt="Clarity Gym Logo" class="h-12 w-12 object-contain" />
          <span class="text-2xl font-bold text-text">Clarity Gym</span>
        </div>

        <!-- Links + Avatar -->
        <div class="flex items-center gap-8">
          <a href="#hero" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Acasă</a>
          <a href="#despre" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Despre</a>
          <a href="#antrenori" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Antrenori</a>
          <a href="#abonamente" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Abonamente</a>
          <a href="#recenzii" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Recenzii</a>
          <a href="#locatie" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Locație</a>
          <a href="#suport" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Suport</a>
          <a href="#contact" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Contact</a>

          <!-- Clopot + Avatar -->
          <div class="flex items-center gap-4">

            <!-- Clopot -->
            <a href="#antrenori" class="relative" @click="handleBellClick">
              <button class="w-11 h-11 rounded-full bg-bg flex items-center justify-center hover:bg-primary hover:text-white transition-colors duration-200 text-text">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
              </button>
              <span
                v-if="notifCount > 0"
                class="absolute -top-1 -right-1 w-5 h-5 bg-red-400 text-white text-xs rounded-full flex items-center justify-center font-bold">
                {{ notifCount }}
              </span>
            </a>

            <!-- Avatar -->
            <div class="relative">
              <button
                @click="toggleDropdown"
                class="w-11 h-11 rounded-full bg-primary text-white font-bold text-lg flex items-center justify-center hover:bg-secondary transition-colors duration-200"
              >
                {{ initial }}
              </button>
              <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-card rounded-xl shadow-lg py-2 z-50">
                <button class="w-full text-left px-4 py-2 text-text hover:bg-bg transition-colors duration-200">
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
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import axios from 'axios'

const router = useRouter()
const authStore = useAuthStore()
const dropdownOpen = ref(false)
const initial = ref('F')
const notifCount = ref(0)
const seen = ref(false)

onMounted(async () => {
  try {
    const response = await axios.get('/api/collaborations/status')
    notifCount.value = response.data.filter(r => r.status !== 'pending').length
  } catch (err) {
    console.error(err)
  }
})

function handleBellClick() {
  seen.value = true
  notifCount.value = 0
}

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function handleLogout() {
  authStore.logout()
  router.push('/')
}

const myRequests = ref([])

onMounted(async () => {
  try {
    const response = await axios.get('/api/collaborations/status')
    // Arata notificare doar daca exista cereri accepted sau rejected
    notifCount.value = response.data.filter(
      r => r.status === 'accepted' || r.status === 'rejected'
    ).length
  } catch (err) {
    console.error(err)
  }
})

function getRequestStatus(trainerId) {
  const req = myRequests.value.find(r => r.trainer_id === trainerId)
  return req ? req.status : null
}
</script>