<template>
  <div id="app" >
    <div id="wrapper" tabindex="0" @keydown.esc="showModal = false;" class="outline-none">
      <transition name="modal">
        <full-picture 
          v-if="showModal" 
          :index="1"
          @close="showModal = false"
          :img="fullImg"
          :alt="altTag"
          :is-portrait="isPortrait"
        />
      </transition>
    
      <div id="content" class="z-50">
        <div id="nav">
          <router-link to="/" class="text-5xl mt-3">RWWIV</router-link>
          <div class="flex justify-center space-x-4 mt-1 text-lg">
            <router-link to="/about">About</router-link>
            <router-link to="/resume">Resume</router-link>
          </div>
          <div class="flex justify-center space-x-4 mt-4">
            <a href="https://www.instagram.com/rwwiv/" target="_blank">Instagram</a>
            <a href="https://github.com/rwwiv" target="_blank">GitHub</a>
            <a href="https://www.linkedin.com/in/rwwiv/" target="_blank">LinkedIn</a>
            <a href="https://www.upwork.com/o/profiles/users/~01be2089b2b24074bc/" target="_blank">UpWork</a>
          </div>
        </div>
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from './main'

import FullPicture from './components/FullPicture.vue'

export default {
  components: { FullPicture },
  data: () => ({
    showModal: false,
    fullImg: '',
    altTag: '',
    isPortrait: false,
    colorSchemeString: ''
  }),
  beforeMount() {
    EventBus.$on('show-img-modal', (img, alt, portrait) => {
      this.showModal = true
      this.fullImg = img
      this.altTag = alt
      this.isPortrait = portrait
    })
  },
  mounted () {
    this.colorSchemeString = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
      ? '-dark'
      : ''

    const darkMediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

    try {
      darkMediaQuery.addEventListener('change', () => this.setFavicon())
    } catch (e) {
      try {
        darkMediaQuery.addListener(() => this.setFavicon())
      } catch (e) {}
    }
    this.setFavicon()
  },
  methods: {
    setFavicon () {
      window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? this.colorSchemeString = '-dark' : this.colorSchemeString = ''
      const fav = document.getElementById('favicon')
      fav.href = `/favicon/favicon${this.colorSchemeString}.png`
    }
  },
}
</script>

<style lang="postcss">

html {
  background: black;
}

#nav {
  @apply flex flex-col items-center w-full pt-4;
}

#nav a {
  font-weight: bold;
  color: grey;
}

#nav a.router-link-exact-active {
  color: white;
}

.modal-enter-active, .modal-leave-active {
  transition: opacity .15s ease;
}
.modal-enter, .modal-leave-to {
  opacity: 0;
}
</style>
