<script>
	// @ts-ignore
	import underConstruction from "$lib/assets/XT35055.jpg?w=300";
	// @ts-ignore
	import teachmeto from "$lib/assets/teachmeto-ui.jpg";
	// @ts-ignore
	import banjohealth from "$lib/assets/banjohealth-ui.jpg";

	import MoreDivider from "$lib/components/MoreDivider.svelte";
	import ProjectPreview from "$lib/components/ProjectPreview.svelte";
</script>

<svelte:head>
	<title>RWWIV</title>
</svelte:head>

<div class="flex flex-col text-center items-center space-y-10">
	<h1>Under Construction</h1>
	<img class="w-[300px]" src={underConstruction} alt="Under construction" />
	<div class="h-[50px]" />
	<MoreDivider />
	<h1>Recent Web Exercises</h1>
	<div class="flex flex-col md:flex-row space-y-24 md:space-y-0 md:space-x-12">
		<ProjectPreview
			src={teachmeto}
			title="T3 listing page"
			subtitle="Take-home technical interview"
			sourceUrl="https://github.com/rwwiv/teachmeto-demo/"
			siteUrl="https://teachmeto.rwwiv.com"
		/>
		<ProjectPreview
			src={banjohealth}
			title="Sveltekit paginated table"
			subtitle="Take-home technical interview"
			sourceUrl="https://github.com/rwwiv/banjo-ui-demo"
		/>
	</div>
</div>

<style lang="postcss">
	h1 {
		@apply text-4xl;
	}
</style>
