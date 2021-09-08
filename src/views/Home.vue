<template>
  <div class="home">
    <div class="mt-12 lg:mt-16 mb-24 flex flex-col items-center justify-center">
      <div class="text-3xl font-bold text-white mb-10">PHOTOGRAPHY</div>
      <div class="text-2xl text-white">ABSTRACT</div>
      <div class="flex flex-row sm:flex-col">
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-4 sm:mt-14 mb-8 sm:mb-2 transition-transform transform hover-d:hover:scale-105">
          <film-strip img="DSF5351.jpg" alt="Fall sky in Georgia" />
          <film-strip img="XT37921.jpg" alt="Lamar Building, Augusta, GA"/>
          <film-strip img="XT30374.jpg" alt="Mother and baby ducks" />
          <film-strip img="XT30469.jpg" alt="Boarded up drive-thru bank" />
          <film-strip img="DSF5336.jpg" alt="Monochrome trees" />
        </div>
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-4 sm:mt-4 mb-8 sm:mb-14 transition-transform transform hover-d:hover:scale-105">
          <film-strip img="XT30749.jpg" alt="Fall sky in Georgia"/>
          <film-strip img="XT39366.jpg" alt="Lamar Building, Augusta, GA" />
          <film-strip img="MTH7650.jpg" alt="Mother and baby ducks" :portrait="true" />
          <film-strip img="XT38677.jpg" alt="Boarded up drive-thru bank" />
          <film-strip img="XT39735.jpg" alt="Monochrome trees" />
        </div>
      </div>
      <div class="text-2xl text-white">ARCHITECTURE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="DSF3482.jpg" alt="Barely lit pillow on bed" :portrait="true" />
        <film-strip img="XT30929.jpg" alt="Close up plants"/>
        <film-strip img="XT39591.jpg" alt="Light falling across stucco building" :portrait="true" />
        <film-strip img="XT39569.jpg" alt="Buildings and light strip from below" />
        <film-strip img="DSF4762.jpg" alt="Flowers on branch" />
      </div>
      <div class="text-2xl text-white">PEOPLE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="XT39772.jpg" alt="Barely lit pillow on bed" />
        <film-strip img="XT30836.jpg" alt="Close up plants"/>
        <film-strip img="XT37993.jpg" alt="Light falling across stucco building" :portrait="true" />
        <film-strip img="XT39684.jpg" alt="Buildings and light strip from below" />
        <film-strip img="XT39701.jpg" alt="Flowers on branch" />
      </div>
      <div class="text-2xl text-white">B/W NATURE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="MTH4273.jpg" alt="Barely lit pillow on bed" />
        <film-strip img="MTH4442.jpg" alt="Close up plants"/>
        <film-strip img="XT30148.jpg" alt="Light falling across stucco building"/>
        <film-strip img="XT39897.jpg" alt="Buildings and light strip from below" />
        <film-strip img="XT39887.jpg" alt="Flowers on branch" />
      </div>
    </div>
    <div class="my-24 flex flex-col items-center justify-center">
      <div class="text-3xl font-bold text-white">PERSONAL DEV PROJECTS</div>
      <!-- <code-block app-name="sysmon-web-manager (Sysmon remote management)" url="https://github.com/rwwiv/sysmon-web-manager" language="python" :code="sysmonCode" /> -->
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
