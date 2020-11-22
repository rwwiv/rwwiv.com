<template>
  <div class="home">
    <div class="mt-12 lg:mt-16 mb-24 flex flex-col items-center justify-center">
      <div class="text-3xl font-bold text-white">PHOTOGRAPHY</div>
      <div class="flex flex-row sm:flex-col ">
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-8 sm:mt-14 transition-transform transform hover:scale-105">
          <film-strip img="MTH4703.jpg" alt="Fall sky in Georgia" />
          <film-strip img="DSF3482.jpg" alt="Lamar Building, Augusta, GA" :portrait="true" />
          <film-strip img="MTH4273.jpg" alt="Mother and baby ducks" />
          <film-strip img="DSF3532.jpg" alt="Boarded up drive-thru bank" />
          <film-strip img="MTH4442.jpg" alt="Monochrome trees" />
        </div>
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-8 transition-transform transform hover:scale-105">
          <film-strip img="DSF5336.jpg" alt="Barely lit pillow on bed" />
          <film-strip img="DSF5438.jpg" alt="Close up plants"/>
          <film-strip img="DSF3448.jpg" alt="Light falling across stucco building" />
          <film-strip img="DSF4762.jpg" alt="Buildings and light strip from below" />
          <film-strip img="MG0855.jpg" alt="Flowers on branch" />
        </div>
      </div>
    </div>
    <div class="my-24 flex flex-col items-center justify-center">
      <div class="text-3xl font-bold text-white">PERSONAL DEV PROJECTS</div>
      <code-block app-name="sysmon-web-manager (Sysmon remote management)" url="https://github.com/rwwiv/sysmon-web-manager" language="python" :code="sysmonCode" />
      <code-block app-name="turma (Contrived Turing machine language compiler)" url="https://github.com/rwwiv/turma" language="kotlin" :code="turmaCode" />

    </div>
  </div>
</template>

<script>
import CodeBlock from '../components/CodeBlock.vue';
import FilmStrip from '../components/FilmStrip.vue';

import sysmon from '../assets/code/sysmon-agent.txt'
import turma from '../assets/code/turma.txt'

export default {
  name: 'Home',
  components: {
    FilmStrip,
    CodeBlock
  },
  data: () => ({
    sysmonCode: sysmon,
    turmaCode: turma
  })
};
</script>
