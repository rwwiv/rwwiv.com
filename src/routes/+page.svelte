<script lang="ts">
	import FilmStrip from "$lib/components/FilmStrip.svelte";
	import ImageModal from "$lib/components/ImageModal.svelte";

	// Import thumbnails (300px) for fast loading
	// @ts-ignore
	import thumb01 from "$lib/assets/_DSF6125.jpg?w=300";
	// @ts-ignore
	import thumb02 from "$lib/assets/_DSF6147.jpg?w=300";
	// @ts-ignore
	import thumb03 from "$lib/assets/_MTH7383.jpg?w=300";
	// @ts-ignore
	import thumb04 from "$lib/assets/_MTH7513_1.jpg?w=300";
	// @ts-ignore
	import thumb05 from "$lib/assets/_MTH7639.jpg?w=300";
	// @ts-ignore
	import thumb06 from "$lib/assets/_XT30112.jpg?w=300";
	// @ts-ignore
	import thumb07 from "$lib/assets/_XT30374.jpg?w=300";
	// @ts-ignore
	import thumb08 from "$lib/assets/_XT30469_1.jpg?w=300";
	// @ts-ignore
	import thumb09 from "$lib/assets/_XT31495_1.jpg?w=300";
	// @ts-ignore
	import thumb10 from "$lib/assets/_XT31619_1.jpg?w=300";
	// @ts-ignore
	import thumb11 from "$lib/assets/_XT31726.jpg?w=300";
	// @ts-ignore
	import thumb12 from "$lib/assets/_XT31746.jpg?w=300";
	// @ts-ignore
	import thumb13 from "$lib/assets/_XT32284.jpg?w=300";
	// @ts-ignore
	import thumb14 from "$lib/assets/_XT33429.jpg?w=300";
	// @ts-ignore
	import thumb15 from "$lib/assets/_XT33483.jpg?w=300";
	// @ts-ignore
	import thumb16 from "$lib/assets/_XT33875.jpg?w=300";
	// @ts-ignore
	import thumb17 from "$lib/assets/_XT34709.jpg?w=300";
	// @ts-ignore
	import thumb18 from "$lib/assets/_XT34713_1.jpg?w=300";
	// @ts-ignore
	import thumb19 from "$lib/assets/_XT34721.jpg?w=300";
	// @ts-ignore
	import thumb20 from "$lib/assets/_XT35173.jpg?w=300";
	// @ts-ignore
	import thumb21 from "$lib/assets/_XT35265_1.jpg?w=300";
	// @ts-ignore
	import thumb22 from "$lib/assets/_XT35380.jpg?w=300";
	// @ts-ignore
	import thumb23 from "$lib/assets/_XT35386.jpg?w=300";
	// @ts-ignore
	import thumb24 from "$lib/assets/_XT35563.jpg?w=300";
	// @ts-ignore
	import thumb25 from "$lib/assets/_XT37908_2.jpg?w=300";
	// @ts-ignore
	import thumb26 from "$lib/assets/_XT37919_1.jpg?w=300";
	// @ts-ignore
	import thumb27 from "$lib/assets/_XT37993.jpg?w=300";
	// @ts-ignore
	import thumb28 from "$lib/assets/_XT38716_1.jpg?w=300";
	// @ts-ignore
	import thumb29 from "$lib/assets/_XT39365.jpg?w=300";
	// @ts-ignore
	import thumb30 from "$lib/assets/_XT39772_3.jpg?w=300";
	// @ts-ignore
	import thumb31 from "$lib/assets/_XT39877_1.jpg?w=300";
	// @ts-ignore
	import thumb32 from "$lib/assets/_XT39887.jpg?w=300";

	// Import full resolution with metadata for modal preview
	// @ts-ignore
	import full01 from "$lib/assets/_DSF6125.jpg?as=meta";
	// @ts-ignore
	import full02 from "$lib/assets/_DSF6147.jpg?as=meta";
	// @ts-ignore
	import full03 from "$lib/assets/_MTH7383.jpg?as=meta";
	// @ts-ignore
	import full04 from "$lib/assets/_MTH7513_1.jpg?as=meta";
	// @ts-ignore
	import full05 from "$lib/assets/_MTH7639.jpg?as=meta";
	// @ts-ignore
	import full06 from "$lib/assets/_XT30112.jpg?as=meta";
	// @ts-ignore
	import full07 from "$lib/assets/_XT30374.jpg?as=meta";
	// @ts-ignore
	import full08 from "$lib/assets/_XT30469_1.jpg?as=meta";
	// @ts-ignore
	import full09 from "$lib/assets/_XT31495_1.jpg?as=meta";
	// @ts-ignore
	import full10 from "$lib/assets/_XT31619_1.jpg?as=meta";
	// @ts-ignore
	import full11 from "$lib/assets/_XT31726.jpg?as=meta";
	// @ts-ignore
	import full12 from "$lib/assets/_XT31746.jpg?as=meta";
	// @ts-ignore
	import full13 from "$lib/assets/_XT32284.jpg?as=meta";
	// @ts-ignore
	import full14 from "$lib/assets/_XT33429.jpg?as=meta";
	// @ts-ignore
	import full15 from "$lib/assets/_XT33483.jpg?as=meta";
	// @ts-ignore
	import full16 from "$lib/assets/_XT33875.jpg?as=meta";
	// @ts-ignore
	import full17 from "$lib/assets/_XT34709.jpg?as=meta";
	// @ts-ignore
	import full18 from "$lib/assets/_XT34713_1.jpg?as=meta";
	// @ts-ignore
	import full19 from "$lib/assets/_XT34721.jpg?as=meta";
	// @ts-ignore
	import full20 from "$lib/assets/_XT35173.jpg?as=meta";
	// @ts-ignore
	import full21 from "$lib/assets/_XT35265_1.jpg?as=meta";
	// @ts-ignore
	import full22 from "$lib/assets/_XT35380.jpg?as=meta";
	// @ts-ignore
	import full23 from "$lib/assets/_XT35386.jpg?as=meta";
	// @ts-ignore
	import full24 from "$lib/assets/_XT35563.jpg?as=meta";
	// @ts-ignore
	import full25 from "$lib/assets/_XT37908_2.jpg?as=meta";
	// @ts-ignore
	import full26 from "$lib/assets/_XT37919_1.jpg?as=meta";
	// @ts-ignore
	import full27 from "$lib/assets/_XT37993.jpg?as=meta";
	// @ts-ignore
	import full28 from "$lib/assets/_XT38716_1.jpg?as=meta";
	// @ts-ignore
	import full29 from "$lib/assets/_XT39365.jpg?as=meta";
	// @ts-ignore
	import full30 from "$lib/assets/_XT39772_3.jpg?as=meta";
	// @ts-ignore
	import full31 from "$lib/assets/_XT39877_1.jpg?as=meta";
	// @ts-ignore
	import full32 from "$lib/assets/_XT39887.jpg?as=meta";

	// All images in one array
	const allImages = [
		{ src: thumb01, fullSrc: full01, alt: "DSF6125" },
		{ src: thumb02, fullSrc: full02, alt: "DSF6147" },
		{ src: thumb03, fullSrc: full03, alt: "MTH7383" },
		{ src: thumb04, fullSrc: full04, alt: "MTH7513" },
		{ src: thumb05, fullSrc: full05, alt: "MTH7639" },
		{ src: thumb06, fullSrc: full06, alt: "XT30112" },
		{ src: thumb07, fullSrc: full07, alt: "XT30374" },
		{ src: thumb08, fullSrc: full08, alt: "XT30469" },
		{ src: thumb09, fullSrc: full09, alt: "XT31495" },
		{ src: thumb10, fullSrc: full10, alt: "XT31619" },
		{ src: thumb11, fullSrc: full11, alt: "XT31726" },
		{ src: thumb12, fullSrc: full12, alt: "XT31746" },
		{ src: thumb13, fullSrc: full13, alt: "XT32284" },
		{ src: thumb14, fullSrc: full14, alt: "XT33429" },
		{ src: thumb15, fullSrc: full15, alt: "XT33483" },
		{ src: thumb16, fullSrc: full16, alt: "XT33875" },
		{ src: thumb17, fullSrc: full17, alt: "XT34709" },
		{ src: thumb18, fullSrc: full18, alt: "XT34713" },
		{ src: thumb19, fullSrc: full19, alt: "XT34721" },
		{ src: thumb20, fullSrc: full20, alt: "XT35173" },
		{ src: thumb21, fullSrc: full21, alt: "XT35265" },
		{ src: thumb22, fullSrc: full22, alt: "XT35380" },
		{ src: thumb23, fullSrc: full23, alt: "XT35386" },
		{ src: thumb24, fullSrc: full24, alt: "XT35563" },
		{ src: thumb25, fullSrc: full25, alt: "XT37908" },
		{ src: thumb26, fullSrc: full26, alt: "XT37919" },
		{ src: thumb27, fullSrc: full27, alt: "XT37993" },
		{ src: thumb28, fullSrc: full28, alt: "XT38716" },
		{ src: thumb29, fullSrc: full29, alt: "XT39365" },
		{ src: thumb30, fullSrc: full30, alt: "XT39772" },
		{ src: thumb31, fullSrc: full31, alt: "XT39877" },
		{ src: thumb32, fullSrc: full32, alt: "XT39887" },
	];

	// Generate random tilt angles on page load (range: -1.5 to 1.5 degrees)
	function randomTilt() {
		return (Math.random() - 0.5) * 3; // -1.5 to 1.5
	}
	const tilts = Array.from({ length: 12 }, randomTilt);

	// Responsive: track window width and compute strips
	// Default to smallest to avoid layout flash on load
	let innerWidth = 0;

	$: imagesPerStrip = innerWidth > 1400 ? 8 : innerWidth > 1000 ? 6 : 4;

	$: strips = Array.from(
		{ length: Math.ceil(allImages.length / imagesPerStrip) },
		(_, i) => allImages.slice(i * imagesPerStrip, (i + 1) * imagesPerStrip)
	);

	// Modal state
	let modalOpen = false;
	let modalIndex = 0;

	// All images prepared for modal (thumbnail + full resolution with dimensions)
	const modalImages = allImages.map((img) => ({
		thumb: img.src,
		src: img.fullSrc.src,
		width: img.fullSrc.width,
		height: img.fullSrc.height,
		alt: img.alt,
	}));

	function handleSelect(globalIndex: number) {
		modalIndex = globalIndex;
		modalOpen = true;
	}

	function handleClose() {
		modalOpen = false;
	}
</script>

<svelte:window bind:innerWidth />

<svelte:head>
	<title>RWWIV</title>
</svelte:head>

<div class="editing-table">
	{#if innerWidth > 0}
		<div class="strips-container">
			{#each strips as strip, i}
				<FilmStrip
					images={strip}
					maxImages={imagesPerStrip}
					startIndex={i * imagesPerStrip}
					tilt={tilts[i % tilts.length]}
					on:select={(e) => handleSelect(i * imagesPerStrip + e.detail.index)}
				/>
			{/each}
		</div>
	{/if}
</div>

<ImageModal images={modalImages} currentIndex={modalIndex} open={modalOpen} on:close={handleClose} />

<style lang="postcss">
	.editing-table {
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 1rem 2rem;
		overflow: hidden;
	}

	.strips-container {
		display: flex;
		flex-direction: column;
		gap: 4rem;
		padding: 1rem;
	}

	@media (max-width: 768px) {
		.strips-container {
			gap: 3rem;
			padding: 0.5rem;
		}

		.editing-table {
			padding: 0.5rem 0.25rem;
		}
	}

	@media (max-width: 480px) {
		.editing-table {
			padding: 0.5rem 0;
		}

		.strips-container {
			padding: 0.25rem;
		}
	}
</style>
