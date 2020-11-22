<template>
  <div class="flex flex-col md:flex-row content-center md:items-start text-left md:justify-left space-y-2 md:space-y-0 md:space-x-5 w-10/12 lg:w-1/2 my-6">
      <div class="w-full md:w-1/6 mb-4 text-3xl md:text-lg font-semibold">
        <slot name="title" />
      </div>
      <div class="w-full md:w-5/6">
        <slot name="content" />
      </div>
    </div>
</template>

<script>
export default {

}
</script>

<style>

</style>