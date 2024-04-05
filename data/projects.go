package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ism0080/profile/models"
)

type ProjectType string

const (
	Personal    ProjectType = "personal"
	Completed   ProjectType = "completed"
	ApiBasePath             = "https://admin.isaacmackle.com/api"
)

func completedProjects() models.CompletedProjects {
	return models.CompletedProjects{
		Title:    "Completed Projects",
		Projects: getProjects(Completed),
	}
}

func personalProjects() models.PersonalProjects {
	return models.PersonalProjects{
		Title:    "Personal Projects",
		Projects: getProjects(Personal),
	}
}

func getProjects(projectType ProjectType) models.Projects {
	res, err := http.Get(fmt.Sprintf("%s/collections/projects/records?fields=title,description,image,project_type,source,target,collectionId,id,img_alt&filter=(project_type='%s')", ApiBasePath, projectType))
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer res.Body.Close()

	var data models.ProjectResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return data.ToProjects(ApiBasePath)
}
