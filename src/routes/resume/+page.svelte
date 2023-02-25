<script lang="ts">
	import EmploymentBlock from '$lib/components/EmploymentBlock.svelte';
	import ExperienceDots from '$lib/components/ExperienceDots.svelte';
	import ResumeBlock from '$lib/components/ResumeBlock.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>RESUME</title>
</svelte:head>

<div class="flex flex-col items-center">
	<div class="text-center mb-6">
		<div class="text-4xl text-red-400">DRAFT</div>
		<br />
		<div class="text-4xl">William Wernert</div>
		<div class="text-lg font-light">Full Stack Developer</div>
	</div>
	<ResumeBlock title="Summary">
		<div>
			Experienced software engineer with a track record of delivering successful projects in various
			tech stacks including Python, Go, Vue, and Spring Boot. Specializes in developing and
			deploying web applications, microservices, and machine learning-based tools. Excited to bring
			my skills and experience to your software engineering team.
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
					experience={e.experience}
				/>
			{/each}
		</div>
	</ResumeBlock>
	<ResumeBlock title="Education">
		<div class="font-semibold text-xl">{data.education.degree}</div>
		<div>{data.education.university}</div>
		<div class="font-light text-gray-600 dark:text-gray-300">
			{data.education.graduationDate.toLocaleDateString('en-us', {
				year: 'numeric',
				month: 'short'
			})}
		</div>
	</ResumeBlock>
	<ResumeBlock title="Languages &amp; Frameworks">
		<div class="flex flex-col xs:flex-row">
			<ul class="w-full xs:w-1/2">
				{#each data.skills as skill, i}
					{#if i % 2 == 0}
						<ExperienceDots num={skill.level} name={skill.name} />
					{/if}
				{/each}
			</ul>
			<ul class="w-full xs:w-1/2">
				{#each data.skills as skill, i}
					{#if i % 2 != 0}
						<ExperienceDots num={skill.level} name={skill.name} />
					{/if}
				{/each}
			</ul>
		</div>
	</ResumeBlock>
</div>
