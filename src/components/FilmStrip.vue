<template>
 <div class="bg-gray-900">
    <div class="thumbnail my-3 sm:my-6 mx-4 sm:mx-2 cursor-pointer" @click="handleClick()">
      <img :src="importImg()" :alt="alt" :class="[ portrait ? 'portrait' : '' ]" />
    </div>
 </div>
</template>

<script>
import { importImage } from '../helpers/img'
import { EventBus } from '../main'

export default {
  props: {
    img: { type: String, default: '' },
    alt: { type: String, default: '' },
    portrait: { type: Boolean, default: false }
  },
  methods: {
    importImg() {
      return importImage('thumbs', this.img)
    },
    handleClick() {
      EventBus.$emit('show-img-modal', this.img, this.alt, this.portrait)
    }
  }
}
</script>

<style scoped lang="postcss">

.thumbnail {
  position: relative;
  width: 8rem;
  min-width: 8rem;
  height: 8rem;
  overflow: hidden;
}

@media (min-width: 640px) {
  .thumbnail {
  position: relative;
  width: 6rem;
  min-width: 6rem;
  height: 6rem;
  overflow: hidden;
}
}

@media (min-width: 768px) {
  .thumbnail {
    position: relative;
    width: 8rem;
    min-width: 8rem;
    height: 8rem;
    overflow: hidden;
  }
}

@media (min-width: 1280px) {
  .thumbnail {
    position: relative;
    width: 12rem;
    min-width: 12rem;
    height: 12rem;
    overflow: hidden;
  }
}

img {
  position: absolute;
  left: 50%;
  top: 50%;
  height: 100%;
  width: auto;
  -webkit-transform: translate(-50%,-50%);
      -ms-transform: translate(-50%,-50%);
          transform: translate(-50%,-50%);
  object-fit: cover;
  filter: invert(1)
}
img.portrait {
  width: 100%;
  height: auto;
}
</style>