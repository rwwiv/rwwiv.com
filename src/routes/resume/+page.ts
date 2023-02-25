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
		`Contributed to an open-source Django server template using cutting-edge libraries such as Vite, AlpineJs, and HTMX.
		This led to a 30% improvement in first sprint velocity on future projects using the template`,
		`Guided future contributors in developing a frontend component system for this open-source Django project
		which resulted in a twofold increase in components produced`,
		`Collaborated in an agile team to engineer a web application for a health tech client,
		leveraging the aforementioned technologies to reduce spin-up time of new developers by 40%`,
		`Consulted and served as a staff engineer for this client after being invited to stay on,
		playing an instrumental role in bringing the aforementioned application to market`,
		`Also contributed to an open source Python API wrapper for the EHR used by several projects at the company,
		ultimately facilitating a successful deal with this EHR company to refer clients on an ongoing basis`
	]
};

const SOS: Employment = {
	jobTitle: 'Software Developer',
	employerName: 'Security Onion Solutions',
	location: 'Remote',
	startDate: new Date(2019, 5),
	endDate: new Date(2022, 1),
	experience: [
		`Played an instrumental role in launching a new Docker-based version of a Linux distribution product,
		bringing it from beta to release on a highly accelerated timeline`,
		`Introduced automated end-to-end testing to the product via the Cypress testing framework and served as an advocate for its adoption among team members,
		which resulted in time spent on releasing a new build dropping from 1-2 weeks to half a day`,
		`Streamlined the installation and upgrade code of the product, written in Bash,
		allowing for a 20% decrease in time spent on further modifications`,
		`Redesigned the product and company logo as part of a branding update initiative with founding members,
		which contributed to several instrumental support contracts being brought to the company`,
		`Developed and maintained a new website using Nuxt as part of this initiative, and created a CI/CD pipeline that included a staging site for change approval.
		This system allowed for a 100% increase in velocity on the project and significantly more ownership of the content by those involved`,
		`Incorporated a beta machine learning-based log anomaly detection tool into the product and established an extensible ML framework using Tensorflow for future applications,
		which led to the introduction of the first ML feature to the product while allowing for significant increases in value of the product moving forward`
	]
};

const RSI: Employment = {
	jobTitle: 'Junior Developer',
	employerName: 'Rural Sourcing Inc.',
	location: 'Augusta, GA',
	startDate: new Date(2018, 6),
	endDate: new Date(2019, 5),
	experience: [
		`As part of an agile team, worked to rewrite a client's legacy desktop application as a full stack Java/Angular web app
		reducing the time spent in the app by its users by 75%`,
		`Engineered and contributed to several Java (Spring Boot) microservices, including developing unit tests using Junit and Mockito with 100% test coverage,
		producing fully working services in half the time of other developers on the project`
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
