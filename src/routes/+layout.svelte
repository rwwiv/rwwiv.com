<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import '../app.css';

	console.log($page);

	onMount(async () => {
		const colorSchemeString =
			window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? '-dark' : '';

		const darkMediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
	});
</script>

<svelte:head>
	<title>RWWIV</title>
</svelte:head>

<div id="content" class="z-50">
	<div id="nav">
		<a href="/" class:active={$page.url.pathname === '/'} class="text-5xl mt-3">RWWIV</a>
		<div class="flex justify-center space-x-4 mt-1 text-lg">
			<a href="/about" class:active={$page.url.pathname === '/about'}>About</a>
			<a href="/resume" class:active={$page.url.pathname === '/resume'}>Resume</a>
		</div>
		<div class="flex justify-center space-x-4 mt-4">
			<a href="https://www.instagram.com/rwwiv/" target="_blank" rel="noreferrer">Instagram</a>
			<a href="https://github.com/rwwiv" target="_blank" rel="noreferrer">GitHub</a>
			<a href="https://www.linkedin.com/in/rwwiv/" target="_blank" rel="noreferrer">LinkedIn</a>
		</div>
	</div>
	<div class="mt-12 lg:mt-16 mb-24">
		<slot />
	</div>
</div>

<style lang="postcss">
	#nav {
		@apply flex flex-col items-center w-full pt-4;
	}

	#nav a {
		font-weight: bold;
		color: grey;
	}

	#nav a.active {
		color: white;
	}
</style>
