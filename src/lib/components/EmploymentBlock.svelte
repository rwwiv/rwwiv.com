<script lang="ts">
	export let jobTitle: string;
	export let employerName: string;
	export let location: string;
	export let startDate: Date;
	export let endDate: Date | string = "Present";
	export let tech: string[];
	export let experience: string[];
</script>

<div>
	<div class="flex flex-col text-left space-y-1">
		<div class="font-semibold text-xl">{jobTitle} @ {employerName} - {location}</div>
		<div class="">
			{#each tech as t, i}
				{t}
				{#if i != tech.length - 1}&bullet;&nbsp;{/if}
			{/each}
		</div>
		<div class="text-sm text-neutral-600 dark:text-neutral-300">
			{startDate.toLocaleDateString("en-us", {
				year: "numeric",
				month: "short",
			})}
			-
			{#if typeof endDate === "string"}
				{endDate}
			{:else}
				{endDate.toLocaleDateString("en-us", {
					year: "numeric",
					month: "short",
				})}
			{/if}
		</div>

		<ul class="space-y-2 list-disc ml-4 mt-4">
			{#each experience as exp}
				<li>{exp}</li>
			{/each}
		</ul>
	</div>
</div>
