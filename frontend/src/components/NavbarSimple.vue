<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-8 py-4 flex justify-between items-center">

      <!-- Logo + Nume -->
      <RouterLink to="/" class="flex items-center gap-3">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-12 w-12 object-contain" />
        <span class="text-2xl font-bold text-text">Clarity Gym</span>
      </RouterLink>

      <!-- Inapoi la Home -->
      <RouterLink
        to="/"
        class="px-6 py-2 rounded-xl border border-primary text-text hover:bg-primary hover:text-white transition-colors duration-200 font-medium"
      >
        Înapoi la Home
      </RouterLink>

    </div>
  </nav>
</template>