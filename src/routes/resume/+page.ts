import type { PageLoad } from './$types';

interface Employment {
	jobTitle: string;
	employerName: string;
	location: string;
	startDate: Date;
	endDate?: Date;
	experience: string[];
}

const LM: Employment = {
	jobTitle: 'Software Engineer',
	employerName: 'Lightmatter',
	location: 'Remote',
	startDate: new Date(2022, 2),
	endDate: new Date(2023, 1),
	experience: [
		`Contributed to an open-source Django project template using cutting-edge libraries such as Vite, AlpineJs, and HTMX`,
		`Provided guidance to an intern working on a frontend component system for the open-source Django project`,
		`Collaborated in an agile team to engineer a web application for a health tech client, leveraging the aforementioned technologies and adhering to design documentation`,
		`Acted as a consultant and staff engineer for this client after being asked to stay on, and played an instrumental role in bringing the aforementioned application to market`,
		`Contributed to the development of a Python API wrapper for the EHR used by several projects at the company`
	]
};

const SOS: Employment = {
	jobTitle: 'Software Developer',
	employerName: 'Security Onion Solutions',
	location: 'Remote',
	startDate: new Date(2019, 5),
	endDate: new Date(2022, 1),
	experience: [
		`Played an instrumental role in launching a new Docker-based version of a Linux distribution product, working as part of a team to bring it from beta to release`,
		`Introduced automated end-to-end testing to the product via the Cypress testing framework and served as an advocate for its adoption among team members`,
		`Streamlined the installation and upgrade code of the product, written in Bash, through significant refactoring efforts`,
		`Redesigned the product and company logo as part of a branding update initiative with founding members of the company`,
		`Developed and maintained a new website using Nuxt as part of this initiative, and created a CI/CD pipeline that included a staging site for change approval`,
		`Incorporated a beta machine learning-based log anomaly detection tool to the product and established an extensible ML framework using Tensorflow for future applications`
	]
};

const RSI: Employment = {
	jobTitle: 'Junior Developer',
	employerName: 'Rural Sourcing Inc.',
	location: 'Augusta, GA',
	startDate: new Date(2018, 6),
	endDate: new Date(2019, 5),
	experience: [
		`As part of an agile team, worked to rewrite a client's legacy desktop application as a full stack Java/Angular web app`,
		`Engineered and contributed to several Java (Spring Boot) microservices, including developing unit tests using Junit and Mockito with 100% test coverage`
	]
};

export const load = (() => {
	return {
		employments: [LM, SOS, RSI],
		education: {
			degree: 'BS in Computer Science',
			university: 'Augusta University',
			graduationDate: new Date(2019, 4)
		},
		skills: [
			{ name: 'JS / TS', level: 10 },
			{ name: 'Python', level: 10 },
			{ name: 'Vue', level: 8 },
			{ name: 'Go', level: 7 },
			{ name: 'Postgres', level: 5 },
			{ name: 'Tensorflow', level: 4 },
			{ name: 'Java', level: 6 },
			{ name: 'C#', level: 6 }
		]
	};
}) satisfies PageLoad;
