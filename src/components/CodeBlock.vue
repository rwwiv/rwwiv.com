<template>
  <a class="flex flex-col justify-center mt-12 cursor-pointer transition-transform transform hover-d:hover:scale-105 w-10/12 sm:w-3/4 lg:w-2/3 xl:w-1/2 3xl:w-1/3" :href="url" target="_blank">
    <div class="text-2xl text-gray-200 text-center">{{ appName }}</div>
    <div class="w-full">
      <vue-code-highlight :language="language">
        <pre>
{{ code }}
        </pre>
      </vue-code-highlight>
    </div>
  </a>
</template>

<script>
import { component as VueCodeHighlight } from 'vue-code-highlight';
import "vue-code-highlight/themes/duotone-sea.css";
import 'prism-es6/components/prism-kotlin';
import 'prism-es6/components/prism-python';

export default {
  components:{
    VueCodeHighlight,
  },
  props: {
    appName: { type: String, default: '' },
    url: { type: String, default: '' },
    language: { type: String, default: '' },
    code: { type: String, default: '' }
  },
  methods: {
    handleClick() {
      window.location.href = this.url
    }
  }
}
</script>

<style>

</style>