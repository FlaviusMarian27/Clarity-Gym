<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-4 md:px-8 py-4 flex justify-between items-center">

      <!-- Logo + Nume -->
      <div class="flex items-center gap-3">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-10 w-10 md:h-14 md:w-14 object-contain" />
        <span class="text-xl md:text-3xl font-bold text-text">Clarity Gym</span>
      </div>

      <!-- Desktop Links + Login -->
      <div class="hidden lg:flex items-center gap-6">
        <a href="#hero" class="text-text hover:text-primary transition-colors duration-200 font-medium">Acasă</a>
        <a href="#despre" class="text-text hover:text-primary transition-colors duration-200 font-medium">Despre</a>
        <a href="#abonamente" class="text-text hover:text-primary transition-colors duration-200 font-medium">Abonamente</a>
        <a href="#recenzii" class="text-text hover:text-primary transition-colors duration-200 font-medium">Recenzii</a>
        <a href="#locatie" class="text-text hover:text-primary transition-colors duration-200 font-medium">Locație</a>
        <a href="#suport" class="text-text hover:text-primary transition-colors duration-200 font-medium">Suport</a>
        <a href="#contact" class="text-text hover:text-primary transition-colors duration-200 font-medium">Contact</a>
        <RouterLink to="/login" class="px-5 py-2 rounded-xl bg-primary text-white hover:bg-secondary transition-colors duration-200 font-semibold">
          Login
        </RouterLink>
      </div>

      <!-- Mobile: Login + Hamburger -->
      <div class="flex lg:hidden items-center gap-3">
        <RouterLink to="/login" class="px-4 py-2 rounded-xl bg-primary text-white hover:bg-secondary transition-colors duration-200 text-sm font-semibold">
          Login
        </RouterLink>
        <button @click="menuOpen = !menuOpen" class="w-10 h-10 flex items-center justify-center rounded-xl bg-bg hover:bg-primary hover:text-white transition-colors duration-200 text-text">
          <svg v-if="!menuOpen" class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

    </div>

    <!-- Mobile Menu -->
    <div v-if="menuOpen" class="lg:hidden bg-card border-t border-bg px-4 py-4 flex flex-col gap-3">
      <a href="#hero" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Acasă</a>
      <a href="#despre" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Despre</a>
      <a href="#abonamente" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Abonamente</a>
      <a href="#recenzii" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Recenzii</a>
      <a href="#locatie" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Locație</a>
      <a href="#suport" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Suport</a>
      <a href="#contact" @click="menuOpen = false" class="text-text hover:text-primary transition-colors duration-200 font-medium py-2">Contact</a>
    </div>

  </nav>
</template>

<script setup>
import { ref } from 'vue'
const menuOpen = ref(false)
</script>