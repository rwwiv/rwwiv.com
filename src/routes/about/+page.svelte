<svelte:head>
	<title>ABOUT</title>
</svelte:head>

<div class="flex flex-col items-center justify-center">
	<div class="flex flex-col items-center justify-center px-3 md:w-auto space-y-10">
		<div
			id="polaroid"
			class="flex flex-col items-center justify-center rounded-sm img-wrapper bg-black dark:bg-white"
		>
			<img
				src="$lib/assets/me.jpg"
				alt="Self portrait"
				class="mt-4 px-3 h-full"
				height="350px"
				width="350px"
			/>
			<div class="h-[80px]" />
		</div>
		<div class="flex flex-col space-y-6 w-[20em] mb-12">
			<h1 class="text-3xl">Current Projects</h1>
			<div class="flex flex-col space-y-2">
				<div class="text-xl font-bold">Software Engineering</div>
				<ul class="list-disc space-y-2 ml-3 text-sm">
					<li class="text-sm">
						Currently working on an embedded programming project using Rust to create a wifi-enabled
						plant moisture and light meter based on the ESP32-C3!
						<br />
						Check back soon for more on that project.
					</li>
					<li class="text-sm">
						This website! This site is built using Sveltekit for now, with a Go backend coming soon.
						<br />
						Check my
						<span class="underline hover:text-red-400">
							<a href="https://github.com/rwwiv">Github </a>
						</span>
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
						Check this site's home page soon for some of my work!
					</li>
					<li>Slowly learning film photography</li>
				</ul>
			</div>
		</div>
	</div>
</div>
