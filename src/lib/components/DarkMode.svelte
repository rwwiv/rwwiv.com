<script lang="ts">
	import { browser } from "$app/environment";
	const toggleTheme = () => {
		const isDark = window.document.documentElement.classList.toggle("dark");
		localStorage.setItem("color-theme", isDark ? "dark" : "light");
	};

	const useSystemColor = () => {
		localStorage.removeItem("color-theme");
		if (browser) {
			window.matchMedia("(prefers-color-scheme: dark)").matches
				? window.document.documentElement.classList.add("dark")
				: window.document.documentElement.classList.remove("dark");
		}
	};
</script>

<svelte:head>
	<script>
		if (window) {
			function setDarkMode() {
				localStorage.getItem("color-theme") === "dark" ||
				(!("color-theme" in localStorage) && darkModePreference.matches)
					? window.document.documentElement.classList.add("dark")
					: window.document.documentElement.classList.remove("dark");
			}

			const darkModePreference = window.matchMedia("(prefers-color-scheme: dark)");

			setDarkMode();

			darkModePreference.addEventListener("change", () => setDarkMode());
		}
	</script>
</svelte:head>

<div class="flex flex-col items-center space-y-3 select-none">
	<div
		on:click={toggleTheme}
		on:keypress={toggleTheme}
		class="cursor-pointer hover:dark:text-gray-500 hover:text-gray-700 "
	>
		<div class="hidden dark:block active:text-purple-500">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 20 20"
				fill="currentColor"
				class="w-5 h-5"
			>
				<path
					d="M10 2a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 2zM10 15a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 0110 15zM10 7a3 3 0 100 6 3 3 0 000-6zM15.657 5.404a.75.75 0 10-1.06-1.06l-1.061 1.06a.75.75 0 001.06 1.06l1.06-1.06zM6.464 14.596a.75.75 0 10-1.06-1.06l-1.06 1.06a.75.75 0 001.06 1.06l1.06-1.06zM18 10a.75.75 0 01-.75.75h-1.5a.75.75 0 010-1.5h1.5A.75.75 0 0118 10zM5 10a.75.75 0 01-.75.75h-1.5a.75.75 0 010-1.5h1.5A.75.75 0 015 10zM14.596 15.657a.75.75 0 001.06-1.06l-1.06-1.061a.75.75 0 10-1.06 1.06l1.06 1.06zM5.404 6.464a.75.75 0 001.06-1.06l-1.06-1.06a.75.75 0 10-1.061 1.06l1.06 1.06z"
				/>
			</svg>
		</div>
		<div class="block dark:hidden active:text-orange-500">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 20 20"
				fill="currentColor"
				class="w-5 h-5"
			>
				<path
					fill-rule="evenodd"
					d="M7.455 2.004a.75.75 0 01.26.77 7 7 0 009.958 7.967.75.75 0 011.067.853A8.5 8.5 0 116.647 1.921a.75.75 0 01.808.083z"
					clip-rule="evenodd"
				/>
			</svg>
		</div>
	</div>
	<div
		on:click={useSystemColor}
		on:keypress={useSystemColor}
		class="cursor-pointer hover:dark:text-gray-500 hover:text-gray-700 active:text-blue-400 dark:active:text-blue-500 text-xs"
	>
		Use system dark mode
	</div>
</div>
