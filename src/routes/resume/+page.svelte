<script lang="ts">
	import EmploymentBlock from "$lib/components/EmploymentBlock.svelte";
	import ExperienceDots from "$lib/components/ExperienceDots.svelte";
	import ResumeBlock from "$lib/components/ResumeBlock.svelte";
	import type { PageData } from "./$types";

	export let data: PageData;
</script>

<svelte:head>
	<title>RESUME</title>
</svelte:head>

<div class="flex flex-col items-center">
	<div class="text-center mb-6">
		<!--
			<div class="text-4xl text-red-400">DRAFT</div>
			<br />
		-->
		<div class="text-4xl">William Wernert</div>
		<div class="text-lg font-light">Full Stack Engineer</div>
	</div>
	<ResumeBlock title="Summary">
		<div>
			Experienced software engineer with a proven track record of delivering successful projects in
			various technologies including Python (Flask, Django), Go, Vue.js, and Java (Spring Boot).
			Specializes in developing and deploying web applications, micro-services, and machine
			learning-based tools. Excited to bring my skills and experience to your software engineering
			team.
		</div>
	</ResumeBlock>
	<ResumeBlock title="Work Experience">
		<div class="flex flex-col space-y-6">
			{#each data.employments as e}
				<EmploymentBlock
					jobTitle={e.jobTitle}
					employerName={e.employerName}
					location={e.location}
					startDate={e.startDate}
					endDate={e.endDate}
					tech={e.tech}
					experience={e.experience}
				/>
			{/each}
		</div>
	</ResumeBlock>
	<ResumeBlock title="Education">
		<div class="flex flex-col space-y-1">
			<div class="font-semibold text-xl">{data.education.degree}</div>
			<div>{data.education.university}</div>
			<div class="text-sm text-gray-600 dark:text-gray-300">
				{data.education.graduationDate.toLocaleDateString("en-us", {
					year: "numeric",
					month: "short",
				})}
			</div>
		</div>
	</ResumeBlock>
	<ResumeBlock title="Personal Projects">
		<ul class="list-disc space-y-1">
			{#each data.projects as project}
				<li class="ml-4">{project}</li>
			{/each}
		</ul>
	</ResumeBlock>
	<ResumeBlock title="Volunteer Work">
		<div class="flex flex-col space-y-1">
			{#each data.volunteerRoles as v}
				<div class="font-semibold text-xl">{v.group}</div>
				<div>
					{v.role}
				</div>
				<div class="text-sm text-gray-600 dark:text-gray-300">
					{v.startYear.toLocaleDateString("en-us", { year: "numeric" })}
					-
					{#if v.endYear}
						{v.endYear.toLocaleDateString("en-us", { year: "numeric" })}
					{:else}Present{/if}
				</div>
			{/each}
		</div>
	</ResumeBlock>
	<ResumeBlock title="Languages &amp; Frameworks">
		<div class="flex flex-col xs:flex-row">
			<ul class="w-full xs:w-1/2 space-y-2">
				{#each data.skills as skill, i}
					{#if i % 2 == 0}
						<ExperienceDots num={skill.level} name={skill.name} />
					{/if}
				{/each}
			</ul>
			<ul class="w-full xs:w-1/2 space-y-2">
				{#each data.skills as skill, i}
					{#if i % 2 != 0}
						<ExperienceDots num={skill.level} name={skill.name} />
					{/if}
				{/each}
			</ul>
		</div>
	</ResumeBlock>
</div>
