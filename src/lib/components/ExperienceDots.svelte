<script lang="ts">
	export let num: number;
	export let name: string;

	$: remainder = 10 - num;
</script>

<div class="flex flex-row">
	<div class="mr-2 w-1/2 xxs:w-1/3 text-right">
		{name}
	</div>
	<div class="flex flex-row w-1/2 xxs:w-2/3 font-mono	">
		{#each Array(num) as _}
			<div class="font-extrabold">&bull;</div>
		{/each}
		{#each Array(remainder) as _}
			<div class="text-gray-400">-</div>
		{/each}
	</div>
</div>
