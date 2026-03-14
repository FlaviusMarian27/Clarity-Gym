import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { 
      path: '/', 
      name: 'home', 
      component: () => import('../views/HomeView.vue') 
    },

    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },

    { 
      path: '/register', 
      name: 'register', 
      component: () => import('../views/RegisterView.vue') 
    },

    {
      path: '/client',
      name: 'client',
      component: () => import('../views/ClientView.vue'),
      meta: { requiresAuth: true, role: 'client' }
    },

    {
      path: '/trainer',
      name: 'trainer',
      component: () => import('../views/TrainerView.vue'),
      meta: { requiresAuth: true, role: 'trainer' }
    },
    
    {
      path: '/admin',
      name: 'admin',
      component: () => import('../views/AdminView.vue'),
      meta: { requiresAuth: true, role: 'admin' }
    },

    {
      path: '/profile',
      component: () => import('../views/ProfileView.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth) {
    if (!authStore.isLoggedIn) {
      next('/login')
      return
    }
    if (to.meta.role && authStore.role !== to.meta.role) {
      if (authStore.role === 'client') next('/client')
      else if (authStore.role === 'trainer') next('/trainer')
      else if (authStore.role === 'admin') next('/admin')
      return
    }
  }

  next()
})

export default router