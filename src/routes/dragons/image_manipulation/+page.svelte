<script lang="ts">
	import { env } from "$env/dynamic/public";
	import Code from "$lib/components/Code.svelte";
	import InlineCode from "$lib/components/InlineCode.svelte";
	import MoreDivider from "$lib/components/MoreDivider.svelte";
	import Slider from "$lib/components/Slider.svelte";
	import { tick } from "svelte";
	import { fade } from "svelte/transition";

	let image: Blob | null;
	let contentType: string;
	let form: HTMLFormElement;
	let input: HTMLInputElement;
	let imgEl: HTMLImageElement;

	let showPreview = false;
	let loading = false;
	let processError = false;

	const actions = {
		grayscale: "Grayscale",
		invert: "Negative",
		blur: "Blur",
		pixelate: "Pixelate",
	};

	let action: string;
	let blurStrength = [10];
	let pixelSize = [10];

	// Chain mode state
	let chainMode = false;
	type Transform = {
		type: string;
		blur_strength?: number;
		pixel_size?: number;
	};
	let chainTransforms: Transform[] = [];
	let newTransformType = "grayscale";
	let newBlurStrength = [10];
	let newPixelSize = [10];

	function addTransform() {
		const transform: Transform = { type: newTransformType };
		if (newTransformType === "blur") {
			transform.blur_strength = newBlurStrength[0];
		} else if (newTransformType === "pixelate") {
			transform.pixel_size = newPixelSize[0];
		}
		chainTransforms = [...chainTransforms, transform];
	}

	function removeTransform(index: number) {
		chainTransforms = chainTransforms.filter((_, i) => i !== index);
	}

	function moveTransform(index: number, direction: "up" | "down") {
		const newIndex = direction === "up" ? index - 1 : index + 1;
		if (newIndex < 0 || newIndex >= chainTransforms.length) return;
		const newTransforms = [...chainTransforms];
		[newTransforms[index], newTransforms[newIndex]] = [
			newTransforms[newIndex],
			newTransforms[index],
		];
		chainTransforms = newTransforms;
	}

	async function handleSubmit() {
		if (image) {
			processError = false;
			loading = true; // show loading animation

			// build form
			const fData = new FormData(form);
			fData.set("image", image);

			let endpoint: string;
			if (chainMode) {
				endpoint = `${env.PUBLIC_API_URL}/api/image/chain`;
				fData.set("transforms", JSON.stringify(chainTransforms));
			} else {
				endpoint = `${env.PUBLIC_API_URL}/api/image/${action}`;
			}

			const gRes = await fetch(endpoint, {
				method: "POST",
				body: fData,
			});

			if (gRes.status < 400) {
				image = await gRes.blob(); // set image to blob from response
			} else {
				processError = true;
			}

			loading = false;
			setPreviewSrc(image);
		}
	}

	async function resetState() {
		image = null;
		showPreview = false;
		processError = false;
		chainTransforms = [];
	}

	function setPreviewSrc(img: Blob) {
		imgEl.src = URL.createObjectURL(img);
		imgEl.onload = () => {
			URL.revokeObjectURL(imgEl.src); // free memory
		};
	}

	async function previewImage() {
		// check if anything was uploaded
		if (!input.files) {
			showPreview = false;
			return;
		}
		showPreview = true;
		// wait for DOM to settle so we can access the preview element
		await tick();

		image = input.files[0];
		contentType = input.files[0].type;
		input.value = ""; // clear input since we set the image field ourselves

		setPreviewSrc(image);
	}

	const previewCode = `function setPreviewSrc(img: Blob) {
	// imgEl refers to a bound <img /> element
	imgEl.src = URL.createObjectURL(img);
	imgEl.onload = () => {
		URL.revokeObjectURL(imgEl.src); // free memory
	};
}`;

	const formText = `{#if showPreview}

	<!-- PREVIEW -->

	<form bind:this={form} on:submit|preventDefault={handleSubmit} class="...">
		<input hidden name="content_type" bind:value={contentType} />
		<div>
			<label for="formAction" class="..."> Select an action </label>
			<select id="formAction" name="action" bind:value={action} class="...">
				{#each Object.entries(actions) as [value, name]}
					<option {value}>{name}</option>
				{/each}
			</select>
		</div>
		{#if action == "blur"}
			<!-- BLUR SLIDER -->
		{:else if action == "pixelate"}
			<!-- PIXELATE SLIDER -->
		{/if}
		...
	</form>
{/if}`;

	const transformRequestText = `// build form
const fData = new FormData(form);
fData.set("image", image);

const gRes = await fetch(\`http://localhost:3000/api/image/\${action}\`, {
	method: "POST",
	body: fData,
});`;

	const goImageExtractCode = `// this keeps the whole form in memory, so be careful of space
err := r.ParseMultipartForm(32 << 20) // max 32MB
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	return
}

file, _, err := r.FormFile("image")
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	return
}
defer file.Close()

contentType := r.PostFormValue("content_type")

// util.Decode() is an internal function that contains
// a simple switch statement used to decode the image
// correctly based on contentType
img, err := util.Decode(file, contentType)
if err != nil {
	w.WriteHeader(http.StatusUnsupportedMediaType)
	return
}`;

	const goImageResponseCode = `processedImg := someProcess(img)

outputBuf := new(bytes.Buffer)
// like util.Decode(), util.Encode() is an internal function
// that contains a simple switch statement used to encode the image
err = util.Encode(outputBuf, processedImg, contentType)
if err != nil {
	panic(err)
}

w.Header().Set("Content-Type", "application/octet-stream")
w.Write(outputBuf.Bytes())`;

	const responseToPreviewCode = `if (gRes.status < 400) {
	// set image to blob from response
	image = await gRes.blob();
}

setPreviewSrc(image);`;

	const chainRequestCode = `// Build the transforms array
const transforms = [
  { type: "blur", blur_strength: 5 },
  { type: "grayscale" },
  { type: "pixelate", pixel_size: 10 }
];

const fData = new FormData(form);
fData.set("image", image);
fData.set("transforms", JSON.stringify(transforms));

const gRes = await fetch(\`\${API_URL}/api/image/chain\`, {
  method: "POST",
  body: fData,
});`;

	const chainGoCode = `// Parse the transforms JSON from the form
transformsJSON := r.PostFormValue("transforms")
var specs []TransformSpec
if err := json.Unmarshal([]byte(transformsJSON), &specs); err != nil {
  w.WriteHeader(http.StatusBadRequest)
  return
}

// Apply each transform in order using the registry
currentImg := img
for _, spec := range specs {
  constructor, ok := transform.Registry[spec.Type]
  if !ok {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  t, err := constructor(spec.Params)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  currentImg, err = t.Apply(currentImg)
  if err != nil {
    panic(err)
  }
}`;
</script>

