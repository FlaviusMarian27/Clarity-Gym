<template>
  <div class="min-h-screen bg-bg">
    <NavbarClient v-if="role === 'client'" />
    <NavbarTrainer v-else-if="role === 'trainer'" />

    <div class="max-w-3xl mx-auto px-8 pt-36 pb-24">

      <!-- Header -->
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-text mb-4">Profilul Meu</h1>
        <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
      </div>

      <!-- Card profil -->
      <div class="bg-card rounded-2xl shadow-sm p-10">

        <!-- Buton inapoi -->
        <button
          @click="router.back()"
          class="flex items-center gap-2 text-text opacity-60 hover:opacity-100 transition-opacity duration-200 mb-8"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
          </svg>
          Înapoi
        </button>

        <!-- Poza profil -->
        <div class="flex flex-col items-center mb-10">
          <div class="relative">
            <div class="w-32 h-32 rounded-full bg-primary flex items-center justify-center text-white font-bold text-5xl overflow-hidden">
              <img v-if="form.avatar_url" :src="form.avatar_url.startsWith('http') ? form.avatar_url : 'http://localhost:8080' + form.avatar_url" class="w-full h-full object-cover" />
              <span v-else>{{ form.name ? form.name[0].toUpperCase() : '?' }}</span>
            </div>
            <label class="absolute bottom-0 right-0 w-9 h-9 bg-secondary rounded-full flex items-center justify-center cursor-pointer hover:bg-primary transition-colors duration-200">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536M9 13l6.586-6.586a2 2 0 112.828 2.828L11.828 15.828a2 2 0 01-1.414.586H9v-2a2 2 0 01.586-1.414z" />
              </svg>
              <input type="file" accept=".jpg,.jpeg,.png,.webp" class="hidden" @change="onFileSelected" />
            </label>
          </div>
          <p v-if="uploadingAvatar" class="text-primary text-sm mt-3">Se încarcă poza...</p>
          <h2 class="text-2xl font-bold text-text mt-4">{{ form.name }}</h2>
          <p class="text-primary font-medium mt-1">{{ role }}</p>
        </div>

        <!-- Modal Crop -->
        <div v-if="showCropper" class="fixed inset-0 bg-black bg-opacity-70 z-50 flex items-center justify-center">
          <div class="bg-card rounded-2xl p-6 w-full max-w-lg mx-4">
            <h3 class="text-xl font-bold text-text mb-4 text-center">Decupează poza</h3>
            <div class="w-full overflow-hidden rounded-xl" style="max-height: 400px;">
              <img ref="cropperImage" :src="rawImageUrl" class="max-w-full" />
            </div>
            <div class="flex gap-3 mt-6">
              <button
                @click="cancelCrop"
                class="flex-1 py-3 border border-primary/30 text-text rounded-xl font-semibold hover:bg-bg transition-colors duration-200"
              >
                Anulează
              </button>
              <button
                @click="confirmCrop"
                class="flex-1 py-3 bg-primary hover:bg-secondary text-white rounded-xl font-semibold transition-colors duration-200"
              >
                {{ uploadingAvatar ? 'Se încarcă...' : 'Salvează poza' }}
              </button>
            </div>
          </div>
        </div>

        <!-- Formular -->
        <div class="flex flex-col gap-5">

          <!-- Nume -->
          <div>
            <label class="block text-text font-semibold mb-2">Nume complet</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full px-4 py-3 rounded-xl border border-primary/30 bg-bg text-text focus:outline-none focus:border-primary"
              placeholder="Numele tău"
            />
          </div>

          <!-- Bio -->
          <div>
            <label class="block text-text font-semibold mb-2">Despre mine</label>
            <textarea
              v-model="form.bio"
              rows="4"
              class="w-full px-4 py-3 rounded-xl border border-primary/30 bg-bg text-text focus:outline-none focus:border-primary resize-none"
              placeholder="Scrie ceva despre tine..."
            />
          </div>

          <!-- Doar pentru antrenor -->
          <template v-if="role === 'trainer'">
            <div>
              <label class="block text-text font-semibold mb-2">Specialitate</label>
              <input
                v-model="form.specialty"
                type="text"
                class="w-full px-4 py-3 rounded-xl border border-primary/30 bg-bg text-text focus:outline-none focus:border-primary"
                placeholder="ex: Fitness, Yoga, Crossfit..."
              />
            </div>

            <div>
              <label class="block text-text font-semibold mb-2">Ani de experiență</label>
              <input
                v-model="form.experience_years"
                type="number"
                min="0"
                class="w-full px-4 py-3 rounded-xl border border-primary/30 bg-bg text-text focus:outline-none focus:border-primary"
                placeholder="ex: 5"
              />
            </div>
          </template>

          <!-- Buton salvare -->
          <button
            @click="saveProfile"
            :disabled="loading"
            class="mt-4 w-full py-3 bg-primary hover:bg-secondary text-white rounded-xl font-semibold text-lg transition-colors duration-200"
          >
            {{ loading ? 'Se salvează...' : 'Salvează modificările' }}
          </button>

          <p v-if="success" class="text-center text-green-500 font-medium">✓ Profil actualizat cu succes!</p>
          <p v-if="error" class="text-center text-red-400 font-medium">{{ error }}</p>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import Cropper from 'cropperjs'
