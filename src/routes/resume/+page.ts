export const prerender = true;
import type { PageLoad } from "./$types";

interface Employment {
	jobTitle: string;
	employerName: string;
	location: string;
	startDate: Date;
	endDate?: Date;
	tech: string[];
	experience: string[];
}

interface VolunteerRole {
	group: string;
	role: string;
	startYear: Date;
	endYear?: Date;
}

const LM: Employment = {
	jobTitle: "Software Engineer",
	employerName: "Lightmatter",
	location: "Remote",
	startDate: new Date(2022, 2),
	endDate: new Date(2023, 1),
	tech: ["Python", "Django", "Postgres", "Typescript", "AlpineJs", "HTMX"],
	experience: [
		`Contributed to an open-source Django server template using cutting-edge libraries such as Vite, AlpineJs, and HTMX.
		This led to a 30% improvement in first sprint velocity on future projects using the template`,
		`Guided future contributors in developing a frontend component system for this open-source Django project
		which resulted in a twofold increase in components produced`,
		`Collaborated in an agile team to engineer a web application for a health tech client,
		leveraging the aforementioned technologies to reduce the spin-up time of new developers by 40%`,
		`Consulted and served as a staff engineer for this client after being invited to stay on,
		playing an instrumental role in bringing the aforementioned application to market`,
		`Aided in the development of an open-source Python API wrapper for the EHR used by several projects at the company,
		ultimately facilitating a successful deal with this EHR company to refer clients on an ongoing basis`,
	],
};

const SOS: Employment = {
	jobTitle: "Software Developer",
	employerName: "Security Onion Solutions",
	location: "Remote",
	startDate: new Date(2019, 5),
	endDate: new Date(2022, 1),
	tech: ["Python", "Go", "Javascript", "Vue.js", "Bash", "Docker", "Redis"],
	experience: [
		`Developed both an extensible machine learning framework built on Tensorflow and a service app that used this framework for AI-assisted log analysis,
		introducing the first machine learning capability to the product which led to a 50% higher adoption rate among users`,
		`Played an instrumental role in launching a new Docker-based version of a Linux distribution product,
		bringing it from beta to release on a highly accelerated timeline`,
		`Automated end-to-end testing of the product via the Cypress testing framework and served as an advocate for its adoption among team members,
		which reduced release time from 1-2 weeks to half a day`,
		`Architected a Python middleware between the primary Go web app and defined local and remote API services for extensible security data analysis,
		which removed a third-party dependency from the project and allowed for greater product flexibility`,
		// `Streamlined the installation and upgrade code of the product, written in Bash,
		// allowing for a 20% decrease in time spent on further modifications`,
		// `Redesigned the product and company logo as part of a branding update initiative with founding members,
		// which contributed to several instrumental support contracts being brought to the company`,
		// `Developed and maintained a new website using Nuxt as part of this initiative, and created a CI/CD pipeline that included a staging site for change approval.
		// This system allowed for a 200% increase in the velocity of the project and significantly more ownership of the content by those involved`,
	],
};

const RSI: Employment = {
	jobTitle: "Software Consultant",
	employerName: "Rural Sourcing Inc",
	location: "Augusta, GA",
	startDate: new Date(2018, 6),
	endDate: new Date(2019, 5),
	tech: ["Java", "Spring Boot", "Javascript", "Angular", "MySQL"],
	experience: [
		`Worked in an agile team to rewrite a client's legacy desktop application as a full stack Java/Angular web app,
		reducing the time spent in the app by its users by 75%`,
		`Engineered and contributed to several Java (Spring Boot) microservices, including developing unit tests using Junit and Mockito with 100% test coverage,
		which allowed for highly reliable deployments of these services`,
	],
};

const projects = [
	"Building a a wifi-enabled plant moisture and light meter based on the ESP32-C3, written in Rust",
	"Developing and maintaining a portfolio website at rwwiv.com (where you are now!), using Sveltekit",
];

const CFG: VolunteerRole = {
	group: "Code For Greenville",
	role: "Individual contributor",
	startYear: new Date(2022, 1), // Set month to anything since we don't use it
};

export const load = (() => {
	return {
		employments: [LM, SOS, RSI],
		education: {
			degree: "BS in Computer Science",
			university: "Augusta University",
			graduationDate: new Date(2019, 4),
		},
		skills: [
			{ name: "Python", level: 10 },
			{ name: "JS / TS", level: 10 },
			{ name: "Django", level: 8 },
			{ name: "Vue.js", level: 8 },
			{ name: "Go", level: 7 },
			{ name: "Postgres", level: 7 },
			{ name: "Java", level: 6 },
			{ name: "Tensorflow", level: 4 },
			{ name: "MySQL", level: 7 },
			{ name: "Spring Boot", level: 5 },
		],
		projects,
		volunteerRoles: [CFG],
	};
}) satisfies PageLoad;
