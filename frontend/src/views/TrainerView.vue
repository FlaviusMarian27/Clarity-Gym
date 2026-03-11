<template>
  <div>
    <NavbarTrainer />

    <!-- HERO -->
    <section id="hero" class="relative min-h-screen flex items-center justify-center overflow-hidden">
      <video
        ref="heroVideo"
        autoplay
        muted
        playsinline
        class="absolute inset-0 w-full h-full object-cover"
        @ended="switchVideo"
      >
        <source :src="currentVideo" type="video/mp4" />
      </video>

      <div class="absolute inset-0 bg-text opacity-50"></div>

      <div class="relative z-10 text-center px-6">
        <h1 class="text-6xl font-bold text-white mb-4">Clarity Gym</h1>
        <p class="text-2xl text-primary mb-4">Premium Wellness & Fitness</p>
        <p class="text-lg text-white opacity-80 mb-10 max-w-xl mx-auto">
          Descoperă un loc unde performanța întâlnește eleganța. Antrenori dedicați, echipamente moderne și o comunitate care te inspiră.
        </p>
      </div>
    </section>

    <!-- DESPRE -->
    <section id="despre" class="py-24 bg-card">
      <div class="max-w-6xl mx-auto px-8">

        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-text mb-4">Despre Noi</h2>
          <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
        </div>

        <div class="flex gap-10 items-center mb-16">
          <div class="w-1/2">
            <img src="/sala1.png" alt="Clarity Gym" class="rounded-2xl w-full h-80 object-cover shadow-sm" />
          </div>
          <div class="w-1/2 flex flex-col gap-4">
            <p class="text-lg text-text leading-relaxed opacity-80">
              Clarity Gym este un concept de fitness de tip boutique wellness, care redefinește experiența sălii de forță prin transparență și echilibru.
            </p>
            <p class="text-lg text-text leading-relaxed opacity-80">
              Într-un spațiu modern, inundat de lumină naturală și definit de un design minimalist în nuanțe calde de bej și salvie, Clarity Gym oferă o alternativă rafinată sălilor clasice. Accentul nu cade doar pe efortul fizic, ci pe claritatea parcursului tău: antrenamente personalizate, o echipă de profesioniști dedicată și o atmosferă premium care inspiră disciplină și calm.
            </p>
            <p class="text-lg text-text leading-relaxed opacity-80">
              Pe scurt, este locul unde sănătatea devine un stil de viață organizat, simplu și eficient.
            </p>
          </div>
        </div>

        <div class="flex gap-10 items-center">
          <div class="w-1/2 flex flex-col gap-6">
            <div>
              <h3 class="text-xl font-bold text-text mb-2">Echipamente Moderne</h3>
              <p class="text-text opacity-70">Dispunem de echipamente de ultimă generație pentru un antrenament eficient și sigur.</p>
            </div>
            <div>
              <h3 class="text-xl font-bold text-text mb-2">Antrenori Certificați</h3>
              <p class="text-text opacity-70">Echipa noastră de antrenori profesioniști te ghidează pas cu pas spre obiectivele tale.</p>
            </div>
            <div>
              <h3 class="text-xl font-bold text-text mb-2">Experiență Premium</h3>
              <p class="text-text opacity-70">Un mediu elegant și motivant unde fiecare detaliu este gândit pentru confortul tău.</p>
            </div>
          </div>
          <div class="w-1/2">
            <img src="/sala2.png" alt="Clarity Gym" class="rounded-2xl w-full h-80 object-cover shadow-sm" />
          </div>
        </div>

      </div>
    </section>

    <!-- RECENZII -->
    <section id="recenzii" class="py-24 bg-bg">
      <div class="max-w-6xl mx-auto px-8">

        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-text mb-4">Ce Spun Membrii Noștri</h2>
          <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
        </div>

        <div class="grid grid-cols-3 gap-8">
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-primary flex items-center justify-center text-white font-bold text-lg">A</div>
              <div>
                <h4 class="font-bold text-text">Alexandru M.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"Cea mai bună sală la care am fost. Atmosfera e incredibilă, echipamentele sunt mereu curate și antrenorii sunt super profesioniști."</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-secondary flex items-center justify-center text-white font-bold text-lg">M</div>
              <div>
                <h4 class="font-bold text-text">Maria P.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"Am slăbit 12 kg în 3 luni cu ajutorul antrenorului meu personal. Planul VIP merită fiecare leu investit!"</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-primary flex items-center justify-center text-white font-bold text-lg">R</div>
              <div>
                <h4 class="font-bold text-text">Radu T.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"Abonamentul Standard e perfect pentru programul meu. Acces oricând, prețuri corecte și o comunitate extraordinară."</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-secondary flex items-center justify-center text-white font-bold text-lg">I</div>
              <div>
                <h4 class="font-bold text-text">Ioana C.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"Spațiul e mereu curat, luminos și plăcut. Mă simt motivată de fiecare dată când intru pe ușă."</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-primary flex items-center justify-center text-white font-bold text-lg">D</div>
              <div>
                <h4 class="font-bold text-text">Daniel F.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"Am recomandat Clarity Gym tuturor prietenilor mei. Raportul calitate-preț e imbatabil în oraș."</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
          <div class="bg-card rounded-2xl p-6 shadow-sm">
            <div class="flex items-center gap-4 mb-4">
              <div class="w-12 h-12 rounded-full bg-secondary flex items-center justify-center text-white font-bold text-lg">A</div>
              <div>
                <h4 class="font-bold text-text">Andrada N.</h4>
                <div class="flex gap-1 text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-text opacity-70 text-sm leading-relaxed">"De când am început abonamentul de dimineață, ziua mea e complet diferită. Energie, focus și rezultate vizibile!"</p>
            <p class="text-primary font-semibold text-sm mt-3">5.0 / 5</p>
          </div>
        </div>

      </div>
    </section>

    <!-- LOCATIE -->
    <section id="locatie" class="py-24 bg-card">
      <div class="max-w-6xl mx-auto px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-text mb-4">Locație</h2>
          <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
          <p class="text-lg text-text opacity-70 mt-6">Ne găsești în inima Timișoarei</p>
        </div>
        <div class="rounded-2xl overflow-hidden shadow-sm">
          <iframe
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d88280.10983570529!2d21.1782!3d45.7489!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47455d84610655bf%3A0xfd169ff07d1c4c78!2sTimișoara!5e0!3m2!1sro!2sro!4v1"
            width="100%"
            height="450"
            style="border:0;"
            allowfullscreen=""
            loading="lazy"
          ></iframe>
        </div>
      </div>
    </section>

    <!-- SUPORT -->
    <section id="suport" class="py-24 bg-bg">
      <div class="max-w-6xl mx-auto px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-text mb-4">Suport Tehnic</h2>
          <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
          <p class="text-lg text-text opacity-70 mt-6">Ai o problemă? Suntem aici să te ajutăm.</p>
        </div>
        <div class="max-w-2xl mx-auto">
          <div class="flex flex-col gap-4">
            <input type="text" placeholder="Numele tău" class="border border-primary rounded-xl px-4 py-3 text-text bg-card focus:outline-none focus:ring-2 focus:ring-secondary" />
            <input type="email" placeholder="Email" class="border border-primary rounded-xl px-4 py-3 text-text bg-card focus:outline-none focus:ring-2 focus:ring-secondary" />
            <select class="border border-primary rounded-xl px-4 py-3 text-text bg-card focus:outline-none focus:ring-2 focus:ring-secondary">
              <option value="">Selectează categoria problemei</option>
              <option value="cont">Problemă cont</option>
              <option value="abonament">Problemă abonament</option>
              <option value="tehnic">Problemă tehnică</option>
              <option value="altele">Altele</option>
            </select>
            <textarea placeholder="Descrie problema ta..." rows="5" class="border border-primary rounded-xl px-4 py-3 text-text bg-card focus:outline-none focus:ring-2 focus:ring-secondary resize-none"></textarea>
            <button class="bg-primary hover:bg-secondary text-white font-semibold py-3 rounded-xl transition-colors duration-200">Trimite</button>
          </div>
        </div>
      </div>
    </section>

    <!-- CONTACT -->
    <section id="contact" class="py-24 bg-card">
      <div class="max-w-6xl mx-auto px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl font-bold text-text mb-4">Contact</h2>
          <div class="w-20 h-1 bg-primary mx-auto rounded-full"></div>
          <p class="text-lg text-text opacity-70 mt-6">Urmărește-ne pe rețelele sociale</p>
        </div>
        <div class="flex justify-center gap-8">
          <a href="https://facebook.com" target="_blank" class="flex flex-col items-center gap-3 p-6 bg-bg rounded-2xl shadow-sm hover:shadow-md transition-all duration-200 w-36">
            <svg class="w-10 h-10 text-blue-600" fill="currentColor" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
            <span class="text-text font-medium text-sm">Facebook</span>
          </a>
          <a href="https://instagram.com" target="_blank" class="flex flex-col items-center gap-3 p-6 bg-bg rounded-2xl shadow-sm hover:shadow-md transition-all duration-200 w-36">
            <svg class="w-10 h-10 text-pink-500" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
            <span class="text-text font-medium text-sm">Instagram</span>
          </a>
          <a href="https://tiktok.com" target="_blank" class="flex flex-col items-center gap-3 p-6 bg-bg rounded-2xl shadow-sm hover:shadow-md transition-all duration-200 w-36">
            <svg class="w-10 h-10 text-text" fill="currentColor" viewBox="0 0 24 24"><path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.09 1.75 4.17 1.12 1.11 2.7 1.62 4.24 1.79v4.03c-1.44-.05-2.89-.35-4.2-.97-.57-.26-1.1-.59-1.62-.93-.01 2.92.01 5.84-.02 8.75-.08 1.4-.54 2.79-1.35 3.94-1.31 1.92-3.58 3.17-5.91 3.21-1.43.08-2.86-.31-4.08-1.03-2.02-1.19-3.44-3.37-3.65-5.71-.02-.5-.03-1-.01-1.49.18-1.9 1.12-3.72 2.58-4.96 1.66-1.44 3.98-2.13 6.15-1.72.02 1.48-.04 2.96-.04 4.44-.99-.32-2.15-.23-3.02.37-.63.41-1.11 1.04-1.36 1.75-.21.51-.15 1.07-.14 1.61.24 1.64 1.82 3.02 3.5 2.87 1.12-.01 2.19-.66 2.77-1.61.19-.33.4-.67.41-1.06.1-1.79.06-3.57.07-5.36.01-4.03-.01-8.05.02-12.07z"/></svg>
            <span class="text-text font-medium text-sm">TikTok</span>
          </a>
          <a href="https://wa.me/40700000000" target="_blank" class="flex flex-col items-center gap-3 p-6 bg-bg rounded-2xl shadow-sm hover:shadow-md transition-all duration-200 w-36">
            <svg class="w-10 h-10 text-green-500" fill="currentColor" viewBox="0 0 24 24"><path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/></svg>
            <span class="text-text font-medium text-sm">WhatsApp</span>
          </a>
          <a href="mailto:contact@claritygym.ro" class="flex flex-col items-center gap-3 p-6 bg-bg rounded-2xl shadow-sm hover:shadow-md transition-all duration-200 w-36">
            <svg class="w-10 h-10 text-primary" fill="currentColor" viewBox="0 0 24 24"><path d="M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 4l-8 5-8-5V6l8 5 8-5v2z"/></svg>
            <span class="text-text font-medium text-sm">Email</span>
          </a>
        </div>
      </div>
    </section>

    <!-- FOOTER -->
    <footer class="bg-text py-6">
      <div class="text-center">
        <p class="text-white opacity-60 text-sm">© {{ currentYear }} Clarity Gym. Toate drepturile rezervate.</p>
      </div>
    </footer>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import NavbarTrainer from '../components/NavbarTrainer.vue'

const videos = ['/video1.mp4', '/video2.mp4']
const currentIndex = ref(0)
const currentVideo = ref(videos[0])
const heroVideo = ref(null)
const currentYear = new Date().getFullYear()

function switchVideo() {
  currentIndex.value = (currentIndex.value + 1) % videos.length
  currentVideo.value = videos[currentIndex.value]
  heroVideo.value.load()
  heroVideo.value.play()
}
</script>