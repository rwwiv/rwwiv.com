<script lang="ts">
	import Code from "$lib/components/Code.svelte";

	export let link: string;
	export let title: string;
	export let code: string;
	export let language: string;
</script>

<a href={link} class="group focus:outline-none">
	<div>
		<div class="flex flex-col space-y-2">
			<h1
				class="text-2xl text-green-700 dark:text-green-200 group-hover:text-green-500 dark:group-hover:text-green-400"
			>
				{title}
			</h1>
			<Code {code} {language} />
		</div>
	</div>
</a>
