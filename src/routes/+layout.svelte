<script>
	import { browser } from "$app/environment";
	import { page } from "$app/stores";
	import Footer from "$lib/components/Footer.svelte";
	import Nav from "$lib/components/Nav.svelte";
	import { webVitals } from "$lib/vitals";
	import "../app.css";
	import "../github-dark.scss";
	import "../github.scss";

	let analyticsId = import.meta.env.VERCEL_ANALYTICS_ID;

	$: if (browser && analyticsId) {
		webVitals({
			path: $page.url.pathname,
			params: $page.params,
			analyticsId,
		});
	}
</script>

<div class="flex flex-col min-h-screen">
	<Nav />
	<div class="my-12 lg:my-16 flex-1">
		<slot />
	</div>
	<Footer />
</div>
