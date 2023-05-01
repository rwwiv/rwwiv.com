<script lang="ts">
	import DarkMode from "./DarkMode.svelte";
</script>

<div class="flex flex-col items-center justify-center my-4">
	<DarkMode />

	<div id="bottom-nav" class="flex justify-center space-x-4 mt-4 text-sm">
		<a href="https://calendly.com/rwwiv/30min" target="_blank" rel="noreferrer">Calendly</a>
		<a href="https://www.instagram.com/rwwiv/" target="_blank" rel="noreferrer">Instagram</a>
		<a href="https://github.com/rwwiv" target="_blank" rel="noreferrer">GitHub</a>
		<a href="https://www.linkedin.com/in/rwwiv/" target="_blank" rel="noreferrer">LinkedIn</a>
	</div>

	<div class="flex space-x-2 text-sm mt-2">
		<div>&copy; 2023 William Wernert</div>
	</div>
</div>

<style lang="postcss">
	#bottom-nav a {
		@apply hover:dark:text-neutral-400 hover:text-neutral-500 focus:underline focus:outline-none;
	}
</style>
