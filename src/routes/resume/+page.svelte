<script lang="ts">
	import EmploymentBlock from "$lib/components/EmploymentBlock.svelte";
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
		<div class="text-4xl font-semibold text-green-800 dark:text-green-200">William Wernert</div>
		<div class="text-lg">Software Engineer</div>
	</div>
	<ResumeBlock title="Summary">
		<div>
			Software engineer building scalable cloud-native services at Grafana Labs, developing and
			operating alerting infrastructure serving millions of users. Experienced across the stack
			from web applications to distributed systems, with current focus on Go, Kubernetes, and
			infrastructure-as-code (Jsonnet, Flux, Crossplane). Deep familiarity with observability
			tooling through dogfooding the Grafana stack (Mimir, Loki, Tempo). Excited to bring my
			skills and experience to your software engineering team.
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
			<div class="font-semibold text-xl">
				{data.education.degree}
			</div>
			<div>{data.education.university}</div>
			<div class="text-sm text-neutral-600 dark:text-neutral-300">
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
		<div class="flex flex-col space-y-5">
			{#each data.volunteerRoles as v}
				<div>
					<div class="text-lg">
						<span class="font-semibold">{v.group}</span> - {v.role}
					</div>
					<div class="text-sm text-neutral-600 dark:text-neutral-300">
						{#if v.endYear && v.startYear.getFullYear() === v.endYear.getFullYear()}
							{v.startYear.toLocaleDateString("en-us", { year: "numeric" })}
						{:else}
							{v.startYear.toLocaleDateString("en-us", { year: "numeric" })}
							-
							{#if v.endYear}
								{v.endYear.toLocaleDateString("en-us", { year: "numeric" })}
							{:else}Present{/if}
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</ResumeBlock>
</div>
