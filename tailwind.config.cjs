/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			screens: {
				xxs: '360px',
				xs: '480px',
				'2xl': '1536px',
				'3xl': '2200px',
				'hover-d': { raw: '(hover: hover)' }
			}
		}
	},
	plugins: []
};
