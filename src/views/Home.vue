<template>
  <div class="home">
    <div class="mt-12 lg:mt-16 mb-24 flex flex-col items-center justify-center">
      <div class="text-3xl font-bold text-white mb-10">PHOTOGRAPHY</div>
      <div class="text-2xl text-white">ABSTRACT</div>
      <div class="flex flex-row sm:flex-col">
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-4 sm:mt-14 mb-8 sm:mb-2 transition-transform transform hover-d:hover:scale-105">
          <film-strip img="DSF5351.jpg" alt="Sky with light pole" />
          <film-strip img="XT37921.jpg" alt="Turtle"/>
          <film-strip img="XT30374.jpg" alt="Chandelier with recessed lights" />
          <film-strip img="XT30469.jpg" alt="Unfinished wall with pipes" />
          <film-strip img="DSF5336.jpg" alt="Barely lit pillow on bed" />
        </div>
        <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 mt-4 sm:mt-4 mb-8 sm:mb-14 transition-transform transform hover-d:hover:scale-105">
          <film-strip img="XT30749.jpg" alt="Parking garage exterior stairwell"/>
          <film-strip img="XT39366.jpg" alt="Man on stairwell through window with tree leaves" />
          <film-strip img="MTH7650.jpg" alt="Power cable hung on wall" />
          <film-strip img="XT38677.jpg" alt="Pipes under bridge and blue pillar edge" />
          <film-strip img="XT39735.jpg" alt="Building reflected in skinny windows" />
        </div>
      </div>
      <div class="text-2xl text-white">ARCHITECTURE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="DSF3482.jpg" alt="Lamar building facade" />
        <film-strip img="XT30929.jpg" alt="Brutalist building reflecting other building"/>
        <film-strip img="XT39591.jpg" alt="Brutalist skyscraper" />
        <film-strip img="XT39569.jpg" alt="Concrete walkway from below in rain" />
        <film-strip img="DSF4762.jpg" alt="Lights strung from buildings seen from below" />
      </div>
      <div class="text-2xl text-white">PEOPLE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="XT39772.jpg" alt="Well dressed family at bus stop" />
        <film-strip img="XT30836.jpg" alt="Shirtless man walking right"/>
        <film-strip img="XT37993.jpg" alt="Skateboarder in air from behind" />
        <film-strip img="XT39684.jpg" alt="Businessman checking phone, seen through window" />
        <film-strip img="XT39701.jpg" alt="Woman with red hair and colorful clothes walking" />
      </div>
      <div class="text-2xl text-white">B/W NATURE</div>
      <div class="flex flex-col sm:flex-row items-center justify-center content-center mx-2 sm:mx-0 my-8 sm:my-14 transition-transform transform hover-d:hover:scale-105">
        <film-strip img="MTH4273.jpg" alt="Mother and baby ducks - black and white" />
        <film-strip img="MTH4442.jpg" alt="Tree branches and leaves - black and white"/>
        <film-strip img="XT30148.jpg" alt="Waterfall within forest - black and white"/>
        <film-strip img="XT39897.jpg" alt="Rolling forested mountains in fog - black and white" />
        <film-strip img="XT39887.jpg" alt="Dead tree against forested mountains in fog - black and white" />
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
