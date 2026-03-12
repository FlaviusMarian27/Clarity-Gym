<template>
  <NavbarSimple />
    <div class="min-h-screen bg-bg flex items-center justify-center pt-20">
      <div class="bg-card rounded-2xl shadow-lg p-8 w-full max-w-md">

        <!-- Titlu -->
        <div class="text-center mb-8">
          <h1 class="text-3xl font-bold text-text">Clarity Gym</h1>
          <p class="text-primary mt-2">Creează un cont nou</p>
        </div>

        <!-- Eroare -->
        <div v-if="error" class="bg-red-100 text-red-500 rounded-xl px-4 py-3 text-sm mb-4">
          {{ error }}
        </div>

        <!-- Form -->
        <div class="flex flex-col gap-4">
          <input
            v-model="nume"
            type="text"
            placeholder="Nume"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />
          <input
            v-model="prenume"
            type="text"
            placeholder="Prenume"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Data nasterii -->
          <div class="flex flex-col gap-1">
            <label class="text-sm text-text ml-1">Data nașterii</label>
            <input
              v-model="dataNasterii"
              type="date"
              class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
            />
          </div>

          <!-- Email -->
          <input
            v-model="email"
            type="email"
            placeholder="Email"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Parola -->
          <input
            v-model="password"
            type="password"
            placeholder="Parolă"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Confirmare parola -->
          <input
            v-model="confirmPassword"
            type="password"
            placeholder="Confirmă parola"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary"
          />

          <!-- Rol -->
          <select 
            v-model="role"
            class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary">
            <option value="">Selectează rolul</option>
            <option value="client">Client</option>
            <option value="trainer">Antrenor</option>
          </select>

          <button
            @click="handleRegister"
            :disabled="loading"
            class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200 disabled:opacity-50"
            >
            {{ loading ? 'Se încarcă...' : 'Înregistrare' }}
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
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import NavbarSimple from '../components/NavbarSimple.vue'
import axios from 'axios'

const router = useRouter()
const authStore = useAuthStore()

const nume = ref('')
const prenume = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const role = ref('client')
const dataNasterii = ref('')
const error = ref('')
const loading = ref(false)

async function handleRegister() {
  error.value = ''

  if (!nume.value || !prenume.value || !email.value || !password.value || !confirmPassword.value) {
    error.value = 'Completează toate câmpurile!'
    return
  }

  if (password.value !== confirmPassword.value) {
    error.value = 'Parolele nu coincid!'
    return
  }

  loading.value = true

  try {
    const response = await axios.post('/api/auth/register', {
      email: email.value,
      password: password.value,
      name: `${nume.value} ${prenume.value}`,
      role: role.value
    })

    const data = response.data
    authStore.login(data)

    if (data.role === 'client') router.push('/client')
    else if (data.role === 'trainer') router.push('/trainer')

  } catch (err) {
    error.value = 'Email-ul este deja folosit!'
  } finally {
    loading.value = false
  }
}
</script>