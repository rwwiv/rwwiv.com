<template>
  <floating-modal @close="$emit('close')">
    <div :style="imgWidthCss" class="img-wrapper bg-white rounded-sm p-3">
      <v-lazy-image 
        :src="importFullImg()" 
        :src-placeholder="importThumb()" 
        :alt="alt" 
        @intersect="loading = true"
        @load="loading = false"
      />
      <transition name="overlay">
        <div class="overlay-fix w-full h-full p-3" v-if="loading">
        <div class="w-full h-full overlay">
          <div class="centered-text text-black text-6xl">
            ...
          </div>
        </div>
      </div>
      </transition>
    </div>
  </floating-modal>
</template>

<script>
import { importImage } from '../helpers/img'

import FloatingModal from './FloatingModal.vue'
import VLazyImage from "v-lazy-image";

export default {
  components: {
    FloatingModal,
    VLazyImage
  },
  props: {
    img: { type: String, default: ''},
    alt: { type: String, default: ''},
    isPortrait: { type: Boolean, default: false }
  },
  data: () => ({
    loading: true
  }),
  computed: {
    imgWidthCss() {
      return this.isPortrait ? '--img-width: 50rem;' : '--img-width: 60rem;'
    }
  },
  methods: {
    importFullImg() {
      return importImage('full', this.img)
    },
    importThumb() {
      return importImage('thumbs', this.img)
    }
  }

}
</script>

<style scoped>
.img-wrapper {
  width: 30rem;
  position: relative;
  -webkit-touch-callout: none; /* iOS Safari */
  -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
      -moz-user-select: none; /* Old versions of Firefox */
      -ms-user-select: none; /* Internet Explorer/Edge */
          user-select: none; /* Non-prefixed version, currently
                                supported by Chrome, Edge, Opera and Firefox */
}
@media (min-width: 480px) {
  .img-wrapper {
    width: var(--img-width);
  }
}

.v-lazy-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.overlay-fix{
  position: absolute;
  top: 0;
  left: 0;
}

.overlay {
  opacity: 0.45;
  background-color: white;
}

.centered-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
}

.overlay-enter-active, .overlay-leave-active {
  transition: opacity .5s ease;
}
.overlay-enter, .overlay-leave-to {
  opacity: 0;
}
</style>
