<svelte:head>
	<title>DRAGONS</title>
</svelte:head>

<div class="flex flex-col justify-center items-center max-w-xl mx-auto space-y-10">
	<div class="space-y-2">
		<h1 class="text-4xl">
			Development Tests and Learning Exercises <span class="text-base italic"
				>&mdash; Here there be dragons 🐉</span
			>
		</h1>
	</div>
	<div class="flex flex-col space-y-4">
		<h2 class="text-3xl">Current projects:</h2>
		<ul class="max-w-xl ml-4 list-disc">
			<li>
				<a href="/dragons/image_manipulation" class="group focus:outline-none">
					<div class="">
						<div class="flex flex-col space-y-2">
							<h1 class="text-xl">Image manipulation using serverless Go functions</h1>
							<div
								class=" px-4 py-2 italic bg-neutral-100 dark:bg-neutral-800 rounded shadow-sm group-hover:outline outline-neutral-200 dark:outline-neutral-700  group-focus:bg-neutral-200 dark:group-focus:bg-neutral-700"
							>
								This experiment accepts an image uploaded on the browser and submits that image to
								serverless functions written in Go. These process the image and return the processed
								image to the client.
							</div>
						</div>
					</div>
				</a>
			</li>
		</ul>
	</div>
</div>
