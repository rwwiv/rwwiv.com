<script>
	import Footer from '$lib/components/Footer.svelte';
	import Nav from '$lib/components/Nav.svelte';
	import '../app.css';
</script>

<svelte:head>
	<title>RWWIV</title>
</svelte:head>

<div class="flex flex-col min-h-screen">
	<Nav />
	<div class="mt-12 lg:mt-16 mb-24 flex-1">
		<slot />
	</div>
	<Footer />
</div>
