<template>
  <div class="flex flex-col text-left">
    <div class="font-semibold text-xl">{{ jobTitle }}</div>
    <div>
      {{ employerName }}
      -
      {{ location }}
    </div>
    <div class="font-light text-gray-300">
      {{ startDate }} - {{ endDate }}
    </div>
    <slot name="experience" />
  </div>
</template>

<script>
export default {
  props: {
    jobTitle: { type: String, default: ''},
    employerName: { type: String, default: '' },
    location: { type: String, default: '' },
    startDate: { type: String, default: '' },
    endDate: { type: String, default: '' }
  }
}
</script>

<style>

</style>