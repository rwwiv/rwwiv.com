<script>
	import { dev } from "$app/environment";
	import Footer from "$lib/components/Footer.svelte";
	import Nav from "$lib/components/Nav.svelte";
	import { injectAnalytics } from "@vercel/analytics/sveltekit";
	import { injectSpeedInsights } from "@vercel/speed-insights/sveltekit";
	import "../app.postcss";
	import "../github-dark.scss";
	import "../github.scss";

	injectAnalytics({ mode: dev ? "development" : "production" });
	injectSpeedInsights();
</script>

<div class="flex flex-col min-h-screen">
	<Nav />
	<div class="my-12 lg:my-16 flex-1">
		<slot />
	</div>
	<Footer />
</div>
