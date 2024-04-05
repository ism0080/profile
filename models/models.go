package models

import (
	"fmt"

	"github.com/a-h/templ"
)

type MenuItem struct {
	Name string
	Url  string
}

func NewMenuItem(name, url string) MenuItem {
	return MenuItem{
		Name: name,
		Url:  url,
	}
}

type Navigation = []MenuItem

type Meta struct {
	Title       string
	Description string
	Image       string
	Url         string
}

type Footer struct {
	Email string
	Links []FooterLink
}

type FooterLink struct {
	Url   string
	Label string
}

func NewFooterLink(url, label string) FooterLink {
	return FooterLink{
		Url:   url,
		Label: label,
	}
}

type SocialButton struct {
	Url       string
	Label     string
	Component templ.Component
}

func NewSocialButton(url, label string, component templ.Component) SocialButton {
	return SocialButton{
		Url:       url,
		Label:     label,
		Component: component,
	}
}

type SocialButtons = []SocialButton

type Page struct {
	Title      string
	Meta       Meta
	Navigation Navigation
	Footer     Footer
}

type HomePage struct {
	Base              Page
	Paragraph         []string
	TechStack         string
	SocialButtons     SocialButtons
	PersonalProjects  PersonalProjects
	CompletedProjects CompletedProjects
}

type ResumePage struct {
	Base   Page
	Resume Resume
}

type CompletedProjects struct {
	Title    string
	Projects Projects
}

type PersonalProjects struct {
	Title    string
	Projects Projects
}

type Projects = []Project

type Project struct {
	Title       string
	Description string
	Image       ProjectImage
	Target      string
	Source      string
}

type ProjectImage struct {
	Src string
	Alt string
}

type ProjectResponse struct {
	Items []ProjectResponseItem `json:"items"`
}

type ProjectResponseItem struct {
	CollectionId string `json:"collectionId"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	Image        string `json:"image"`
	ImageAlt     string `json:"image_alt"`
	ProjectType  string `json:"project_type"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	Title        string `json:"title"`
}

func (p ProjectResponse) ToProjects(basePath string) Projects {
	var d = Projects{}
	for _, project := range p.Items {
		d = append(d, Project{
			Title:       project.Title,
			Description: project.Description,
			Target:      project.Target,
			Source:      project.Source,
			Image: ProjectImage{
				Src: fmt.Sprintf("%s/files/%s/%s/%s", basePath, project.CollectionId, project.Id, project.Image),
				Alt: project.ImageAlt,
			},
		})
	}

	return d
}

type Resume struct {
	Title      string
	Blurb      string
	Info       ResumeInfo
	Experience ResumeExperience
	Skills     ResumeSkills
}

type ResumeInfo struct {
	Phone    string
	Email    string
	Location string
	Website  string
	GitHub   string
	LinkedIn string
}

type ResumeExperience struct {
	Title string
	Jobs  []ResumeJob
}

type ResumeJob struct {
	Role        string
	Company     string
	Description string
	Examples    []ResumeJobExample
}

type ResumeJobExample struct {
	Role        string
	Experiences []ResumeJobExampleExperience
}

type ResumeJobExampleExperience struct {
	Project     string
	Description string
	Technology  string
}

type ResumeSkills struct {
	Title  string
	Skills []ResumeSkillValue
	Other  ResumeSkillOther
}

type ResumeSkillOther struct {
	Title  string
	Skills []ResumeSkillValue
}

type ResumeSkillValue struct {
	Key   string
	Value string
}

func NewResumeSkillValue(key, value string) ResumeSkillValue {
	return ResumeSkillValue{
		Key:   key,
		Value: value,
	}
}
