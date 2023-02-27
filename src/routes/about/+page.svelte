<script>
	// @ts-ignore
	import me from "$lib/assets/me.jpg?h=350";
</script>

<svelte:head>
	<title>ABOUT</title>
</svelte:head>

<div class="flex flex-col items-center justify-center">
	<div class="flex flex-col items-center justify-center px-3 md:w-auto space-y-10">
		<div
			id="polaroid"
			class="flex flex-col items-center justify-center rounded-sm img-wrapper bg-black dark:bg-white"
		>
			<img src={me} alt="Self portrait" class="mt-4 px-3" height="350px" />
			<div class="h-[75px]" />
		</div>
		<div class="flex flex-col space-y-6 w-[20em] mb-12">
			<h1 class="text-3xl">Current Projects</h1>
			<div class="flex flex-col space-y-2">
				<div class="text-xl font-bold">Software Engineering</div>
				<ul class="list-disc space-y-2 ml-3 text-sm">
					<li class="text-sm">
						Learning embedded programming and Rust by creating a wifi-enabled plant moisture and
						light meter based on the ESP32-C3!
						<br />
						Check back soon for more on that project.
					</li>
					<li class="text-sm">
						This website! This site is built using Sveltekit for now, with a Go backend coming soon.
						<br />
						Check
						<a
							href="https://github.com/rwwiv/rwwiv.com"
							class="underline hover:text-blue-400 cursor-pointer"
						>
							here
						</a>
						for the public repo!
					</li>
				</ul>
			</div>
			<div class="flex flex-col space-y-1">
				<div class="text-2xl font-bold">Photography</div>
				<ul class="list-disc space-y-2 ml-3 text-sm">
					<li>
						Building a street photography portfolio with my Fujifilm X-T3
						<br />
						Check
						<a href="/" class="underline cursor-pointer hover:text-blue-400">
							this site's home page
						</a> soon for some of my work!
					</li>
					<li>Slowly (oh so slowly) learning film photography on a Canon AE-1</li>
				</ul>
			</div>
		</div>
	</div>
</div>
