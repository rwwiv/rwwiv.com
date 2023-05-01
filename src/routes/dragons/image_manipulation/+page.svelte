<script lang="ts">
	import { env } from "$env/dynamic/public";
	import Code from "$lib/components/Code.svelte";
	import InlineCode from "$lib/components/InlineCode.svelte";
	// @ts-ignore
	import Slider from "@bulatdashiev/svelte-slider";
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

	async function handleSubmit() {
		if (image) {
			processError = false;
			loading = true; // show loading animation

			// build form
			const fData = new FormData(form);
			fData.set("image", image);

			const gRes = await fetch(`${env.PUBLIC_API_URL}/api/image/${action}`, {
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
				class="relative m-0 block w-full min-w-0 flex-auto rounded border border-solid border-neutral-300 bg-clip-padding px-3 py-[0.4rem] text-base font-normal text-neutral-700 transition duration-300 ease-in-out file:-mx-3 file:-my-[0.4rem] file:overflow-hidden file:rounded-none file:border-0 file:border-solid file:border-inherit file:bg-neutral-100 file:px-3 file:py-[0.4rem] file:text-neutral-700 file:transition file:duration-150 file:ease-in-out file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] hover:file:bg-neutral-200 focus:text-neutral-700 focus:outline-none focus:border-neutral-400 focus:dark:border-neutral-400 dark:border-neutral-600 dark:text-neutral-200 dark:file:bg-neutral-700 dark:file:text-neutral-100 file:cursor-pointer"
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
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_200ms]" />
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_400ms]" />
							<div class="rounded-full bg-white h-3 w-3 animate-[pulse_2s_infinite_600ms]" />
						</div>
					</div>
				{/if}
			</div>

			{#if processError}
				<div class="text-red-700 dark:text-red-400">Error processing image, try again.</div>
			{/if}

			<form
				bind:this={form}
				on:submit|preventDefault={handleSubmit}
				class="flex flex-col justify-center space-y-6 w-[250px] max-w-[250px]"
			>
				<input hidden name="content_type" bind:value={contentType} />
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
						class="bg-neutral-50 border border-neutral-300 text-neutral-900 text-base rounded focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-neutral-700 dark:border-neutral-600 dark:placeholder-neutral-400 dark:text-neutral-100 dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer"
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
				<button
					type="submit"
					class="py-2 px-4 text-base font-medium text-neutral-900 focus:outline-none bg-white rounded border border-neutral-200 hover:bg-neutral-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-neutral-200 dark:focus:ring-neutral-700 dark:bg-neutral-800 dark:text-neutral-400 dark:border-neutral-600 dark:hover:text-neutral-100 dark:hover:bg-neutral-700"
				>
					Submit
				</button>
			</form>
			<hr
				class="w-[150px] my-6 h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50"
			/>
			<button
				on:click={resetState}
				class="px-6 text-sm font-medium text-neutral-900 focus:outline-none bg-white rounded border border-neutral-200 hover:bg-neutral-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-neutral-200 dark:focus:ring-neutral-700 dark:bg-neutral-800 dark:text-neutral-400 dark:border-neutral-600 dark:hover:text-neutral-100 dark:hover:bg-neutral-700"
			>
				Reset
			</button>
		</div>
	{/if}
	<!-- Here be dragons... -->

	<!-- Fancy divider -->
	<div class="flex flex-col items-center mt-[250px] mb-6">
		<div class="flex flex-col items-center -space-y-1">
			<p>.</p>
			<p>.</p>
			<p>.</p>
		</div>
		<hr
			class="w-[150px] mt-4 mb-10 h-[1px] rounded-full border-t-0 bg-neutral-900 dark:bg-neutral-100 opacity-100 dark:opacity-50"
		/>
	</div>

	<!-- Write up -->
	<div class="px-4 max-w-sm sm:max-w-xl md:max-w-3xl">
		<div id="writeup" class="mb-10 flex flex-col space-y-10">
			<h1 class="font-bold italic">What is this doing?</h1>
			<div class="space-y-6">
				<p>
					At a high level we accept an uploaded image into an in-memory file blob on the client and
					then post that file as part of a multipart form along with some extra info to some
					<InlineCode code="&lcub;HOSTNAME&rcub;/api/image/&lcub;ACTION&rcub;" language="svelte" />
					URL.
				</p>
				<p>
					This URL corresponds to some Go serverless function running on
					<a
						class="focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
						href="https://vercel.com/docs/concepts/functions/serverless-functions"
						target="_blank"
						rel="noreferrer">Vercel</a
					>. That function performs some processing on the image and then responds with the
					processed image, which we then replace the in-memory file blob with client-side.
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
				<sub class="text-sm tracking-wide ml-2">
					* A possible improvement to be made here would be to join the functions into a single
					image processing function and take the action as another field in the form. This would
					greatly reduce code repetition, at the cost of an increasingly long and complicated switch
					statement.
				</sub>
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					Finally, we retrieve the new binary from the response on the client side and set the
					preview to use the new URL:
				</p>
				<Code language="ts" code={responseToPreviewCode} />
			</div>
			<h1 class="self-center">Et voilà!</h1>
			<sub class="self-center text-sm">
				<a
					href="https://github.com/rwwiv/rwwiv.com"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 hover:dark focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Site source
				</a>
				and
				<a
					href="https://github.com/rwwiv/go-serverless"
					target="_blank"
					rel="noreferrer"
					class="underline cursor-pointer hover:text-blue-700 dark:hover:text-blue-400 hover:dark focus:outline-none focus:dark:text-blue-300 focus:text-blue-700"
				>
					Go source
				</a>
			</sub>
		</div>
	</div>
</div>

<style lang="postcss">
	div {
		--thumb-bg: #dadada;
		--progress-bg: #6e6e6e;
		--track-bg: #cdcdcd;
	}
	.parent :global(.thumb) {
		@apply cursor-pointer;
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
