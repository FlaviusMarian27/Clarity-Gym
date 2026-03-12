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
              :disabled="loading"
              class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200 disabled:opacity-50"
            >
              {{ loading ? 'Se încarcă...' : 'Autentificare' }}
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
import axios from 'axios'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function handleLogin() {
  error.value = ''

  if (!email.value || !password.value) {
    error.value = 'Completează toate câmpurile!'
    return
  }

  loading.value = true

  try {
    const response = await axios.post('/api/auth/login', {
      email: email.value,
      password: password.value
    })

    const data = response.data
    authStore.login(data)

    if (data.role === 'client') router.push('/client')
    else if (data.role === 'trainer') router.push('/trainer')
    else if (data.role === 'admin') router.push('/admin')

  } catch (err) {
    error.value = 'Email sau parolă incorecte!'
  } finally {
    loading.value = false
  }
}
</script>