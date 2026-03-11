<template>
  <nav class="fixed top-0 left-0 right-0 z-50 bg-card shadow-sm">
    <div class="w-full px-8 py-4 flex justify-between items-center">

      <!-- Stanga: Logo + Nume -->
      <div class="flex items-center gap-3">
        <img src="/logo.png" alt="Clarity Gym Logo" class="h-14 w-14 object-contain" />
        <span class="text-3xl font-bold text-text">Clarity Gym</span>
      </div>

      <!-- Dreapta: Links + Login -->
      <div class="flex items-center gap-8">
        <a href="#hero" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Acasă</a>
        <a href="#despre" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Despre</a>
        <a href="#abonamente" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Abonamente</a>
        <a href="#recenzii" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Recenzii</a>
        <a href="#locatie" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Locație</a>
        <a href="#suport" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Suport</a>
        <a href="#contact" class="text-text hover:text-primary transition-colors duration-200 text-lg font-medium">Contact</a>

        <RouterLink
          to="/login"
          class="px-6 py-3 rounded-xl bg-primary text-white hover:bg-secondary transition-colors duration-200 text-lg font-semibold"
        >
          Login
        </RouterLink>
      </div>

    </div>
  </nav>
</template>