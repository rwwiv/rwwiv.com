<script lang="ts">
	import { page } from "$app/stores";
</script>

<div id="nav">
	<a href="/" class:active={$page.url.pathname === "/"} class="text-5xl mt-3">RWWIV</a>
	<div class="flex justify-center space-x-4 mt-2 text-2xl">
		<a href="/about" class:active={$page.url.pathname === "/about"}>About</a>
		<a href="/resume" class:active={$page.url.pathname === "/resume"}>Resume</a>
		<a href="/dragons" class:active={$page.url.pathname.includes("/dragons")}>Dragons</a>
	</div>
</div>

<style lang="postcss">
	#nav {
		@apply flex flex-col items-center w-full pt-4;
	}

	#nav a {
		@apply font-bold text-neutral-500 dark:text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-400 focus:underline focus:outline-none;
	}

	#nav a.active {
		@apply text-neutral-900 dark:text-neutral-100 hover:text-neutral-700 dark:hover:text-neutral-300;
	}
</style>
