package data

import (
	"github.com/ism0080/profile/models"
	"github.com/ism0080/profile/templates/icons"
)

func HomePageData(baseUrl string) models.HomePage {
	return models.HomePage{
		Paragraph: []string{"Hi,", "I'm Isaac,", "Software Engineer"},
		TechStack: "{ .Net Core • React • AWS }",
		SocialButtons: models.SocialButtons{
			models.NewSocialButton("https://github.com/ism0080", "GitHub", icons.GitHub()),
			models.NewSocialButton("https://www.linkedin.com/in/ism0080", "LinkedIn", icons.LinkedIn()),
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
			},
			Footer: models.Footer{
				Email: "isaac.mackle@gmail.com",
				Links: []models.FooterLink{
					models.NewFooterLink("https://github.com/ism0080", "GitHub"),
					models.NewFooterLink("https://www.linkedin.com/in/ism0080", "LinkedIn"),
				},
			},
		},
	}
}
