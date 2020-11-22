<template>
  <floating-modal @close="$emit('close')">
    <div :style="imgWidthCss" class="img-wrapper bg-white rounded-sm p-5">
      <v-lazy-image :src="importFullImg()" :src-placeholder="importThumb()" :alt="alt" />
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

</style>
