<template>
  <div class="min-h-screen bg-bg">
    <NavbarAdmin />

    <div class="flex pt-20">

      <!-- SIDEBAR -->
      <aside class="fixed left-0 top-20 h-full w-64 bg-card shadow-sm px-4 py-8 flex flex-col gap-2">
        <button
          v-for="item in menuItems"
          :key="item.key"
          @click="activeSection = item.key"
          :class="[
            'w-full text-left px-4 py-3 rounded-xl font-medium transition-colors duration-200 flex items-center gap-3',
            activeSection === item.key
              ? 'bg-primary text-white'
              : 'text-text hover:bg-bg'
          ]"
        >
          <span>{{ item.icon }}</span>
          <span>{{ item.label }}</span>
        </button>
      </aside>

      <!-- CONTINUT PRINCIPAL -->
      <main class="ml-64 flex-1 p-8">

        <!-- UTILIZATORI -->
        <div v-if="activeSection === 'users'">
          <div class="mb-8">
            <h1 class="text-3xl font-bold text-text">Utilizatori</h1>
            <p class="text-text opacity-60 mt-1">Gestionează toți utilizatorii platformei</p>
          </div>

          <div class="flex gap-4 mb-6">
            <button
              v-for="filter in userFilters"
              :key="filter.key"
              @click="activeUserFilter = filter.key"
              :class="[
                'px-4 py-2 rounded-xl text-sm font-medium transition-colors duration-200',
                activeUserFilter === filter.key
                  ? 'bg-primary text-white'
                  : 'bg-card text-text hover:bg-primary hover:text-white'
              ]"
            >
              {{ filter.label }}
            </button>
          </div>

          <div class="bg-card rounded-2xl shadow-sm overflow-hidden">
            <table class="w-full">
              <thead class="bg-bg">
                <tr>
                  <th class="text-left px-6 py-4 text-text font-semibold text-sm">Nume</th>
                  <th class="text-left px-6 py-4 text-text font-semibold text-sm">Email</th>
                  <th class="text-left px-6 py-4 text-text font-semibold text-sm">Rol</th>
                  <th class="text-left px-6 py-4 text-text font-semibold text-sm">Status</th>
                  <th class="text-left px-6 py-4 text-text font-semibold text-sm">Acțiuni</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="user in filteredUsers"
                  :key="user.id"
                  class="border-t border-bg hover:bg-bg transition-colors duration-150"
                >
                  <td class="px-6 py-4">
                    <div class="flex items-center gap-3">
                      <div class="w-9 h-9 rounded-full bg-primary flex items-center justify-center text-white font-bold text-sm">
                        {{ user.name[0] }}
                      </div>
                      <span class="text-text font-medium">{{ user.name }}</span>
                    </div>
                  </td>
                  <td class="px-6 py-4 text-text opacity-70 text-sm">{{ user.email }}</td>
                  <td class="px-6 py-4">
                    <span :class="[
                      'px-3 py-1 rounded-full text-xs font-semibold',
                      user.role === 'client' ? 'bg-blue-100 text-blue-600' : 'bg-green-100 text-green-600'
                    ]">
                      {{ user.role === 'client' ? 'Client' : 'Antrenor' }}
                    </span>
                  </td>
                  <td class="px-6 py-4">
                    <span :class="[
                      'px-3 py-1 rounded-full text-xs font-semibold',
                      user.is_suspended ? 'bg-red-100 text-red-500' : 'bg-green-100 text-green-600'
                    ]">
                      {{ user.is_suspended ? 'Suspendat' : 'Activ' }}
                    </span>
                  </td>
                  <td class="px-6 py-4">
                    <div class="flex gap-2">
                      <button
                        @click="toggleSuspend(user)"
                        :class="[
                          'px-3 py-1 rounded-lg text-xs font-semibold transition-colors duration-200',
                          user.is_suspended
                            ? 'bg-green-100 text-green-600 hover:bg-green-200'
                            : 'bg-yellow-100 text-yellow-600 hover:bg-yellow-200'
                        ]"
                      >
                        {{ user.is_suspended ? 'Activează' : 'Suspendă' }}
                      </button>
                      <button
                        @click="deleteUser(user.id)"
                        class="px-3 py-1 rounded-lg text-xs font-semibold bg-red-100 text-red-500 hover:bg-red-200 transition-colors duration-200"
                      >
                        Șterge
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- PLANURI ABONAMENTE -->
        <div v-if="activeSection === 'plans'">
          <div class="mb-8 flex justify-between items-center">
            <div>
              <h1 class="text-3xl font-bold text-text">Planuri Abonamente</h1>
              <p class="text-text opacity-60 mt-1">Gestionează planurile tarifare</p>
            </div>
            <button
              @click="showAddPlan = true"
              class="px-6 py-3 bg-primary hover:bg-secondary text-white rounded-xl font-semibold transition-colors duration-200"
            >
              + Adaugă Plan
            </button>
          </div>

          <div class="flex flex-col gap-4">
            <div
              v-for="plan in plans"
              :key="plan.id"
              class="bg-card rounded-2xl p-6 flex justify-between items-center shadow-sm"
            >
              <div>
                <span class="text-xs font-semibold text-primary uppercase tracking-widest">{{ plan.type }}</span>
                <h3 class="text-xl font-bold text-text mt-1">{{ plan.name }}</h3>
                <p class="text-text opacity-60 text-sm mt-1">{{ plan.description }}</p>
              </div>
              <div class="flex items-center gap-6">
                <span class="text-2xl font-bold text-text">{{ plan.price }} RON<span class="text-sm font-normal opacity-60">/lună</span></span>
                <div class="flex gap-2">
                  <button class="px-4 py-2 bg-secondary bg-opacity-20 text-secondary rounded-xl text-sm font-semibold hover:bg-opacity-30 transition-colors duration-200">
                    Editează
                  </button>
                  <button
                    @click="deletePlan(plan.id)"
                    class="px-4 py-2 bg-red-100 text-red-500 rounded-xl text-sm font-semibold hover:bg-red-200 transition-colors duration-200"
                  >
                    Șterge
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Modal Adauga Plan -->
          <div v-if="showAddPlan" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
            <div class="bg-card rounded-2xl p-8 w-full max-w-md shadow-xl">
              <h2 class="text-2xl font-bold text-text mb-6">Adaugă Plan Nou</h2>
              <div class="flex flex-col gap-4">
                <input v-model="newPlan.name" type="text" placeholder="Numele planului" class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary" />
                <input v-model="newPlan.type" type="text" placeholder="Tip (Buget / Standard / VIP)" class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary" />
                <input v-model="newPlan.price" type="number" placeholder="Preț (RON)" class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary" />
                <textarea v-model="newPlan.description" placeholder="Descriere" rows="3" class="border border-primary rounded-xl px-4 py-3 text-text bg-bg focus:outline-none focus:ring-2 focus:ring-secondary resize-none"></textarea>
                <div class="flex gap-3">
                  <button @click="addPlan" class="flex-1 bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200">
                    Adaugă
                  </button>
                  <button @click="showAddPlan = false" class="flex-1 border border-primary text-text font-semibold py-3 rounded-xl hover:bg-bg transition-colors duration-200">
                    Anulează
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- SUPORT -->
        <div v-if="activeSection === 'support'">
          <div class="mb-8">
            <h1 class="text-3xl font-bold text-text">Mesaje Suport</h1>
            <p class="text-text opacity-60 mt-1">Mesajele primite de la utilizatori</p>
          </div>

          <!-- Filtre suport -->
          <div class="flex gap-4 mb-6">
            <button
              @click="activeSupportFilter = 'all'"
              :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors duration-200', activeSupportFilter === 'all' ? 'bg-primary text-white' : 'bg-card text-text hover:bg-primary hover:text-white']"
            >Toate</button>
            <button
              @click="activeSupportFilter = 'open'"
              :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors duration-200', activeSupportFilter === 'open' ? 'bg-primary text-white' : 'bg-card text-text hover:bg-primary hover:text-white']"
            >Deschise</button>
            <button
              @click="activeSupportFilter = 'closed'"
              :class="['px-4 py-2 rounded-xl text-sm font-medium transition-colors duration-200', activeSupportFilter === 'closed' ? 'bg-primary text-white' : 'bg-card text-text hover:bg-primary hover:text-white']"
            >Rezolvate</button>
          </div>

          <div v-if="filteredSupport.length === 0" class="text-center py-16">
            <p class="text-text opacity-50 text-lg">Nu există mesaje momentan.</p>
          </div>

          <div class="flex flex-col gap-4">
            <div
              v-for="msg in filteredSupport"
              :key="msg.id"
              class="bg-card rounded-2xl p-6 shadow-sm"
            >
              <div class="flex justify-between items-start mb-3">
                <div>
                  <h4 class="font-bold text-text">{{ msg.name }}</h4>
                  <p class="text-text opacity-60 text-sm">{{ msg.email }}</p>
                  <p class="text-text opacity-40 text-xs mt-1">{{ msg.created_at.slice(0, 10) }} · {{ msg.role }}</p>
                </div>
                <div class="flex items-center gap-2">
                  <span class="px-3 py-1 bg-primary bg-opacity-10 text-primary rounded-full text-xs font-semibold">{{ msg.category }}</span>
                  <span v-if="msg.status === 'closed'" class="px-3 py-1 bg-green-100 text-green-600 rounded-full text-xs font-semibold">Rezolvat</span>
                  <span v-else class="px-3 py-1 bg-yellow-100 text-yellow-600 rounded-full text-xs font-semibold">Deschis</span>
                </div>
              </div>
              <p class="text-text opacity-70 text-sm leading-relaxed">{{ msg.message }}</p>
              <div class="flex gap-2 mt-4">
                <button
                  v-if="msg.status === 'open'"
                  @click="closeRequest(msg.id)"
                  class="px-4 py-2 bg-green-100 text-green-600 rounded-xl text-sm font-semibold hover:bg-green-200 transition-colors duration-200"
                >
                  Marchează rezolvat
                </button>
              </div>
            </div>
          </div>
        </div>

      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import NavbarAdmin from '../components/NavbarAdmin.vue'
