<script lang="ts">
	export let jobTitle: string;
	export let employerName: string;
	export let location: string;
	export let startDate: Date;
	export let endDate: Date | string = 'Present';
	export let experience: string[];
</script>

<div>
	<div class="flex flex-col text-left">
		<div class="font-semibold text-xl">{jobTitle}</div>
		<div>
			{employerName} - {location}
		</div>
		<div class="font-light text-gray-600 dark:text-gray-300">
			{startDate.toLocaleDateString('en-us', {
				year: 'numeric',
				month: 'short'
			})}
			-
			{#if typeof endDate === 'string'}
				{endDate}
			{:else}
				{endDate.toLocaleDateString('en-us', {
					year: 'numeric',
					month: 'short'
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
