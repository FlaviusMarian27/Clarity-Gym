<template>
  <div class="min-h-screen bg-bg flex items-center justify-center">
    <div class="bg-card rounded-2xl shadow-lg p-8 w-full max-w-md">

      <!-- Titlu -->
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-text">Clarity Gym</h1>
        <p class="text-primary mt-2">Bine ai revenit!</p>
      </div>

      <!-- Form -->
      <div class="flex flex-col gap-4">
        <input
          type="email"
          placeholder="Email"
          class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
        />
        <input
          type="password"
          placeholder="Parola"
          class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
        />
        <button class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200">
          Autentificare
        </button>
      </div>

      <!-- Link register -->
      <p class="text-center text-text mt-6 text-sm">
        Nu ai cont?
        <RouterLink to="/register" class="text-secondary font-semibold hover:underline">
            Înregistrează-te
        </RouterLink>
      </p>

    </div>
  </div>
</template>