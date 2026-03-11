<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-text shadow-sm">
    <div class="w-full px-8 py-4 flex justify-between items-center">

      <!-- Logo + Nume -->
      <div class="flex items-center gap-3">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-10 w-10 object-contain" />
        <span class="text-xl font-bold text-white">Clarity Gym <span class="text-primary text-sm font-normal ml-2">Admin Panel</span></span>
      </div>

      <!-- Clopot + Avatar -->
      <div class="flex items-center gap-4">
        <div class="relative">
          <button class="w-10 h-10 rounded-full bg-white bg-opacity-10 flex items-center justify-center hover:bg-opacity-20 transition-colors duration-200 text-white">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <span class="absolute -top-1 -right-1 w-5 h-5 bg-red-400 text-white text-xs rounded-full flex items-center justify-center font-bold">5</span>
        </div>

        <div class="relative">
          <button
            @click="toggleDropdown"
            class="w-10 h-10 rounded-full bg-primary text-white font-bold flex items-center justify-center hover:bg-secondary transition-colors duration-200"
          >
            A
          </button>
          <div v-if="dropdownOpen" class="absolute right-0 mt-2 w-48 bg-card rounded-xl shadow-lg py-2 z-50">
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
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const dropdownOpen = ref(false)

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function handleLogout() {
  authStore.logout()
  router.push('/')
}
</script>