import axios from 'axios'

const activeSection = ref('users')
const activeUserFilter = ref('all')
const activeSupportFilter = ref('all')
const showAddPlan = ref(false)

const menuItems = [
  { key: 'users', label: 'Utilizatori', icon: '👥' },
  { key: 'plans', label: 'Abonamente', icon: '💳' },
  { key: 'support', label: 'Suport', icon: '🛠️' },
]

const userFilters = [
  { key: 'all', label: 'Toți' },
  { key: 'client', label: 'Clienți' },
  { key: 'trainer', label: 'Antrenori' },
]

const users = ref([])
const supportMessages = ref([])

const plans = ref([
  { id: 1, type: 'Buget', name: 'Dimineața', price: 180, description: 'Acces 06:00 - 12:00' },
  { id: 2, type: 'Standard', name: 'Oricând', price: 220, description: 'Acces nelimitat la sală' },
  { id: 3, type: 'VIP', name: 'Rezultate Garantate', price: 1000, description: 'Antrenor personal + plan nutrițional' },
])

const newPlan = ref({ name: '', type: '', price: '', description: '' })

const filteredUsers = computed(() => {
  if (activeUserFilter.value === 'all') return users.value
  return users.value.filter(u => u.role === activeUserFilter.value)
})

const filteredSupport = computed(() => {
  if (activeSupportFilter.value === 'all') return supportMessages.value
  return supportMessages.value.filter(m => m.status === activeSupportFilter.value)
})

