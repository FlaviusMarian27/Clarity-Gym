<template>
  <div>
    <Navbar />

    <!-- HERO -->
    <section id="hero" class="min-h-screen bg-bg flex items-center justify-center pt-20">
      <div class="text-center px-6">
        <img src="/logo.png" alt="Clarity Gym" class="h-40 w-40 object-contain mx-auto mb-8" />
        <h1 class="text-6xl font-bold text-text mb-4">Clarity Gym</h1>
        <p class="text-2xl text-primary mb-4">Premium Wellness & Fitness</p>
        <p class="text-lg text-text mb-10 max-w-xl mx-auto">
          Descoperă un loc unde performanța întâlnește eleganța. Antrenori dedicați, echipamente moderne și o comunitate care te inspiră.
        </p>
        <RouterLink
          to="/login"
          class="px-10 py-4 bg-primary hover:bg-secondary text-white text-xl font-semibold rounded-2xl transition-colors duration-200"
        >
          Alătură-te nouă
        </RouterLink>
      </div>
    </section>

  </div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue'
</script>