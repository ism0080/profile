package data

import (
	"github.com/ism0080/profile/cmd/models"
	"github.com/ism0080/profile/cmd/templates/icons"
)

func HomePageData(baseUrl string) models.HomePage {
	return models.HomePage{
		Paragraph: []string{"Hi,", "I'm Isaac,", "Software Engineer"},
		TechStack: "{ .Net Core • React • AWS }",
		SocialButtons: models.SocialButtons{
			models.NewSocialButton(githubUrl, "GitHub", icons.GitHub()),
			models.NewSocialButton(linkedinUrl, "LinkedIn", icons.LinkedIn()),
		},
		CompletedProjects: completedProjects(),
		PersonalProjects:  personalProjects(),
		Base: models.Page{
			Title: "Isaac Mackle • Software Engineer",
			Meta: models.Meta{
				Title:       "Isaac Mackle • Software Engineer",
				Description: "Software Engineer who is always looking to learn more. Front-end focused with experience in React, Typescript. Get to know me.",
				Url:         baseUrl,
				Image:       "./assets/meta_share.png",
			},
			Navigation: models.Navigation{
				models.NewMenuItem("Home", "./index.html"),
				models.NewMenuItem("Blog", "/blog/index.html"),
			},
			Footer: models.Footer{
				Email: email,
				Links: []models.FooterLink{
					models.NewFooterLink(githubUrl, "GitHub"),
					models.NewFooterLink(linkedinUrl, "LinkedIn"),
				},
			},
		},
	}
}
