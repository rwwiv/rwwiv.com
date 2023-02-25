<script lang="ts">
	import { page } from "$app/stores";
</script>

<div id="nav">
	<a href="/" class:active={$page.url.pathname === "/"} class="text-5xl mt-3">RWWIV</a>
	<div class="flex justify-center space-x-4 mt-2 text-lg">
		<a href="/about" class:active={$page.url.pathname === "/about"}>About</a>
		<a href="/resume" class:active={$page.url.pathname === "/resume"}>Resume</a>
	</div>
	<div class="flex justify-center space-x-4 mt-4">
		<a href="https://www.instagram.com/rwwiv/" target="_blank" rel="noreferrer">Instagram</a>
		<a href="https://github.com/rwwiv" target="_blank" rel="noreferrer">GitHub</a>
		<a href="https://www.linkedin.com/in/rwwiv/" target="_blank" rel="noreferrer">LinkedIn</a>
	</div>
</div>

<style lang="postcss">
	#nav {
		@apply flex flex-col items-center w-full pt-4;
	}

	#nav a {
		@apply font-bold text-gray-500 hover:text-gray-600;
	}

	#nav a.active {
		@apply text-black dark:text-white hover:text-gray-700 dark:hover:text-gray-300;
	}
</style>
