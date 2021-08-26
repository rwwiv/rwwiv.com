<template>
  <div id="start" class="flex flex-col text-white items-center">
    <div class="text-center mt-12 mb-6">
      <div class="text-4xl">
        William Wernert
      </div>
      <div class="text-lg font-light">
        Full Stack Developer
      </div>
    </div>
    <resume-block>
      <template slot="title">
        Summary
      </template>
      <template slot="content">
        I am constantly challenging myself to improve my development skills and explore new languages and frameworks.
        My current job has me primarily working in JavaScript on the frontend and Python on the backend.
        I am most familiar with Kotlin and Node.js for server-side applications and the Vue framework for client-side apps;
        but have no trouble working in whatever language or framework is thrown at me.
      </template>
    </resume-block>
    <resume-block>
      <template slot="title">
        Work Experience
      </template>
      <template slot="content">
        <employment-block
          job-title="Software Developer"
          employer-name="Security Onion Solutions, LLC"
          location="Augusta, GA"
          start-date="June 2019"
          end-date="Present"
          class="mb-12"
        >
          <template slot="experience">
            <div class="mt-4">
                Worked as part of a team to bring a new, container-based version of the product to release. As part of this:              <ul class="list-disc">
                <li class="ml-5">
                  Introduced automated E2E testing to the project via the Cypress testing framework and pioneered its use among the team.
                </li>
                <li class="ml-5">
                  Added better error handling to and refactored several Python/Flask applications and shell components to improve stability of the overall product.
                </li>
                <li class="ml-5">
                  Worked with the team to add new functionality to and improve extensibility of these apps and several other components, as well as the product’s main Go application.
                </li>
              </ul>
            </div>
            <div class="mt-4">
              Worked with founding members of the company on a product and company brand update. As part of this:
              <ul class="list-disc">
                <li class="ml-5">
                  Designed and built the company’s new website using current web frameworks, namely Vue/Nuxt.
                </li>
              </ul>
            </div>
            <div class="mt-4">
              Introduced a beta machine learning tool to the product and created a ML framework for additional applications to utilize in the future. As part of this:
              <ul class="list-disc">
                <li class="ml-5">
                  Created tooling in Python to train models using supervised learning in a repeatable and automated manner. This tooling was designed to be extensible for the addition of future models as needed.
                </li>
                <li class="ml-5">
                  Wrote a Python application that used the trained models to predict anomalous activity in the product’s authorization logs.
                </li>
                <li class="ml-5">
                  Ensured all code was adequately unit tested, and where appropriate, E2E tested.
                </li>
              </ul>
            </div>
          </template>
        </employment-block>
        <employment-block
          job-title="Junior Developer"
          employer-name="Rural Sourcing Inc."
          location="Augusta, GA"
          start-date="July 2018"
          end-date="June 2019"
        >
          <template slot="experience">
            <div class="mt-4">
              Worked with top executives to create a .NET Core reporting application for employee tenure on clients. As part of this:
             <ul class="list-disc">
               <li class="ml-5">
                 Designed stored procedures to aggregate the information from an internal MySQL database.
               </li>
               <li class="ml-5">
                 Wrote a .NET Core service to run these queries and generate an email to be sent to executives by the same service.
               </li>
             </ul>
            </div>
            <div class="mt-4">
              As part of an agile team worked to rewrite a client’s legacy desktop application as a full stack Java/Angular web app. As part of this:
              <ul class="list-disc">
                <li class="ml-5">
                  Wrote several Java (Spring Boot) microservices from scratch, as well as adding additional functionality and unit testing using Junit and Mockito to already existing microservice applications.
                </li>
                <li class="ml-5">
                  Contributed to the creation of several frontend Angular components and their respective unit tests.
                </li>
                <li class="ml-5">
                  All code personally contributed was 100% unit tested.
                </li>
              </ul>
            </div>
          </template>
        </employment-block>
      </template>
    </resume-block>
    <resume-block>
      <template slot="title">
        Education
      </template>
      <template slot="content">
        <div class="font-semibold text-xl">BS in Computer Science</div>
        <div>
          Augusta University
        </div>
        <div class="font-light text-gray-300">
          May 2019
        </div>
      </template>
    </resume-block>
    <resume-block class="mb-16">
      <template slot="title">
        Languages &amp; Frameworks
      </template>
      <template slot="content">
        <div class="flex flex-col xs:flex-row">
          <ul class="w-full xs:w-1/2">
            <li><experience-dots :num="10">Javascript</experience-dots></li>
            <li><experience-dots :num="9">Node.js</experience-dots></li>
            <li><experience-dots :num="7">Vue</experience-dots></li>
            <li><experience-dots :num="5">Angular</experience-dots></li>
            <li><experience-dots :num="4">golang</experience-dots></li>
            <li><experience-dots :num="6">C#</experience-dots></li>
          </ul>
          <ul class="w-full xs:w-1/2">
            <li><experience-dots :num="10">Kotlin</experience-dots></li>
            <li><experience-dots :num="7">Java</experience-dots></li>
            <li><experience-dots :num="5">Spring Boot</experience-dots></li>
            <li><experience-dots :num="9">Python</experience-dots></li>
            <li><experience-dots :num="3">Tensorflow</experience-dots></li>
            <li><experience-dots :num="5">MySQL</experience-dots></li>
          </ul>
        </div>
      </template>
    </resume-block>
  </div>
</template>

<script>
import EmploymentBlock from '../components/EmploymentBlock.vue'
import ExperienceDots from '../components/ExperienceDots.vue'
import ResumeBlock from '../components/ResumeBlock.vue'
export default {
  components: { ResumeBlock, EmploymentBlock, ExperienceDots },

}
</script>