<svelte:head>
	<title>DRAGON 1</title>
</svelte:head>

<div class="flex flex-col justify-center items-center px-4 parent">
	<h1 class="text-5xl text-center mb-10">Serverless Image Manipulation</h1>
	{#if !image}
		<div>
			<label
				for="formFile"
				class="block mb-2 text-sm font-medium text-neutral-900 dark:text-neutral-100"
			>
				Upload an image
			</label>
			<input
				name="image"
				bind:this={input}
				on:change={previewImage}
				accept="image/jpeg, image/png"
				class="file-input"
				type="file"
				id="formFile"
			/>
		</div>
	{/if}

	{#if showPreview}
		<div transition:fade class="flex flex-col items-center space-y-6">
			<div class="relative max-w-[800px]">
				<img id="preview" bind:this={imgEl} alt="Upload preview" />
				{#if loading}
					<div transition:fade class="absolute top-0 left-0 h-full w-full bg-black bg-opacity-60">
						<div
							class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 flex items-center space-x-2"
						>
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_200ms]"></div>
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_400ms]"></div>
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_600ms]"></div>
						</div>
					</div>
				{/if}
			</div>

			{#if processError}
				<div class="text-red-700 dark:text-red-400">Error processing image, try again.</div>
			{/if}

			<!-- Mode Toggle -->
			<div class="flex rounded border border-neutral-300 dark:border-neutral-600 overflow-hidden">
				<button
					type="button"
					on:click={() => (chainMode = false)}
					class="px-4 py-2 text-sm font-medium transition-colors {!chainMode
						? 'bg-neutral-700 text-white dark:bg-neutral-200 dark:text-neutral-900'
						: 'bg-neutral-100 text-neutral-700 dark:bg-neutral-800 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}"
				>
					Single
				</button>
				<button
					type="button"
					on:click={() => (chainMode = true)}
					class="px-4 py-2 text-sm font-medium transition-colors {chainMode
						? 'bg-neutral-700 text-white dark:bg-neutral-200 dark:text-neutral-900'
						: 'bg-neutral-100 text-neutral-700 dark:bg-neutral-800 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}"
				>
					Chain
				</button>
			</div>

			<form
				bind:this={form}
				on:submit|preventDefault={handleSubmit}
				class="flex flex-col justify-center space-y-6 w-[300px] max-w-[300px]"
			>
				<input hidden name="content_type" bind:value={contentType} />

				{#if !chainMode}
					<!-- Single Action Mode -->
					<div>
						<label
							for="formAction"
							class="block mb-2 text-sm font-medium text-neutral-900 dark:text-neutral-100"
						>
							Select an action
						</label>
						<select
							id="formAction"
							name="action"
							bind:value={action}
							class="bg-white border border-neutral-200 text-neutral-900 text-base rounded focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-neutral-700 dark:border-neutral-600 dark:placeholder-neutral-400 dark:text-neutral-100 dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer"
						>
							{#each Object.entries(actions) as [value, name]}
								<option {value}>{name}</option>
							{/each}
						</select>
					</div>
					{#if action == "blur"}
						<div>
							<label
								for="blurStrength"
								class="block mb-1 text-sm font-medium text-neutral-900 dark:text-neutral-100"
							>
								Blur strength: {blurStrength[0]}
							</label>
							<span id="blurStrength">
								<Slider name={["blur_strength"]} bind:value={blurStrength} min={0} max={100} />
							</span>
						</div>
					{:else if action == "pixelate"}
						<div>
							<label
								for="pixelSize"
								class="block mb-1 text-sm font-medium text-neutral-900 dark:text-neutral-100"
							>
								Pixel size: {pixelSize[0]}
							</label>
							<span id="pixelSize">
								<Slider name={["pixel_size"]} bind:value={pixelSize} min={1} max={50} />
							</span>
						</div>
					{/if}
				{:else}
					<!-- Chain Mode -->
					<div class="space-y-4">
						<p class="text-sm text-neutral-600 dark:text-neutral-400">
							Build a chain of transforms to apply in order.
						</p>

						<!-- Transform List -->
						{#if chainTransforms.length > 0}
							<div class="space-y-2">
								{#each chainTransforms as transform, i}
									<div
										class="flex items-center justify-between p-2 bg-neutral-100 dark:bg-neutral-800 rounded border border-neutral-200 dark:border-neutral-700"
									>
										<div class="flex items-center space-x-2">
											<span class="text-xs text-neutral-500 dark:text-neutral-400 w-4"
												>{i + 1}.</span
											>
											<span class="text-sm font-medium">{actions[transform.type]}</span>
											{#if transform.blur_strength !== undefined}
												<span class="text-xs text-neutral-500 dark:text-neutral-400"
													>({transform.blur_strength})</span
												>
											{/if}
											{#if transform.pixel_size !== undefined}
												<span class="text-xs text-neutral-500 dark:text-neutral-400"
													>({transform.pixel_size}px)</span
												>
											{/if}
										</div>
										<div class="flex items-center space-x-1">
											<button
												type="button"
												on:click={() => moveTransform(i, "up")}
												disabled={i === 0}
												class="p-1 text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300 disabled:opacity-30 disabled:cursor-not-allowed"
												title="Move up"
											>
												<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
													<path
														stroke-linecap="round"
														stroke-linejoin="round"
														stroke-width="2"
														d="M5 15l7-7 7 7"
													/>
												</svg>
											</button>
											<button
												type="button"
												on:click={() => moveTransform(i, "down")}
												disabled={i === chainTransforms.length - 1}
												class="p-1 text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300 disabled:opacity-30 disabled:cursor-not-allowed"
												title="Move down"
											>
												<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
													<path
														stroke-linecap="round"
														stroke-linejoin="round"
														stroke-width="2"
														d="M19 9l-7 7-7-7"
													/>
												</svg>
											</button>
											<button
												type="button"
												on:click={() => removeTransform(i)}
												class="p-1 text-red-500 hover:text-red-700 dark:hover:text-red-400"
												title="Remove"
											>
												<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
													<path
														stroke-linecap="round"
														stroke-linejoin="round"
														stroke-width="2"
														d="M6 18L18 6M6 6l12 12"
													/>
												</svg>
											</button>
										</div>
									</div>
								{/each}
							</div>
						{/if}

						<!-- Add Transform -->
						<div
							class="p-3 border border-dashed border-neutral-300 dark:border-neutral-600 rounded space-y-3"
						>
							<div>
								<label
									for="newTransformType"
									class="block mb-1 text-xs font-medium text-neutral-700 dark:text-neutral-300"
								>
									Transform type
								</label>
								<select
									id="newTransformType"
									bind:value={newTransformType}
									class="bg-white border border-neutral-200 text-neutral-900 text-sm rounded focus:ring-blue-500 focus:border-blue-500 block w-full p-2 dark:bg-neutral-700 dark:border-neutral-600 dark:text-neutral-100 cursor-pointer"
								>
									{#each Object.entries(actions) as [value, name]}
										<option {value}>{name}</option>
									{/each}
								</select>
							</div>
							{#if newTransformType === "blur"}
								<div>
									<label
										for="newBlurStrength"
										class="block mb-1 text-xs font-medium text-neutral-700 dark:text-neutral-300"
									>
										Blur strength: {newBlurStrength[0]}
									</label>
									<span id="newBlurStrength">
										<Slider bind:value={newBlurStrength} min={0} max={100} />
									</span>
								</div>
							{:else if newTransformType === "pixelate"}
								<div>
									<label
										for="newPixelSize"
										class="block mb-1 text-xs font-medium text-neutral-700 dark:text-neutral-300"
									>
										Pixel size: {newPixelSize[0]}
									</label>
									<span id="newPixelSize">
										<Slider bind:value={newPixelSize} min={1} max={50} />
									</span>
								</div>
							{/if}
							<button
								type="button"
								on:click={addTransform}
								class="w-full py-1.5 px-3 text-sm font-medium text-neutral-700 bg-neutral-200 rounded hover:bg-neutral-300 dark:bg-neutral-700 dark:text-neutral-200 dark:hover:bg-neutral-600 transition-colors"
							>
								+ Add Transform
							</button>
						</div>
					</div>
				{/if}

				<button
					type="submit"
					disabled={chainMode && chainTransforms.length === 0}
					class="py-2 px-4 text-base font-medium text-neutral-800 focus:outline-none bg-neutral-50 rounded border border-neutral-200 hover:bg-neutral-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-neutral-200 dark:focus:ring-neutral-700 dark:bg-neutral-800 dark:text-neutral-50 dark:border-neutral-600 dark:hover:text-neutral-100 dark:hover:bg-neutral-700 disabled:opacity-50 disabled:cursor-not-allowed"
				>
					Submit
				</button>
			</form>
			<hr
				class="w-[150px] my-6 h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50"
			/>
			<button
				on:click={resetState}
				class="px-6 text-sm font-medium text-neutral-800 focus:outline-none bg-neutral-50 rounded border border-neutral-200 hover:bg-neutral-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-neutral-200 dark:focus:ring-neutral-700 dark:bg-neutral-800 dark:text-neutral-50 dark:border-neutral-600 dark:hover:text-neutral-100 dark:hover:bg-neutral-700"
			>
				Reset
			</button>
		</div>
	{/if}
	<!-- Here be dragons... -->

	<!-- Fancy divider -->
	<MoreDivider />

	<!-- Write up -->
	<div class="px-4 max-w-sm sm:max-w-xl md:max-w-3xl">
		<div id="writeup" class="mb-10 flex flex-col space-y-10">
			<h1 class="font-bold italic">What is this doing?</h1>
			<div class="space-y-6">
				<p>
					At a high level we accept an uploaded image into an in-memory file blob in the browser,
					and then post that file as part of a multipart form along with some extra info to a URL
					with the shape
					<InlineCode code="&lcub;HOSTNAME&rcub;/api/image/&lcub;ACTION&rcub;" language="svelte" />.
				</p>
				<p>
					This URL corresponds to a Go serverless function deployed alongside this site as part of
					the same monorepo on
					<a
						class="focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
						href="https://vercel.com/docs/functions/runtimes/go"
						target="_blank"
						rel="noreferrer">Vercel</a
					>. That function processes the image and returns the result, which we use to replace the
					in-memory file blob client-side.
				</p>
			</div>
			<hr
				class="w-[150px] h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50 self-center"
			/>
			<h1 class="font-bold italic">How is it doing it?</h1>
			<div class="flex flex-col space-y-2">
				<p>
					When an image is uploaded we collect the image into a file blob, then use <InlineCode
						code="URL.createObjectURL()"
						language="typescript"
					/>
					to display a preview of the working image:
				</p>
				<Code language="ts" code={previewCode} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					After the image is uploaded we show a preview, as well as a form that contains several
					inputs as described here. A select sets the action to be applied to the image, and also
					toggles additional inputs depending on the option selected. We use a hidden input to send
					the content type (collected on upload) along with the form data:
				</p>
				<Code language="svelte" code={formText} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					We then use the Javascript Fetch API to POST to the relevant transform API endpoint
					&mdash; our serverless function:
				</p>
				<Code language="ts" code={transformRequestText} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					In the Go function we extract the image from the multipart form and then decode it based
					on the content type:
				</p>
				<Code language="go" code={goImageExtractCode} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					Once we decode the image we perform some processing on it depending on the function being
					called. We then re-encode the image to a buffer and write it to the response:
				</p>
				<Code language="go" code={goImageResponseCode} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					Finally, we retrieve the new binary from the response on the client side and set the
					preview to use the new URL:
				</p>
				<Code language="ts" code={responseToPreviewCode} />
			</div>
			<hr
				class="w-[150px] h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50 self-center"
			/>
			<h1 class="font-bold italic">New for 2025!</h1>
			<div class="space-y-6">
				<p>
					Rather than applying a single transform at a time and making multiple round trips to the
					server, we now support chaining multiple transforms in a single request. This lets you
					build a pipeline of operations that are applied in sequence.
				</p>
				<p>
					On the client side, we build an array of transform objects and serialize it as JSON in the
					form data:
				</p>
				<Code language="ts" code={chainRequestCode} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					On the server, we parse the transforms array and apply each operation in order. A registry
					maps transform types to their constructors, making it easy to add new transforms without
					modifying the handler:
				</p>
				<Code language="go" code={chainGoCode} />
			</div>
			<!-- <div class="space-y-6"> -->
			<!-- 	<p> -->
			<!-- 		The chain endpoint reduces latency for complex edits and makes the API more flexible -->
			<!-- 		&mdash; you can now blur, convert to grayscale, and pixelate all in one request! -->
			<!-- 	</p> -->
			<!-- </div> -->
			<h1 class="self-center">Et voilà!</h1>
			<sub class="self-center text-sm">
				<a
					href="https://github.com/rwwiv/rwwiv.com"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Source on GitHub
				</a>
			</sub>
		</div>
	</div>
</div>

<style lang="postcss">
	@reference "#styles";

	div {
		--thumb-bg: #dadada;
		--progress-bg: #6e6e6e;
		--track-bg: #cdcdcd;
	}
	.parent :global(.thumb) {
		@apply cursor-pointer;
	}

	.file-input {
		@apply block w-full text-base rounded border border-neutral-300 cursor-pointer;
		@apply bg-neutral-50 text-neutral-700;
		@apply focus:outline-none focus:border-neutral-400;
	}
	.file-input::file-selector-button {
		@apply px-4 py-2 mr-3 border-0 border-r border-neutral-300;
		@apply bg-neutral-200 text-neutral-700 font-medium cursor-pointer;
		@apply hover:bg-neutral-300 transition-colors duration-150;
	}
	:global(.dark) .file-input {
		@apply bg-neutral-800 text-neutral-200 border-neutral-600;
	}
	:global(.dark) .file-input::file-selector-button {
		@apply bg-neutral-700 text-neutral-100 border-neutral-600;
		@apply hover:bg-neutral-600;
	}

	#writeup {
		@apply font-serif;
	}

	#writeup h1 {
		@apply text-2xl;
	}

	#writeup p {
		@apply text-lg leading-relaxed;
	}

	#writeup a {
		@apply underline underline-offset-2;
		@apply dark:hover:text-blue-400 hover:text-blue-700;
	}
</style>
