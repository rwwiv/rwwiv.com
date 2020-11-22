<template>
  <img :src="importImage()" :alt="alt"/>
</template>

<script>
export default {
  props: {
    src: { type: String, default: '' },
    alt: { type: String, default: '' }
  },
  methods: {
    importImage() {
      return require(`../assets/img/thumbs/${this.src}`)
    }
  }
}
</script>

<style>
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