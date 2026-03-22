import { onMounted, onUnmounted } from 'vue'

export function useScrollAnimation() {
  let observer = null

  function observeElements() {
    document.querySelectorAll('.fade-in:not(.observed)').forEach(el => {
      el.classList.add('observed')
      observer.observe(el)
    })
  }

  onMounted(() => {
    observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible')
          observer.unobserve(entry.target)
        }
      })
    }, { threshold: 0.15 })

    // Observa imediat
    observeElements()

    // Re-observa dupa ce datele din API se incarca
    setTimeout(observeElements, 500)
    setTimeout(observeElements, 1000)
  })

  onUnmounted(() => {
    if (observer) observer.disconnect()
  })
}