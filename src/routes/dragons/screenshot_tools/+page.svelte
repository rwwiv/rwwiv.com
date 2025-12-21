<script lang="ts">
	import Code from "$lib/components/Code.svelte";
	import InlineCode from "$lib/components/InlineCode.svelte";

	const colorDetect = `PRESET=$(hyprctl monitors -j | jq -r \\
    '.[] | select(.focused) | .colorManagementPreset // "srgb"')`;

	const colorConvert = `case "$PRESET" in
    dp3)
        CONVERT_CMD="magick png:- -set colorspace DisplayP3 -colorspace sRGB png:-"
        ;;
    adobe)
        CONVERT_CMD="magick png:- -set colorspace Adobe98 -colorspace sRGB png:-"
        ;;
esac`;

	const grimPipe = `if [[ -n "$CONVERT_CMD" ]]; then
    /usr/bin/grim "\${ARGS[@]}" - | eval "$CONVERT_CMD"
else
    /usr/bin/grim "\${ARGS[@]}" "$OUTPUT_TARGET"
fi`;
</script>

<svelte:head>
	<title>DRAGON 2</title>
</svelte:head>

<div class="flex flex-col justify-center items-center px-4">
	<h1 class="text-5xl text-center mb-10">grim Color Space Wrapper</h1>

	<div class="px-4 max-w-sm sm:max-w-xl md:max-w-3xl">
		<div id="writeup" class="mb-10 flex flex-col space-y-10">
			<div class="flex flex-col space-y-6">
				<p>
					Hyprland's color management lets you configure wide-gamut displays to use wider color
					spaces, showing more accurate colors on panels that support them. The problem is that
					grim assumes the captured pixels are sRGB, so it doesn't tone map correctly &mdash; the
					resulting screenshots look washed out. This wrapper uses ImageMagick to set the correct
					source colorspace before converting to sRGB.
				</p>
			</div>

			<h2 class="font-bold italic">Detecting the colorspace</h2>
			<div class="flex flex-col space-y-2">
				<p>
					We query Hyprland for the focused monitor's color management preset, then build an
					ImageMagick command based on the result:
				</p>
				<Code language="bash" code={colorDetect} />
			</div>

			<h2 class="font-bold italic">Converting to sRGB</h2>
			<div class="flex flex-col space-y-2">
				<p>
					ImageMagick's <InlineCode code="-set colorspace" language="bash" /> declares how to
					interpret the input pixels, while <InlineCode code="-colorspace sRGB" language="bash" />
					performs the actual conversion:
				</p>
				<Code language="bash" code={colorConvert} />
			</div>
			<div class="flex flex-col space-y-2">
				<p>
					If a conversion is needed, we pipe grim's output through ImageMagick. Otherwise we pass
					it through unchanged:
				</p>
				<Code language="bash" code={grimPipe} />
			</div>

			<h2 class="font-bold italic">Limitations</h2>
			<div class="flex flex-col space-y-6">
				<p>
					This only covers Display P3 and Adobe RGB. Hyprland supports other presets (<InlineCode
						code="bt2020"
						language="text"
					/>, <InlineCode code="scrgb" language="text" />, etc.) that would need ICC profiles to map
					onto sRGB &mdash; ImageMagick only has certain colorspaces defined natively.
				</p>
			</div>

			<sub class="self-center text-sm">
				<a
					href="https://github.com/rwwiv/.dotfiles-arch/tree/main/.local/bin"
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

	#writeup {
		@apply font-serif;
	}

	#writeup h2 {
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
