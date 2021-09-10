<template>
 <div class="film-bg">
    <div class="thumbnail my-1 sm:my-4 mx-4 sm:mx-1 cursor-pointer" @click="handleClick()">
      <img :src="importImg()" :alt="alt" :class="[ portrait ? 'portrait' : '' ]" />
    </div>
 </div>
</template>

<script>
import { importImage } from '../helpers/img'
import { EventBus } from '../main'

export default {
  data: () => ({
    portrait: false
  }),
  props: {
    img: { type: String, default: '' },
    alt: { type: String, default: '' },
  },
  methods: {
    importImg() {
      return importImage('thumbs', this.img)
    },
    handleClick() {
      EventBus.$emit('show-img-modal', this.img, this.alt, this.portrait)
    }
  },
  async beforeMount () {
    const img = new Image()
    img.onload = () => {
      let portrait = false
      const ratio = img.width / img.height
      portrait = ratio < 0.9 ? true : false
      this.portrait = portrait
    }
    img.src = await importImage('thumbs', this.img)
    await img.onload()
  }
}
</script>

<style scoped lang="postcss">

.film-bg {
  background-color: #e27745ad;
}

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
  filter: invert(1) brightness(85%) hue-rotate(330deg);
}

img.portrait {
  width: 100%;
  height: auto;
}
</style>
