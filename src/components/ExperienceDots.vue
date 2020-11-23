<template>
  <div class="flex flex-row">
    <div class="mr-2 w-1/2 xxs:w-1/4 sm:w-1/3 text-right">
      <slot />
    </div>
   <div class="flex flex-row w-1/2 xxs:w-1/4 sm:w-2/3 font-mono	">
      <div v-for="index in num" :key="index" class="font-extrabold">
        &bull;
      </div>
      <div v-for="index in remainder" :key="index" class="text-gray-400">
        -
      </div>
   </div>
  </div>
</template>

<script>
export default {
  props: {
    num: { type: Number, default: 1 }
  },
  computed: {
    remainder() {
      return 10 - this.num
    }
  }
}
</script>

<style>

</style>