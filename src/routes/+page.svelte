<script>
	// @ts-ignore
	import underConstruction from "$lib/assets/XT35055.jpg?w=300";
</script>

<svelte:head>
	<title>RWWIV</title>
</svelte:head>

<div class="flex flex-col text-center items-center space-y-10">
	<h1>Under Construction</h1>
	<img class="w-[300px] h-[300px]" src={underConstruction} alt="Under construction" />
</div>

<style lang="postcss">
	h1 {
		@apply text-4xl;
	}
</style>