import 'cropperjs/dist/cropper.css'
import { useAuthStore } from '../stores/auth'
import NavbarClient from '../components/NavbarClient.vue'
import NavbarTrainer from '../components/NavbarTrainer.vue'

const router = useRouter()
const authStore = useAuthStore()
const role = authStore.role

const form = ref({
  name: '',
  bio: '',
  avatar_url: '',
  specialty: '',
  experience_years: 0
})

const loading = ref(false)
const success = ref(false)
const error = ref('')
const uploadingAvatar = ref(false)
const showCropper = ref(false)
const rawImageUrl = ref('')
const cropperImage = ref(null)
let cropperInstance = null

onMounted(async () => {
  try {
    const res = await axios.get('/api/profile')
    form.value = res.data
  } catch (err) {
    console.error(err)
  }
})

async function saveProfile() {
  loading.value = true
  success.value = false
  error.value = ''
  try {
    await axios.put('/api/profile', form.value)
    success.value = true
  } catch (err) {
    error.value = 'Eroare la salvare!'
  } finally {
    loading.value = false
  }
}

function onFileSelected(event) {
  const file = event.target.files[0]
  if (!file) return

  rawImageUrl.value = URL.createObjectURL(file)
  showCropper.value = true

  nextTick(() => {
    if (cropperInstance) {
      cropperInstance.destroy()
      cropperInstance = null
    }
    cropperInstance = new Cropper(cropperImage.value, {
      aspectRatio: 1,
      viewMode: 1,
      autoCropArea: 0.8,
    })
  })
}

function cancelCrop() {
  showCropper.value = false
  if (cropperInstance) {
    cropperInstance.destroy()
    cropperInstance = null
  }
}

async function confirmCrop() {
  if (!cropperInstance) return

  uploadingAvatar.value = true

  const canvas = cropperInstance.getCroppedCanvas({ width: 400, height: 400 })
  canvas.toBlob(async (blob) => {
    const formData = new FormData()
    formData.append('avatar', blob, 'avatar.jpg')

    try {
      const res = await axios.post('/api/profile/avatar', formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
      })
      form.value.avatar_url = res.data.avatar_url
      showCropper.value = false
      cropperInstance.destroy()
      cropperInstance = null
      setTimeout(() => window.location.reload(), 300)
    } catch (err) {
      console.error('Eroare upload:', err)
    } finally {
      uploadingAvatar.value = false
    }
  }, 'image/jpeg')
}
</script>