<template>
  <div>
    <NavbarSimple />
    <div class="min-h-screen bg-bg flex items-center justify-center pt-20">
      <div class="bg-card rounded-2xl shadow-lg p-8 w-full max-w-md">

        <div class="text-center mb-8">
          <h1 class="text-3xl font-bold text-text">Clarity Gym</h1>
          <p class="text-primary mt-2">Bine ai revenit!</p>
        </div>

        <!-- Eroare -->
        <div v-if="error" class="bg-red-100 text-red-500 rounded-xl px-4 py-3 text-sm mb-4">
          {{ error }}
        </div>

        <div class="flex flex-col gap-4">
          <input
            v-model="email"
            type="email"
            placeholder="Email"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />
          <input
            v-model="password"
            type="password"
            placeholder="Parolă"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />
          <button
            @click="handleLogin"
            class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200"
          >
            Autentificare
          </button>
        </div>

        <p class="text-center text-text mt-6 text-sm">
          Nu ai cont?
          <RouterLink to="/register" class="text-secondary font-semibold hover:underline">
            Înregistrează-te
          </RouterLink>
        </p>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import NavbarSimple from '../components/NavbarSimple.vue'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const error = ref('')

// Date fake pentru testare - vor fi inlocuite cu API call
const fakeUsers = [
  { email: 'client@test.com', password: '1234', role: 'client', name: 'Flavius' },
  { email: 'trainer@test.com', password: '1234', role: 'trainer', name: 'Alexandru' },
  { email: 'admin@test.com', password: '1234', role: 'admin', name: 'Admin' },
]

function handleLogin() {
  error.value = ''

  if (!email.value || !password.value) {
    error.value = 'Completează toate câmpurile!'
    return
  }

  const user = fakeUsers.find(
    u => u.email === email.value && u.password === password.value
  )

  if (!user) {
    error.value = 'Email sau parolă incorecte!'
    return
  }

  authStore.login(user)

  if (user.role === 'client') router.push('/client')
  else if (user.role === 'trainer') router.push('/trainer')
  else if (user.role === 'admin') router.push('/admin')
}
</script>