<template>
  <NavbarSimple />
    <div class="min-h-screen bg-bg flex items-center justify-center pt-20">
      <div class="bg-card rounded-2xl shadow-lg p-8 w-full max-w-md">

        <!-- Titlu -->
        <div class="text-center mb-8">
          <h1 class="text-3xl font-bold text-text">Clarity Gym</h1>
          <p class="text-primary mt-2">Creează un cont nou</p>
        </div>

        <!-- Form -->
        <div class="flex flex-col gap-4">
          <input
            type="text"
            placeholder="Nume"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />
          <input
            type="text"
            placeholder="Prenume"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Data nasterii -->
          <div class="flex flex-col gap-1">
            <label class="text-sm text-text ml-1">Data nașterii</label>
            <input
              type="date"
              class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
            />
          </div>

          <!-- Email -->
          <input
            type="email"
            placeholder="Email"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Parola -->
          <input
            type="password"
            placeholder="Parolă"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Confirmare parola -->
          <input
            type="password"
            placeholder="Confirmă parola"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Rol -->
          <select class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary">
            <option value="">Selectează rolul</option>
            <option value="client">Client</option>
            <option value="trainer">Antrenor</option>
          </select>

          <button class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200">
            Înregistrare
          </button>
        </div>

        <!-- Link login -->
        <p class="text-center text-text mt-6 text-sm">
          Ai deja cont?
          <RouterLink to="/" class="text-secondary font-semibold hover:underline">
              Autentifică-te
          </RouterLink>
        </p>

      </div>
    </div>
</template>

<script setup>
  import NavbarSimple from '../components/NavbarSimple.vue'
</script>