async function fetchUsers() {
  try {
    const response = await axios.get('/api/admin/users')
    users.value = response.data
  } catch (err) {
    console.error('Eroare la încărcarea userilor:', err)
  }
}

async function fetchSupport() {
  try {
    const response = await axios.get('/api/support')
    supportMessages.value = response.data
  } catch (err) {
    console.error('Eroare la încărcarea suportului:', err)
  }
}

async function toggleSuspend(user) {
  try {
    const response = await axios.put(`/api/admin/users/${user.id}/suspend`)
    user.is_suspended = response.data.suspended
  } catch (err) {
    console.error('Eroare la suspendare:', err)
  }
}

async function deleteUser(id) {
  try {
    await axios.delete(`/api/admin/users/${id}`)
    users.value = users.value.filter(u => u.id !== id)
  } catch (err) {
    console.error('Eroare la ștergere:', err)
  }
}

async function closeRequest(id) {
  try {
    await axios.put(`/api/support/close?id=${id}`)
    const msg = supportMessages.value.find(m => m.id === id)
    if (msg) msg.status = 'closed'
  } catch (err) {
    console.error('Eroare:', err)
  }
}

function deletePlan(id) {
  plans.value = plans.value.filter(p => p.id !== id)
}

function addPlan() {
  if (!newPlan.value.name || !newPlan.value.price) return
  plans.value.push({ id: Date.now(), ...newPlan.value })
  newPlan.value = { name: '', type: '', price: '', description: '' }
  showAddPlan.value = false
}

onMounted(() => {
  fetchUsers()
  fetchSupport()
})
</script>