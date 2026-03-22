<template>
  <Transition name="fade">
    <div v-if="visible" class="fixed inset-0 z-[999] bg-text flex flex-col items-center justify-center">
      <img src="/logo.png" class="w-24 h-24 object-contain animate-bounce mb-6" />
      <h1 class="text-4xl font-bold text-white mb-2">Clarity Gym</h1>
      <p class="text-primary text-lg">Premium Wellness & Fitness</p>
      <div class="mt-8 flex gap-2">
        <div class="w-2 h-2 bg-primary rounded-full animate-bounce" style="animation-delay: 0s"></div>
        <div class="w-2 h-2 bg-primary rounded-full animate-bounce" style="animation-delay: 0.2s"></div>
        <div class="w-2 h-2 bg-primary rounded-full animate-bounce" style="animation-delay: 0.4s"></div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const visible = ref(true)

onMounted(() => {
  setTimeout(() => {
    visible.value = false
  }, 2000)
})
</script>

<style scoped>
.fade-leave-active {
  transition: opacity 0.8s ease;
}
.fade-leave-to {
  opacity: 0;
}
</style>