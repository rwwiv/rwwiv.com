@@ -1,39 +0,0 @@
<template>
  <div class="flex flex-col items-center justify-center">
    <div class="flex flex-col items-center justify-center my-10 bg-white rounded-sm img-wrapper">
      <v-lazy-image :src="require('@/assets/img/full/DSF4898.jpg')" :src-placeholder="require('@/assets/img/thumbs/DSF4898.jpg')" class="mt-4 px-3 h-full" height="250" width="250" />
      <div class="my-7 sm:my-10 font-bold sm:text-xl">SELF PORTRAIT - ME, 2020</div>
    </div>
    <div class="text-white text-center w-5/6 mb-12">
      Open source software developer for <a href="https://securityonionsolutions.com/" target="_blank" class="underline">Security Onion Solutions</a>
      <br>
      &bull;
      <br>
      BS in Computer Science from Augusta University
      <br>
      &bull;
      <br>
      Amateur photographer
      <br>
      (Canon A-1, Fujifilm X-T3)
    </div>
  </div>
</template>

<script>
import VLazyImage from "v-lazy-image";

export default {
  components: {
    VLazyImage
  }
}
</script>

<style scoped>

.v-lazy-image {
  width: 20em;
  min-height: auto;
  object-fit: cover;
}


@media (min-width: 640px) {
  .v-lazy-image {
    width: 30em;
  }
}
</style>
