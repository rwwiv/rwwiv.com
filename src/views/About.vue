@@ -1,39 +0,0 @@
<template>
  <div class="flex flex-col items-center justify-center">
    <div class="flex flex-col items-center justify-center my-10 bg-white rounded-sm img-wrapper">
      <v-lazy-image :src="require('@/assets/img/full/DSF4898.jpg')" :src-placeholder="require('@/assets/img/thumbs/DSF4898.jpg')" class="mt-4"/>
      <div class="my-6 sm:my-8">Self Portrait - me, 2020</div>
    </div>
    <div class="text-white text-center w-5/6">
      Open source software developer for <a href="https://securityonionsolutions.com/" target="_blank" class="underline">Security Onion Solutions</a>
      <br>
      &bull;
      <br>
      BS in Computer Science from Augusta University
      <br>
      &bull;
      <br>
      Amateur photographer (Canon A-1, Fujifilm X-T3)
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
  width: 95%;
  height: 100%;
  object-fit: cover;
}

.img-wrapper {
  width: 75%;
  max-width: 20em;
}

@media (min-width: 640px) {
  .img-wrapper {
    width: 75%;
    max-width: 30em;
  }
}
</style>
