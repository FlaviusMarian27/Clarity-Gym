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
        <div>
          <input
            v-model="email"
            type="email"
            placeholder="Email"
            :class="[
              'w-full border rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2',
              emailError ? 'border-red-400 focus:ring-red-300' : 'border-primary focus:ring-secondary'
            ]"
            @blur="validateEmail"
          />
          <p v-if="emailError" class="text-red-400 text-xs mt-1 ml-1">{{ emailError }}</p>
        </div>

        <!-- Parola -->
        <div>
          <div class="relative">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Parolă"
              :class="[
                'w-full border rounded-xl px-4 py-3 pr-12 text-text bg-bg focus:outline-none focus:ring-2',
                passwordError ? 'border-red-400 focus:ring-red-300' : 'border-primary focus:ring-secondary'
              ]"
              @blur="validatePassword"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-text opacity-50 hover:opacity-100"
            >
              <svg v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
              <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
              </svg>
            </button>
          </div>
          <p v-if="passwordError" class="text-red-400 text-xs mt-1 ml-1">{{ passwordError }}</p>
          <!-- Indicatori parola -->
          <div v-if="password" class="flex gap-2 mt-2">
            <span :class="['text-xs px-2 py-1 rounded-full', password.length >= 8 ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-400']">
              Min. 8 caractere
            </span>
            <span :class="['text-xs px-2 py-1 rounded-full', /[A-Z]/.test(password) ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-400']">
              Literă mare
            </span>
            <span :class="['text-xs px-2 py-1 rounded-full', /[0-9]/.test(password) ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-400']">
              Număr
            </span>
            <span :class="['text-xs px-2 py-1 rounded-full', hasSpecialChar ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-400']">
              Caracter special
            </span>
          </div>
        </div>

        <!-- Confirmare parola -->
        <div>
          <div class="relative">
            <input
              v-model="confirmPassword"
              :type="showConfirmPassword ? 'text' : 'password'"
              placeholder="Confirmă parola"
              :class="[
                'w-full border rounded-xl px-4 py-3 pr-12 text-text bg-bg focus:outline-none focus:ring-2',
                confirmError ? 'border-red-400 focus:ring-red-300' : 'border-primary focus:ring-secondary'
              ]"
              @blur="validateConfirm"
            />
            <button
              type="button"
              @click="showConfirmPassword = !showConfirmPassword"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-text opacity-50 hover:opacity-100"
            >
              <svg v-if="!showConfirmPassword" class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
              <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
              </svg>
            </button>
          </div>
          <p v-if="confirmError" class="text-red-400 text-xs mt-1 ml-1">{{ confirmError }}</p>
        </div>

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
import { ref, computed } from 'vue'
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
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const emailError = ref('')
const passwordError = ref('')
const confirmError = ref('')

const hasSpecialChar = computed(() => /[!@#$%^&*()\-_=+[\]{};,.]/.test(password.value))

function validateAge() {
  if (!dataNasterii.value) return true
  const birth = new Date(dataNasterii.value)
  const today = new Date()
  let age = today.getFullYear() - birth.getFullYear()
  const m = today.getMonth() - birth.getMonth()
  if (m < 0 || (m === 0 && today.getDate() < birth.getDate())) age--
  return age >= 16
}

function validateEmail() {
  const emailRegex = /^[^\s@]+@(gmail\.com|yahoo\.com|yahoo\.ro)$/i
  if (!email.value) {
    emailError.value = 'Email-ul este obligatoriu!'
  } else if (!emailRegex.test(email.value)) {
    emailError.value = 'Acceptăm doar adrese Gmail sau Yahoo!'
  } else {
    emailError.value = ''
  }
}

function validatePassword() {
  if (!password.value) {
    passwordError.value = 'Parola este obligatorie!'
  } else if (password.value.length < 8) {
    passwordError.value = 'Parola trebuie să aibă minim 8 caractere!'
  } else if (!/[A-Z]/.test(password.value)) {
    passwordError.value = 'Parola trebuie să conțină cel puțin o literă mare!'
  } else if (!/[0-9]/.test(password.value)) {
    passwordError.value = 'Parola trebuie să conțină cel puțin un număr!'
  } else if (!hasSpecialChar.value) {
    passwordError.value = 'Parola trebuie să conțină cel puțin un caracter special!'
  } else {
    passwordError.value = ''
  }
}

function validateConfirm() {
  if (confirmPassword.value !== password.value) {
    confirmError.value = 'Parolele nu coincid!'
  } else {
    confirmError.value = ''
  }
}

async function handleRegister() {
  error.value = ''
  validateEmail()
  validatePassword()
  validateConfirm()

  if (!nume.value || !prenume.value) {
    error.value = 'Completează numele și prenumele!'
    return
  }

  if (emailError.value || passwordError.value || confirmError.value) return

  if (!validateAge()) {
    error.value = 'Trebuie să ai minim 16 ani pentru a te înregistra!'
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