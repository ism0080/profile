package data

import (
	"github.com/ism0080/profile/models"
)

func ResumeData() models.Resume {
	return models.Resume{
		Title: "Isaac Mackle",
		Blurb: "Software engineer specialising in React, .Net Core and AWS. Prepared to face any challenge thrown at me, I am a hardworking fast learner with proven experience to work will in a team environment.",
		Info: models.ResumeInfo{
			Phone:    "+6427 788 4011",
			Email:    "isaac.mackle@gmail.com",
			Location: "Christchurch",
			Website:  "isaacmackle.com",
			GitHub:   "github.com/ism0080",
			LinkedIn: "linkedin.com/in/ism0080",
		},
		Experience: models.ResumeExperience{
			Title: "Experience",
			Jobs: []models.ResumeJob{
				{
					Role:        "Software Engineer",
					Company:     "Domino's Pizza Enterprise, Brisbane - (2019 - 2021)",
					Description: "I contributed towards many projects for different markets (AU, NZ, JP, NL, FR, BE, DE, DK) as a frontend/backend software developer.",
					Examples: []models.ResumeJobExample{
						{
							Role: "Graduate (2019 - 2021)",
							Experiences: []models.ResumeJobExampleExperience{
								{
									Project:    "New online ordering platform (website/ native app)",
									Technology: "React, React Native, TypeScript, Jest, Storybook, Monorepo, Redux",
								},
								{
									Project:     "Pizza Prep",
									Description: "Trained in a Domino's store to gain knowledge of the business and how IT improves the pizza making process.",
								},
								{
									Project:     "Adyen payment service migration and report notifications",
									Description: "This included updating current Apple pay implementation",
									Technology:  ".Net core, C#, redis, swagger, AWS (Lambda, S3, SQS, Cloud Formation), NUnit, Azure DevOps pipelines",
								},
								{
									Project:     "Online ordering website/ wrapper app maintenance",
									Description: "This included implementing Launch Darkly feature toggles and implementing Apple Sign In OAuth flow for wrapper app and web",
									Technology:  "ASP .Net framework, C#, JavaScript, SASS, Azure (CosmosDB, App Insights, App services), TeamCity (pipelines), Octopus (deploying builds to Azure)",
								},
							},
						},
						{
							Role: "Associate (2021)",
							Experiences: []models.ResumeJobExampleExperience{
								{
									Project:     "GPS Driver Tracker",
									Description: "Working with Domino's US to implement their GPS solution for DPE through the use of messaging systems, APIs, and frontend applications",
									Technology:  ".Net core, C#, ActiveMQ, Docker, React, AWS (Kinesis Firehose, Cloudformation, ECS, CloudWatch), Azure Devops Pipelines, Application Insights",
								},
							},
						},
					},
				},
			},
		},
		Skills: models.ResumeSkills{
			Title: "Skills",
			Skills: []models.ResumeSkillValue{
				models.NewResumeSkillValue("HTML5 | CSS", "Less, CSS Modules"),
				models.NewResumeSkillValue("JavaScript", "ES6, React, TypeScript, GraphQL, Apollo, NextJS"),
				models.NewResumeSkillValue("C#", ".Net Core, API"),
				models.NewResumeSkillValue("Testing", "Jest, React Testing Library, XUnit/NUnit"),
				models.NewResumeSkillValue("UI | Design", "Storybook, Sketch, Zeplin, Figma"),
				models.NewResumeSkillValue("Server-Side", "NodeJS, C#, Go"),
				models.NewResumeSkillValue("AWS", "S3, CloudFormation, Lambda, DynamoDB, Kinesis Firehose, SQS, SNS"),
				models.NewResumeSkillValue("Other", "CI/CD Pipelines, Docker"),
				models.NewResumeSkillValue("Version Control", "Git, GitHub, BitBucket, Azure DevOps, Perforce"),
				models.NewResumeSkillValue("Project Management", "Agile/Scrum, Azure Devops, Jira, Asana"),
			},
			Other: models.ResumeSkillOther{
				Title: "Experience With",
				Skills: []models.ResumeSkillValue{
					models.NewResumeSkillValue("iOS", "Swift, SwiftUI"),
					models.NewResumeSkillValue("Python", "Django, Flask"),
				},
			},
		},
	}
}
